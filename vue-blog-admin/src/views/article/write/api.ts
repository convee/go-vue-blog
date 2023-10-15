import { request } from '@/utils'

export default {
  getArticles: (params = {}) => request.get('/article/list', { params }),
  getArticleById: (id: any) => request.get(`/article/detail?id=${id}`),
  addArticle: (data: any) => request.post('/article/add', data),
  updateArticle: (data: any) => request.post(`/article/update`, data),
  deleteArticle: (id: string) => request.post(`/article/delete`, { id }),
  getCategoryOption: () => request.get('/category/all'),
  getTagOption: () => request.get('/tag/all'),
}
