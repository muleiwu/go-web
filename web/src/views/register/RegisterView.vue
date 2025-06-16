<script setup lang="ts">
import From from "@/views/register/components/from.vue";
</script>

<template>
  <div class="app">
    <n-space class="login-card" vertical>
      <h1 style="text-align: center;color: #ffffff;font-weight: bold;">木雷坞 · 通行证</h1>
      <n-card>
        <template #header>
          <div style="font-size: 22px;text-align: center;">注册</div>
        </template>
        <from />
      </n-card>
    </n-space>
  </div>
</template>

<style lang="less" scoped>
.app {
  height: 100vh;
  width: 100%;
  background: url("https://image.hepeichun.com/background/zip/1.jpg") no-repeat center center fixed;
  background-size: cover;
  position: relative;
  .login-card {
    position: absolute;
    top: 20%;
    left: 50%;
    width: 80%;
    min-width: 300px;
    max-width: 400px;
    transform: translateX(-50%);
  }
}
</style>