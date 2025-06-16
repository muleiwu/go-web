<script setup lang="ts">
import './assets/main.css'
import {RouterView} from 'vue-router'
import hljs from 'highlight.js'
import {ref} from "vue";

hljs.initHighlightingOnLoad();
const theme = ref(null)
</script>

<template>
  <n-config-provider :theme="theme" :hljs="hljs">
    <RouterView/>
  </n-config-provider>
</template>