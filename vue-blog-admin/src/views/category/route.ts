import type { RouteType } from '~/types/router'

const Layout = () => import('@/layout/index.vue')

export default {
  name: 'category',
  path: '/category',
  component: Layout,
  redirect: '/category/list',
  meta: {
    title: '分类管理',
    icon: 'tabler:category',
    role: ['admin'],
    requireAuth: true,
    order: 1,
  },
  children: [
    {
      name: 'categoryList',
      path: 'list',
      component: () => import('@/views/category/list/index.vue'),
      meta: {
        title: '分类列表',
        icon: 'tabler:category',
        role: ['admin'],
        requireAuth: true,
      },
    },
  ],
} as RouteType
