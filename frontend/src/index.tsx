import '@/assets/styles/index.css';

import { router } from '@/router';
import { createRoot } from 'react-dom/client';

import { StrictMode } from 'react';
import { RouterProvider } from '@tanstack/react-router';

const rootElement = document.getElementById('root') as HTMLElement;

if (rootElement && !rootElement.innerHTML) {
   const container = createRoot(rootElement);

   container.render(
      <StrictMode>
         <RouterProvider router={router} />
      </StrictMode>
   );
}
