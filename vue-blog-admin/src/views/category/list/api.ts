import { request } from '@/utils'

export default {
  getCategory: (params = {}) => request.get('/category/list', { params }),
  getCategoryById: (id: string) => request.get(`/category/detail?id=${id}`),
  addCategory: (data: any) => request.post('/category/add',  data),
  updateCategory: (data: any) => request.post(`/category/update`, data),
  deleteCategory: (id: string) => request.post(`/category/delete`, { id }),
}
