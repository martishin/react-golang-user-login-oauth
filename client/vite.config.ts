import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8100', // Your local backend URL
        changeOrigin: true,
        secure: false, // Disable SSL verification for local development
      },
      '/auth': {
        target: 'http://localhost:8100', // Your local backend URL
        changeOrigin: true,
        secure: false,
      },
    },
  },
});
