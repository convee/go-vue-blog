import { request } from '@/utils'

export default {
  getPages: (params = {}) => request.get('/page/list', { params }),
  getPageById: (id: any) => request.get(`/page/detail?id=${id}`),
  addPage: (data: any) => request.post('/page/add', data),
  updatePage: (data: any) => request.post(`/page/update`, data),
  deletePage: (id: string) => request.post(`/page/delete`, { id }),
}
