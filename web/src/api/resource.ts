import request from '../utils/request'

export interface Classroom {
  id?: number
  name: string
  capacity: number
  equipment: string
  status: 'available' | 'maintenance' | 'closed'
}

export const getClassrooms = () => {
  return request.get<{ data: Classroom[] }>('/resources/classrooms')
}

export const createClassroom = (data: Classroom) => {
  return request.post('/resources/classrooms', data)
}

export const updateClassroom = (id: number, data: Classroom) => {
  return request.put(`/resources/classrooms/${id}`, data)
}

export const deleteClassroom = (id: number) => {
  return request.delete(`/resources/classrooms/${id}`)
}
