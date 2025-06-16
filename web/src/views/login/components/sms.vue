<script setup lang="ts">
import { ref } from 'vue'
import router from "@/router";

const count = ref(0)
const timer = ref<any>(null)
const data = ref({
  phone: '',
  code: '',
  autoLogin: false
})

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
  <div class="login-container">
    <div class="login-form">
      <n-input size="large" v-model:value="data.phone" placeholder="手机号" style="margin-top: 16px;"/>
      <n-input-group style="margin-top: 16px;">
        <n-input size="large" v-model:value="data.code" placeholder="验证码"/>
        <n-button type="primary" size="large" ghost @click="getSms">
          {{ count > 0 ? `${count}秒` : '获取验证码' }}
        </n-button>
      </n-input-group>
      <n-button type="primary" size="large" :disabled="data.phone.length === 0 || data.code.length === 0" style="width: 100%;margin-top: 16px;">登录</n-button>
    </div>
    <div class="auto-login">
      <n-checkbox v-model:checked="data.autoLogin" style="margin-right: 12px">下次自动登录</n-checkbox>
      <n-space style="align-items: center;">
        <n-a @click="() => { router.push({ name:'recovery' }) }" style="color: #2c3e50;">忘记密码</n-a>
        <div style="font-size: 12px;">|</div>
        <n-a @click="() => { router.push({name: 'register'}) }" style="color: #2c3e50;">注册</n-a>
      </n-space>
    </div>
    <div class="login-tips">
      <div>登录或注册即代表同意 用户协议 和 个人信息保护政策</div>
      <div>未满18周岁未成年人请勿自行注册，其注册、登录账号及使用服务需征得监护人同意</div>
    </div>
  </div>
</template>

<style lang="less" scoped>
.login-container {
  display: flex;
  flex-direction: column;
  .login-form {
    margin-top: 8px;
    display: flex;
    flex-direction: column;
  }
  .auto-login {
    margin-top: 16px;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }
  .login-tips{
    display: flex;
    flex-direction: column;
    margin-top: 16px;
    font-size: 10px;
    color: rgb(153, 153, 153);
    div{
      margin-top: 4px;
    }
  }
}
</style>