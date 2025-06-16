<script setup lang="ts">
import Password from "@/views/login/components/password.vue";
import Sms from "@/views/login/components/sms.vue";
</script>

<template>
  <div class="app">
    <n-space class="login-card" vertical>
      <h1 style="text-align: center;color: #ffffff;font-weight: bold;">木雷坞 · 通行证</h1>
      <n-card>
        <n-space vertical>
          <n-tabs type="line" animated>
            <n-tab-pane name="oasis" tab="短信登录">
              <sms />
            </n-tab-pane>
            <n-tab-pane name="the beatles" tab="密码登录">
              <password />
            </n-tab-pane>
          </n-tabs>
        </n-space>
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
    top: 25%;
    left: 50%;
    width: 80%;
    min-width: 300px;
    max-width: 400px;
    transform: translateX(-50%);
  }
}
</style>