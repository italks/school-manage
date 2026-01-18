<template>
  <div class="bg-white p-6 rounded-lg shadow">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-bold">排课管理</h2>
      <el-button type="primary" @click="handleAdd">新增排课</el-button>
    </div>

    <el-table :data="tableData" v-loading="loading" style="width: 100%">
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="course_name" label="课程名称" />
      <el-table-column label="老师" width="120">
        <template #default="{ row }">
          {{ row.teacher?.username }}
        </template>
      </el-table-column>
      <el-table-column label="教室" width="120">
        <template #default="{ row }">
          {{ row.classroom?.name }}
        </template>
      </el-table-column>
      <el-table-column label="上课时间" width="300">
        <template #default="{ row }">
          {{ formatTime(row.start_time) }} - {{ formatTime(row.end_time) }}
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态" width="100">
        <template #default="{ row }">
          <el-tag :type="row.status === 'scheduled' ? 'success' : 'info'">{{ row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="120">
        <template #default="{ row }">
          <el-button 
            v-if="row.status === 'scheduled'"
            link 
            type="danger" 
            @click="handleDelete(row)"
          >
            取消
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- Dialog -->
    <el-dialog
      v-model="dialogVisible"
      title="新增排课"
      width="500px"
    >
      <el-form :model="form" label-width="80px">
        <el-form-item label="课程名称" required>
          <el-input v-model="form.course_name" />
        </el-form-item>
        <el-form-item label="老师" required>
          <el-select v-model="form.teacher_id" placeholder="请选择老师">
            <el-option
              v-for="item in teachers"
              :key="item.id"
              :label="item.username"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="教室" required>
          <el-select v-model="form.classroom_id" placeholder="请选择教室">
            <el-option
              v-for="item in classrooms"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="开始时间" required>
          <el-date-picker
            v-model="form.start_time"
            type="datetime"
            placeholder="选择开始时间"
            value-format="YYYY-MM-DDTHH:mm:ssZ"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="结束时间" required>
          <el-date-picker
            v-model="form.end_time"
            type="datetime"
            placeholder="选择结束时间"
            value-format="YYYY-MM-DDTHH:mm:ssZ"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="满员人数">
          <el-input-number v-model="form.max_students" :min="1" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确认</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue'
import { getSchedules, createSchedule, deleteSchedule, type Schedule } from '../../api/schedule'
import { getClassrooms, type Classroom } from '../../api/resource'
import { getUsers, type User } from '../../api/user'
import { ElMessage, ElMessageBox } from 'element-plus'
import dayjs from 'dayjs'

const loading = ref(false)
const tableData = ref<Schedule[]>([])
const dialogVisible = ref(false)
const teachers = ref<User[]>([])
const classrooms = ref<Classroom[]>([])

const form = reactive<Schedule>({
  teacher_id: undefined as any,
  classroom_id: undefined as any,
  course_name: '',
  start_time: '',
  end_time: '',
  max_students: 20,
  status: 'scheduled'
})

const formatTime = (time: string) => {
  return dayjs(time).format('YYYY-MM-DD HH:mm')
}

const fetchData = async () => {
  loading.value = true
  try {
    const [schedulesRes, teachersRes, classroomsRes] = await Promise.all([
      getSchedules(),
      getUsers({ role: 'teacher' }),
      getClassrooms()
    ])
    
    tableData.value = (schedulesRes as any).data || []
    teachers.value = (teachersRes as any).data || []
    classrooms.value = (classroomsRes as any).data || []
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

const handleAdd = () => {
  form.teacher_id = undefined as any
  form.classroom_id = undefined as any
  form.course_name = ''
  form.start_time = ''
  form.end_time = ''
  form.max_students = 20
  dialogVisible.value = true
}

const handleDelete = (row: Schedule) => {
  ElMessageBox.confirm('确认取消该排课吗？', '提示', {
    type: 'warning'
  }).then(async () => {
    if (row.id) {
      await deleteSchedule(row.id)
      ElMessage.success('取消成功')
      fetchData()
    }
  })
}

const handleSubmit = async () => {
  try {
    // 简单验证
    if (!form.teacher_id || !form.classroom_id || !form.start_time || !form.end_time) {
      ElMessage.warning('请填写完整信息')
      return
    }
    
    await createSchedule(form)
    ElMessage.success('排课成功')
    dialogVisible.value = false
    fetchData()
  } catch (error: any) {
    if (error.response?.status === 409) {
      ElMessage.error('排课冲突：该时间段老师或教室已被占用')
    } else {
      ElMessage.error('操作失败')
    }
    console.error(error)
  }
}

onMounted(() => {
  fetchData()
})
</script>
