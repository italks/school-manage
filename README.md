# School Management System

## 项目简介
基于 Go (Gin) + Vue 3 (Element Plus) 的学校管理系统。

## 功能模块
- **仪表盘**: 系统概览
- **资源管理**: 教室管理 (CRUD)
- **排课管理**: 课程安排、冲突检测 (老师/教室时间冲突)
- **用户管理**: 老师/学生管理

## 技术栈
- **后端**: Go 1.20+, Gin, GORM, MySQL, Redis
- **前端**: Vue 3, TypeScript, Vite, Element Plus, Tailwind CSS
- **部署**: Docker Compose

## 快速开始

### 1. 启动基础设施 (MySQL & Redis)
```bash
docker-compose up -d
```

### 2. 启动后端
```bash
cd api
go mod tidy
go run main.go
```
后端服务运行在 `http://localhost:8080`

### 3. 启动前端
```bash
cd web
pnpm install
pnpm dev
```
前端服务运行在 `http://localhost:5173`

## 默认账号
- 注册一个新账号即可使用。
