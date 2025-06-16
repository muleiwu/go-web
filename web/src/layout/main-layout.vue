<script setup lang="ts">
import {ref} from "vue";
import type {MenuDividerOption, MenuGroupOption, MenuOption} from "naive-ui/es/menu/src/interface";
import logo from '@/assets/logo.svg'
import {useRoute, useRouter} from "vue-router";

const router = useRouter()

const menuOptions = ref<MenuOption[] | MenuGroupOption[] | MenuDividerOption[]>([
  {label: '首页', key: 'index', show: true}
])
const activeKey = ref(null)

const handleUpdateValue = (key: string, item: MenuOption) => {
  if (typeof item.key === 'string') {
    router.push({ name: item.key });
  } else {
    console.error('Invalid route name:', item.key);
  }
}
</script>

<template>
  <main>
    <n-layout class="layout-main">
      <n-layout-header bordered class="layout-header">
        <div class="header-logo">
          <img :src="logo" width="100" alt="logo" />
        </div>
        <n-menu v-model:value="activeKey" mode="horizontal" :options="menuOptions" responsive @update:value="handleUpdateValue"/>
        <div class="nav-end">
<!--          <div style="display: flex;flex-direction: row;">-->
<!--            <n-button quaternary>登录</n-button>-->
<!--            <n-button strong secondary round>注册</n-button>-->
<!--          </div>-->
        </div>
      </n-layout-header>
      <n-layout-content bordered class="layout-content">
        <n-flex class="one-ms-container" align="center" justify="center">
          <div class="one-ms-content">
            <slot />
          </div>
        </n-flex>
      </n-layout-content>
      <n-layout-footer bordered class="layout-footer" >
        <n-space vertical>
          <div>Copyright &copy; 2024 合肥木雷坞信息技术有限公司. All rights reserved.</div>
          <div style="display: flex;align-items: center;justify-content: center;">
            <n-space vertical>
              <div>违法违规信息举报QQ：3094285305</div>
              <span>商务合作QQ：3094285305</span>
            </n-space>
          </div>
          <div><a href="http://beian.miit.gov.cn/" target="_blank" class="link-secondary" rel="noopener">皖ICP备2024056083号-1</a></div>
        </n-space>
      </n-layout-footer>
    </n-layout>
  </main>
</template>

<style scoped>

.layout-main {
  display: flex;
  flex-direction: column;
  width: 100%;
  .layout-header {
    padding: 0 32px;
    display: flex;
    flex-direction: row;
    height: 64px;
    align-items: center;
    .header-logo {
      width: 240px;
    }
  }

  .layout-content {
    min-height: calc(100vh - 128px);
    width: 100%;
    .one-ms-container {
      width: 100%;
      .one-ms-content {
        width: 100%;
        max-width: 1200px;
        padding: 16px 24px;
        box-sizing: border-box;
      }
    }
  }

  .layout-footer {
    padding: 32px;
    align-items: center;
    text-align: center;
  }

}
</style>
