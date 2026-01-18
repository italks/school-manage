<template>
  <div class="p-4">
    <h1 class="text-2xl font-bold">仪表盘</h1>
    <p class="mt-4">欢迎进入学校管理系统</p>
    <el-button @click="logout" type="danger" class="mt-4">退出登录</el-button>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('userInfo')
  router.push('/login')
}
</script>
