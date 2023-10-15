import { request } from '@/utils'

export default {
  getArticles: (params = {}) => request.get('/article/list', { params }),
  getArticleById: (id: string) => request.get(`/article/detail?id=${id}`),
  addArticle: (params = {}) => request.post('/article/add', { params }),
  updateArticle: (params = {}) => request.post(`/article/update`, { params }),
  deleteArticle: (id: string) => request.post(`/article/delete`, { id }),
}
