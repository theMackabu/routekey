import { Fragment } from 'react';
import { Outlet, createRootRoute } from '@tanstack/react-router';
import { TanStackRouterDevtools } from '@tanstack/react-router-devtools';

export const rootRoute = createRootRoute({
   component: () => (
      <Fragment>
         <Outlet />
         <TanStackRouterDevtools />
      </Fragment>
   )
});
