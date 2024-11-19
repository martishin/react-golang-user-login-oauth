import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig(({ mode }) => ({
  plugins: [react()],
  server:
    mode === "development"
      ? {
          proxy: {
            "/api": {
              target: "http://localhost:8100",
              changeOrigin: true,
              secure: false,
            },
            "/auth": {
              target: "http://localhost:8100",
              changeOrigin: true,
              secure: false,
            },
          },
        }
      : undefined, // No proxy in production
}));
