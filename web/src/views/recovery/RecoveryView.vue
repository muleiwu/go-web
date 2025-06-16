<script setup lang="ts">
import {ref} from "vue";
import router from "@/router";

const data = ref({
  account: '',
  code: ''
})

const count = ref(0)
const timer = ref<any>(null)

const getSms = () => {
  if (count.value > 0) {
    return
  }
  count.value = 60
  // 倒计时
  timer.value = setInterval(() => {
    count.value--
    if (count.value <= 0) {
      clearInterval(timer.value)
    }
  }, 1000)
}
</script>

<template>
  <div class="app">
    <n-space class="login-card" vertical>
      <h1 style="text-align: center;color: #ffffff;font-weight: bold;">木雷坞 · 通行证</h1>
      <n-card>
        <template #header>
          <div style="font-size: 22px;text-align: center;">找回密码</div>
        </template>
        <n-space vertical class="recovery-from">
          <n-input v-model="data.account" placeholder="手机号/邮箱"></n-input>
          <n-input-group>
            <n-input size="large" v-model:value="data.code" placeholder="验证码" />
            <n-button type="primary" size="large" ghost @click="getSms">
              {{ count > 0 ? `${count}秒` : '获取验证码' }}
            </n-button>
          </n-input-group>
          <n-button type="primary" style="width: 100%;margin-top: 16px;">下一步</n-button>
        </n-space>
        <div class="auto-login">
          <n-space style="align-items: center;">
            已有账号，<n-a @click="() => { router.push({ name:'login' }) }">马上登录 </n-a>
          </n-space>
        </div>
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
    .recovery-from {
      margin-top: 8px;
      display: flex;
      flex-direction: column;
    }
    .auto-login{
      margin-top: 16px;
      display: flex;
      flex-direction: row;
      justify-content: end;
    }
  }
}
</style>