<template>
  <el-container class="h-screen">
    <el-aside width="200px" class="bg-gray-800 text-white">
      <div class="h-14 flex items-center justify-center text-xl font-bold border-b border-gray-700">
        学校管理系统
      </div>
      <el-menu
        router
        :default-active="$route.path"
        background-color="#1f2937"
        text-color="#fff"
        active-text-color="#409eff"
        class="border-r-0"
      >
        <el-menu-item index="/dashboard">
          <el-icon><Odometer /></el-icon>
          <span>仪表盘</span>
        </el-menu-item>
        <el-sub-menu index="/resources">
          <template #title>
            <el-icon><School /></el-icon>
            <span>资源管理</span>
          </template>
          <el-menu-item index="/resources/classrooms">教室管理</el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header class="bg-white shadow-sm flex items-center justify-between px-4">
        <div class="text-lg font-medium">{{ $route.meta.title || '首页' }}</div>
        <el-dropdown @command="handleCommand">
          <span class="cursor-pointer flex items-center">
            管理员
            <el-icon class="ml-1"><ArrowDown /></el-icon>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="logout">退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-header>
      <el-main class="bg-gray-50">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { Odometer, School, ArrowDown, Calendar } from '@element-plus/icons-vue'

const router = useRouter()

const handleCommand = (command: string) => {
  if (command === 'logout') {
    localStorage.removeItem('token')
    localStorage.removeItem('userInfo')
    router.push('/login')
  }
}
</script>
