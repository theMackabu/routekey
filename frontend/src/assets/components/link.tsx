import React, { forwardRef } from 'react';
import * as Headless from '@headlessui/react';
import { Link as RouterLink } from '@tanstack/react-router';

export const Link = forwardRef(function Link(
   props: { to: string } & React.ComponentPropsWithoutRef<'a'>,
   ref: React.ForwardedRef<HTMLAnchorElement>
) {
   return (
      <Headless.DataInteractive>
         <RouterLink {...props} ref={ref} />
      </Headless.DataInteractive>
   );
});
