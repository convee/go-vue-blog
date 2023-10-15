import { request } from '@/utils'

export default {
  getPages: (params = {}) => request.get('/page/list', { params }),
  deletePage: (id: string) => request.post(`/page/delete`, { id }),
}
