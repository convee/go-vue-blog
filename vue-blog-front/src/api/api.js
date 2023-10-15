import { defAxios as request } from '@/utils/http'

// 获取文章列表
export function getArticles(page) {
  return request({
    url: "/backend/article/list",
    method: "get",
    params: data,
  });
}



const loadArticles = async (page = 0) => {
  if (page !== 0) {
    pageInfo.page = page;
  }
  let res = await axios.get(`/api/article/list?keyword=${pageInfo.keyword}&page=${pageInfo.page}&pageSize=${pageInfo.pageSize}&categoryId=${pageInfo.categoryId}`)
  console.log(res)
  let temp_rows = res.data.data.data;
  // 处理获取的文章列表数据
  for (let row of temp_rows) {
    row.content += "..."
    // 把时间戳转换为年月日
    let d = new Date(row.createdAt)
    row.create_time = `${d.getFullYear()}年${d.getMonth() + 1}月${d.getDate()}日`
  }
  articleListInfo.value = temp_rows;
  pageInfo.total = res.data.data.total;
  //计算分页大小
  pageInfo.totalPage = res.data.data.totalPage
  console.log(res)
}

// 获取文章详情
export function getArticleDetail(id) {
  return request({
    url: `/backend/article/${id}`,
    method: 'get',
  })
}
// 获取分类列表
export function getCategoryList() {
  return request({
    url: '/api/categories/',
    method: 'get',
  })
}

// 登录
export function userLogin(data) {
  return request({
    url: '/users/login/',
    method: 'post',
    data,
  })
}

// 注册
export function userRegister(data) {
  return request({
    url: '/users/register/',
    method: 'post',
    data,
  })
}

// 新增文章
export function addArticle(data) {
  return request({
    url: '/articles/token/',
    method: 'post',
    data,
  })
}

// 根据id获取文章/articles/1
export function getArticleById(id) {
  return request({
    url: `/articles/${id}`,
    method: 'get',
  })
}

// 修改文章/articles/token/16
export function updateArticleById(id, data) {
  return request({
    url: `/articles/token/${id}`,
    method: 'put',
    data,
  })
}

// 删除文章/articles/token/2
export function deleteArticleById(id) {
  return request({
    url: `/articles/token/${id}`,
    method: 'delete',
  })
}

// 根据id获取分类
export function getCategoryById(id) {
  return request({
    url: `/categories/${id}`,
    method: 'get',
  })
}

// 新增分类
export function addCategory(data) {
  return request({
    url: '/categories/token/',
    method: 'post',
    data,
  })
}

// 修改分类
export function updateCategoryById(id, data) {
  return request({
    url: `/categories/token/${id}`,
    method: 'put',
    data,
  })
}

// 删除分类
export function deleteCategoryById(id) {
  return request({
    url: `/categories/token/${id}`,
    method: 'delete',
  })
}
