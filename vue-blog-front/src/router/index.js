import { createRouter, createWebHistory } from "vue-router"; //引入路由

// 路由配置
let routes = [
  // 首页
  { path: "/", component: () => import("@/views/home/index.vue") },
  // 文章详情页
  { path: "/detail/:id", component: () => import("@/views/detail/index.vue") },
  // 标签页
  { path: "/tags", component: () => import("@/views/tags/index.vue") },
  // 自定义页面
  { path: "/page/:ident", component: () => import("@/views/page/index.vue") },
];
const router = createRouter({
  history: createWebHistory(),
  routes,
});

export { router, routes };
