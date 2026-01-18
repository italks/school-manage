<template>
  <div class="bg-white p-6 rounded-lg shadow">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-bold">教室管理</h2>
      <el-button type="primary" @click="handleAdd">新增教室</el-button>
    </div>

    <el-table :data="tableData" v-loading="loading" style="width: 100%">
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="name" label="教室名称" />
      <el-table-column prop="capacity" label="容纳人数" width="120" />
      <el-table-column prop="equipment" label="设备" />
      <el-table-column prop="status" label="状态" width="120">
        <template #default="{ row }">
          <el-tag :type="getStatusType(row.status)">{{ getStatusLabel(row.status) }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180">
        <template #default="{ row }">
          <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
          <el-button link type="danger" @click="handleDelete(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- Dialog -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogType === 'add' ? '新增教室' : '编辑教室'"
      width="500px"
    >
      <el-form :model="form" label-width="80px">
        <el-form-item label="教室名称" required>
          <el-input v-model="form.name" />
        </el-form-item>
        <el-form-item label="容纳人数" required>
          <el-input-number v-model="form.capacity" :min="1" />
        </el-form-item>
        <el-form-item label="设备">
          <el-input v-model="form.equipment" placeholder="如：投影仪,音响" />
        </el-form-item>
        <el-form-item label="状态" v-if="dialogType === 'edit'">
          <el-select v-model="form.status">
            <el-option label="可用" value="available" />
            <el-option label="维护中" value="maintenance" />
            <el-option label="关闭" value="closed" />
          </el-select>
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
import { getClassrooms, createClassroom, updateClassroom, deleteClassroom, type Classroom } from '../../api/resource'
import { ElMessage, ElMessageBox } from 'element-plus'

const loading = ref(false)
const tableData = ref<Classroom[]>([])
const dialogVisible = ref(false)
const dialogType = ref<'add' | 'edit'>('add')

const form = reactive<Classroom>({
  name: '',
  capacity: 30,
  equipment: '',
  status: 'available'
})

const fetchList = async () => {
  loading.value = true
  try {
    const res = await getClassrooms()
    // Adjust based on actual API response structure (e.g. res.data.data or res.data)
    // Assuming res.data.data based on axios + backend response
    tableData.value = (res as any).data || []
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

const getStatusLabel = (status: string) => {
  const map: Record<string, string> = {
    available: '可用',
    maintenance: '维护中',
    closed: '关闭'
  }
  return map[status] || status
}

const getStatusType = (status: string) => {
  const map: Record<string, string> = {
    available: 'success',
    maintenance: 'warning',
    closed: 'info'
  }
  return map[status] || ''
}

const handleAdd = () => {
  dialogType.value = 'add'
  form.id = undefined
  form.name = ''
  form.capacity = 30
  form.equipment = ''
  form.status = 'available'
  dialogVisible.value = true
}

const handleEdit = (row: Classroom) => {
  dialogType.value = 'edit'
  Object.assign(form, row)
  dialogVisible.value = true
}

const handleDelete = (row: Classroom) => {
  ElMessageBox.confirm('确认删除该教室吗？', '提示', {
    type: 'warning'
  }).then(async () => {
    if (row.id) {
      await deleteClassroom(row.id)
      ElMessage.success('删除成功')
      fetchList()
    }
  })
}

const handleSubmit = async () => {
  try {
    if (dialogType.value === 'add') {
      await createClassroom(form)
      ElMessage.success('创建成功')
    } else {
      if (form.id) {
        await updateClassroom(form.id, form)
        ElMessage.success('更新成功')
      }
    }
    dialogVisible.value = false
    fetchList()
  } catch (error) {
    console.error(error)
  }
}

onMounted(() => {
  fetchList()
})
</script>
