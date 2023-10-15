import type { RouteType } from '~/types/router'

const Layout = () => import('@/layout/index.vue')

export default {
  name: 'Article',
  path: '/article',
  component: Layout,
  redirect: '/article/list',
  meta: {
    title: '文章管理',
    icon: 'ic:twotone-article',
    role: ['admin'],
    requireAuth: true,
    order: 1,
  },
  children: [
    {
      name: 'ArticleWrite',
      path: 'write',
      component: () => import('@/views/article/write/index.vue'),
      meta: {
        title: '发布文章',
        icon: 'icon-park-outline:write',
        role: ['admin'],
        requireAuth: true,
      },
    },
    {
      name: 'ArticleEdit',
      path: 'write/:id',
      component: () => import('@/views/article/write/index.vue'),
      isHidden: true,
      meta: {
        title: '编辑文章',
        icon: 'icon-park-outline:write',
        role: ['admin'],
        requireAuth: true,
      },
    },
    {
      name: 'ArticleList',
      path: 'list',
      component: () => import('@/views/article/list/index.vue'),
      meta: {
        title: '文章列表',
        icon: 'material-symbols:format-list-bulleted',
        role: ['admin'],
        requireAuth: true,
      },
    },
  ],
} as RouteType
