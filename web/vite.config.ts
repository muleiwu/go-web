import { fileURLToPath, URL } from 'node:url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import vueDevTools from 'vite-plugin-vue-devtools'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'
import { visualizer } from 'rollup-plugin-visualizer';
import PortFinder from 'portfinder';


// https://vitejs.dev/config/
export default defineConfig(async () => {
  const port = await PortFinder.getPortPromise({ port: 5173 }); // 从3000开始寻找可用端口
  return {
    server: {
      port: port,
    },
    plugins: [
      vue(),
      vueJsx(),
      visualizer(),
      vueDevTools(),
      AutoImport({
        imports: [
          'vue',
          {
            'naive-ui': [
              'useDialog',
              'useMessage',
              'useNotification',
              'useLoadingBar'
            ]
          }
        ]
      }),
      Components({
        resolvers: [NaiveUiResolver()]
      }),
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      }
    },
    build:{
      chunkSizeWarningLimit: 500,
      rollupOptions:{
        output: {
          manualChunks(id: any) {
            if (id.includes("node_modules")) {
              // 让每个插件都打包成独立的文件
              return id.toString().split("node_modules/")[1].split("/")[0].toString();
            }
          }
        }
      }
    }
  }
})
