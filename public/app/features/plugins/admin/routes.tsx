import { SafeDynamicImport } from 'app/core/components/DynamicImports/SafeDynamicImport';
import { RouteDescriptor } from 'app/core/navigation/types';

import { PluginAdminRoutes } from './types';

const DEFAULT_ROUTES = [
  {
    path: '/plugins',
    navId: 'plugins',
    roles: () => ['Admin', 'ServerAdmin'],
    routeName: PluginAdminRoutes.Home,
    component: SafeDynamicImport(() => import(/* webpackChunkName: "PluginListPage" */ './pages/Browse')),
  },
  {
    path: '/plugins/browse',
    navId: 'plugins',
    roles: () => ['Admin', 'ServerAdmin'],
    routeName: PluginAdminRoutes.Browse,
    component: SafeDynamicImport(() => import(/* webpackChunkName: "PluginListPage" */ './pages/Browse')),
  },
  {
    path: '/plugins/:pluginId/',
    navId: 'plugins',
    roles: () => ['Admin', 'ServerAdmin'],
    routeName: PluginAdminRoutes.Details,
    component: SafeDynamicImport(() => import(/* webpackChunkName: "PluginPage" */ './pages/PluginDetails')),
  },
];

export function getRoutes(): RouteDescriptor[] {
  return DEFAULT_ROUTES;
}
