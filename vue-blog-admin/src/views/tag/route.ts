import type { RouteType } from '~/types/router'

const Layout = () => import('@/layout/index.vue')

export default {
  name: 'tag',
  path: '/tag',
  component: Layout,
  redirect: '/tag/list',
  meta: {
    title: '标签管理',
    icon: 'tabler:tag',
    role: ['admin'],
    requireAuth: true,
    order: 1,
  },
  children: [
    {
      name: 'tagList',
      path: 'list',
      component: () => import('@/views/tag/list/index.vue'),
      meta: {
        title: '标签列表',
        icon: 'tabler:tag',
        role: ['admin'],
        requireAuth: true,
      },
    },
  ],
} as RouteType
