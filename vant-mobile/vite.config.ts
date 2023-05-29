import * as path from "path";
import vuePlugin from "@vitejs/plugin-vue";
import { defineConfig, loadEnv } from "vite";
// import { VantResolver } from "unplugin-vue-components/resolvers";
// import Components from "unplugin-vue-components/vite";

const rollupOptions = {
  output: {
    entryFileNames: `static/[name].[hash].js`,
    chunkFileNames: `js/[name].[hash].js`,
    assetFileNames: `assets/[name].[hash].[ext]`,
  },
};
// console.log(import.meta);
// https://vitejs.dev/config/

// mode 什么环境
export default defineConfig(({ mode }: { mode: string }) => {
  const env = loadEnv(mode, process.cwd());
  return {
    plugins: [
      // vue(),
      vuePlugin({}),
      // Components({
      //   resolvers: [VantResolver()],
      // }),
    ],
    css: {
      preprocessorOptions: {
        less: {
          javascriptEnabled: true,
        },
      },
    },
    base: "./", // index.html文件所在位置
    root: "./", // js导入的资源路径，src
    server: {
      // 如果使用docker-compose开发模式，设置为false
      cors: true,
      hmr: true,
      host: "0.0.0.0",
      port: 3800,
      proxy: {
        // 把key的路径代理到target位置
        // detail: https://cli.vuejs.org/config/#devserver-proxy
        [env.VITE_BASE_API]: {
          // 需要代理的路径   例如 '/api'
          target: `${env.VITE_BASE_PATH}:${env.VITE_SERVER_PORT}`, // 代理到 目标路径
          changeOrigin: true,
          secure: true,
          rewrite: (path) =>
            path.replace(new RegExp(`^${env.VITE_BASE_API}`), "/api"),
          // rewrite: (path) => path.replace("", ""),
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
        "@": path.resolve(__dirname, "./src/"),
      },
    },
  };
});
