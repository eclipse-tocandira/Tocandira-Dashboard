package dashboards

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/grafana/grafana/pkg/services/dashboards"

	"github.com/grafana/grafana/pkg/bus"

	"github.com/grafana/grafana/pkg/components/simplejson"
	"github.com/grafana/grafana/pkg/log"
	"github.com/grafana/grafana/pkg/models"
	gocache "github.com/patrickmn/go-cache"
)

type fileReader struct {
	Cfg           *DashboardsAsConfig
	Path          string
	log           log.Logger
	dashboardRepo dashboards.Repository
	cache         *gocache.Cache
}

func NewDashboardFilereader(cfg *DashboardsAsConfig, log log.Logger) (*fileReader, error) {
	path, ok := cfg.Options["folder"].(string)
	if !ok {
		return nil, fmt.Errorf("Failed to load dashboards. folder param is not a string")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Error("Cannot read directory", "error", err)
	}

	return &fileReader{
		Cfg:           cfg,
		Path:          path,
		log:           log,
		dashboardRepo: dashboards.GetRepository(),
		cache:         gocache.New(5*time.Minute, 10*time.Minute),
	}, nil
}

func (fr *fileReader) addCache(key string, json *dashboards.SaveDashboardItem) {
	fr.cache.Add(key, json, time.Minute*10)
}

func (fr *fileReader) getCache(key string) (*dashboards.SaveDashboardItem, bool) {
	obj, exist := fr.cache.Get(key)
	if !exist {
		return nil, exist
	}

	dash, ok := obj.(*dashboards.SaveDashboardItem)
	if !ok {
		return nil, ok
	}

	return dash, ok
}

func (fr *fileReader) ReadAndListen(ctx context.Context) error {
	ticker := time.NewTicker(time.Second * 5)

	if err := fr.walkFolder(); err != nil {
		fr.log.Error("failed to search for dashboards", "error", err)
	}

	for {
		select {
		case <-ticker.C:
			fr.walkFolder()
		case <-ctx.Done():
			return nil
		}
	}
}

func (fr *fileReader) walkFolder() error {
	if _, err := os.Stat(fr.Path); err != nil {
		if os.IsNotExist(err) {
			return err
		}
	}

	return filepath.Walk(fr.Path, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if f.IsDir() {
			if strings.HasPrefix(f.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}

		if !strings.HasSuffix(f.Name(), ".json") {
			return nil
		}

		cachedDashboard, exist := fr.getCache(path)
		if exist && cachedDashboard.UpdatedAt == f.ModTime() {
			return nil
		}

		dash, err := fr.readDashboardFromFile(path)
		if err != nil {
			fr.log.Error("failed to load dashboard from ", "file", path, "error", err)
			return nil
		}

		cmd := &models.GetDashboardQuery{Slug: dash.Dashboard.Slug}
		err = bus.Dispatch(cmd)

		if err == models.ErrDashboardNotFound {
			fr.log.Debug("saving new dashboard", "file", path)
			_, err = fr.dashboardRepo.SaveDashboard(dash)
			return err
		}

		if err != nil {
			fr.log.Error("failed to query for dashboard", "slug", dash.Dashboard.Slug, "error", err)
			return nil
		}

		if cmd.Result.Updated.Unix() >= f.ModTime().Unix() {
			return nil
		}

		fr.log.Debug("no dashboard in cache. loading dashboard from disk into database.", "file", path)
		_, err = fr.dashboardRepo.SaveDashboard(dash)
		return err
	})
}

func (fr *fileReader) readDashboardFromFile(path string) (*dashboards.SaveDashboardItem, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	data, err := simplejson.NewFromReader(reader)
	if err != nil {
		return nil, err
	}

	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	dash, err := createDashboardJson(data, stat.ModTime(), fr.Cfg)
	if err != nil {
		return nil, err
	}

	fr.addCache(path, dash)

	return dash, nil
}
