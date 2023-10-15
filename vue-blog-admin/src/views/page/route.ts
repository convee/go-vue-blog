import type { RouteType } from '~/types/router'

const Layout = () => import('@/layout/index.vue')

export default {
  name: 'Page',
  path: '/page',
  component: Layout,
  redirect: '/page/list',
  meta: {
    title: '页面管理',
    icon: 'iconoir:journal-page',
    role: ['admin'],
    requireAuth: true,
    order: 2,
  },
  children: [
    {
      name: 'PageWrite',
      path: 'write',
      component: () => import('@/views/page/write/index.vue'),
      meta: {
        title: '添加页面',
        icon: 'icon-park-outline:write',
        role: ['admin'],
        requireAuth: true,
      },
    },
    {
      name: 'PageEdit',
      path: 'write/:id',
      component: () => import('@/views/page/write/index.vue'),
      isHidden: true,
      meta: {
        title: '编辑页面',
        icon: 'icon-park-outline:write',
        role: ['admin'],
        requireAuth: true,
      },
    },
    {
      name: 'PageList',
      path: 'list',
      component: () => import('@/views/page/list/index.vue'),
      meta: {
        title: '页面列表',
        icon: 'material-symbols:format-list-bulleted',
        role: ['admin'],
        requireAuth: true,
      },
    },
  ],
} as RouteType
