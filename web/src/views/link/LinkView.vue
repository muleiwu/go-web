<script setup lang="ts">

import {useRoute} from "vue-router";

const route = useRoute()

const link = route.query.target as string

const click = () => {
  window.location.href = link
}
</script>

<template>
  <div class="link-background">
    <div class="link-middle-page">
      <n-card :segmented="{ content: true, footer: 'soft', }">
        <n-space vertical style="padding: 16px 8px 0;">
          <div style="font-size: 18px;line-height: 22px;">即将离开木雷坞，请注意账号财产安全</div>
          <div style="font-size: 14px;color: gray;">{{link}}</div>
        </n-space>
        <template #footer>
          <div style="display: flex;justify-content: end;">
            <n-button type="info" @click="click">继续访问</n-button>
          </div>
        </template>
      </n-card>
    </div>
  </div>

</template>

<style lang="less" scoped>
.link-background {
  font-size: 12px;
  background-color: #f4f5f5;
  height: 100vh;
  position: relative;
  .link-middle-page {
    width: 86%;
    max-width: 624px;
    position: absolute;
    top: 30%;
    left: 50%;
    transform: translateX(-50%);
  }
}
</style>