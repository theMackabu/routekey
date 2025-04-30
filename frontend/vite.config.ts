import { defineConfig } from 'vite';
import { resolve } from 'node:path';

import react from '@vitejs/plugin-react';
import tailwindcss from '@tailwindcss/vite';

export default defineConfig({
   plugins: [react(), tailwindcss()],
   test: {
      globals: true,
      environment: 'jsdom'
   },
   resolve: {
      alias: { '@': resolve(__dirname, './src') }
   }
});
