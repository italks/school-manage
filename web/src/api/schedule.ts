import request from '../utils/request'
import type { Classroom } from './resource'
import type { User } from './user'

export interface Schedule {
  id?: number
  teacher_id: number
  classroom_id: number
  course_name: string
  start_time: string
  end_time: string
  max_students: number
  status: string
  teacher?: User
  classroom?: Classroom
}

export const getSchedules = () => {
  return request.get<{ data: Schedule[] }>('/schedules')
}

export const createSchedule = (data: Schedule) => {
  return request.post('/schedules', data)
}

export const deleteSchedule = (id: number) => {
  return request.delete(`/schedules/${id}`)
}
