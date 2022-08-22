import { fileURLToPath, URL } from "node:url";
import vuePlugin from "@vitejs/plugin-vue";
import { defineConfig } from "vite";
// import vue from "@vitejs/plugin-vue";
import legacyPlugin from "@vitejs/plugin-legacy";

const rollupOptions = {
  output: {
    entryFileNames: `static/[name].[hash].js`,
    chunkFileNames: `js/[name].[hash].js`,
    assetFileNames: `assets/[name].[hash].[ext]`,
  },
};

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    legacyPlugin({
      targets: ["Chrome >= 98"],
    }),
    // vue(),
    vuePlugin({}),
  ],
  css: {
    preprocessorOptions: {
      less: {
        // 支持内联 JavaScript
        javascriptEnabled: true,
      },
    },
  },
  base: "./", // index.html文件所在位置
  root: "./", // js导入的资源路径，src
  server: {
    // 如果使用docker-compose开发模式，设置为false
    port: 3500,
    proxy: {
      // 把key的路径代理到target位置
      // detail: https://cli.vuejs.org/config/#devserver-proxy
      "/api": {
        // 需要代理的路径   例如 '/api'
        target: `${process.env.VITE_BASE_PATH}:${process.env.VITE_SERVER_PORT}/`, // 代理到 目标路径
        changeOrigin: true,
        rewrite: (path) => path.replace(new RegExp("^" + process.env.VITE_BASE_API), ""),
      },
    },
  },
  build: {
    target: "es2015",
    minify: "terser", // 是否进行压缩,boolean | 'terser' | 'esbuild',默认使用terser
    manifest: false, // 是否产出maifest.json
    sourcemap: false, // 是否产出soucemap.json
    outDir: "dist", // 产出目录
    rollupOptions,
  },
  esbuild: {},
  optimizeDeps: {},
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
});
