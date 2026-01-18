import request from '../utils/request'

export interface User {
  id: number
  username: string
  role: string
}

export const getUsers = (params?: { role?: string }) => {
  return request.get<{ data: User[] }>('/users', { params })
}
