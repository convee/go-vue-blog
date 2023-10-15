import { defineConfig } from 'vite'
import { resolve } from 'path'
import { createVitePlugins } from './build/plugin'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: createVitePlugins(),
  server: {
    host: "0.0.0.0", //内网访问
  },
  resolve: {
    alias: {
      // 关键代码
      "@": resolve(__dirname, "src")
    },
  },
});