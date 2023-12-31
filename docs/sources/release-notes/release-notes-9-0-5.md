---
aliases: []
hide_menu: true
labels:
  products:
    - cloud
    - enterprise
    - oss
title: Release notes for Grafana 9.0.5
---

<!-- Auto generated by update changelog github action -->

# Release notes for Grafana 9.0.5

### Features and enhancements

- **Access control:** Show dashboard settings to users who can edit dashboard. [#52535](https://github.com/grafana/grafana/pull/52535), [@grafanabot](https://github.com/grafanabot)
- **Alerting:** Allow the webhook notifier to support a custom Authorization header. [#52515](https://github.com/grafana/grafana/pull/52515), [@gotjosh](https://github.com/gotjosh)
- **Chore:** Upgrade to Go version 1.17.12. [#52523](https://github.com/grafana/grafana/pull/52523), [@sakjur](https://github.com/sakjur)
- **Plugins:** Add signature wildcard globbing for dedicated private plugin type. [#52163](https://github.com/grafana/grafana/pull/52163), [@wbrowne](https://github.com/wbrowne)
- **Prometheus:** Don't show errors from unsuccessful API checks like rules or exemplar checks. [#52193](https://github.com/grafana/grafana/pull/52193), [@darrenjaneczek](https://github.com/darrenjaneczek)

### Bug fixes

- **Access control:** Allow organisation admins to add existing users to org (#51668). [#52553](https://github.com/grafana/grafana/pull/52553), [@vtorosyan](https://github.com/vtorosyan)
- **Alerting:** Fix alert panel instance-based rules filtering. [#52583](https://github.com/grafana/grafana/pull/52583), [@konrad147](https://github.com/konrad147)
- **Apps:** Fixes navigation between different app plugin pages. [#52571](https://github.com/grafana/grafana/pull/52571), [@torkelo](https://github.com/torkelo)
- **Cloudwatch:** Upgrade grafana-aws-sdk to fix auth issue with secret keys. [#52420](https://github.com/grafana/grafana/pull/52420), [@sarahzinger](https://github.com/sarahzinger)
- **Grafana/toolkit:** Fix incorrect image and font generation for plugin builds. [#52661](https://github.com/grafana/grafana/pull/52661), [@academo](https://github.com/academo)
- **Loki:** Fix `show context` not working in some occasions. [#52458](https://github.com/grafana/grafana/pull/52458), [@svennergr](https://github.com/svennergr)
- **RBAC:** Fix permissions on dashboards and folders created by anonymous users. [#52615](https://github.com/grafana/grafana/pull/52615), [@gamab](https://github.com/gamab)
