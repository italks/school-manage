import axios from 'axios'

const service = axios.create({
  baseURL: '/api', // Proxy will handle this or set full URL
  timeout: 5000
})

// Request interceptor
service.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// Response interceptor
service.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error) => {
    console.error('Request Error:', error)
    return Promise.reject(error)
  }
)

export default service
