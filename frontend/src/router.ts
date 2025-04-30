import { Home } from '@/pages/home';
import { Error } from '@/pages/error';
import { rootRoute } from '@/pages/root';
import { createRoute, createRouter } from '@tanstack/react-router';

const indexRoute = createRoute({
   path: '/',
   component: Home,
   getParentRoute: () => rootRoute
});

export const router = createRouter({
   defaultPreload: 'intent',
   scrollRestoration: true,
   defaultStructuralSharing: true,
   defaultNotFoundComponent: Error,
   defaultPreloadStaleTime: 0,
   routeTree: rootRoute.addChildren([indexRoute])
});
