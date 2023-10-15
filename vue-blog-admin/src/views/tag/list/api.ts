import { request } from '@/utils'

export default {
  getTag: (params = {}) => request.get('/tag/list', { params }),
  getTagById: (id: string) => request.get(`/tag/detail?id=${id}`),
  addTag: (data: any) => request.post('/tag/add',  data),
  updateTag: (data: any) => request.post(`/tag/update`, data),
  deleteTag: (id: string) => request.post(`/tag/delete`, { id }),
}
