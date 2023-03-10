import { createRouter, createWebHistory } from "vue-router"; //引入路由

// 路由配置
let routes = [
  //每个路由都需要映射到一个组件。
  { path: "/", component: () => import("./views/front/Home.vue") },
  { path: "/login", component: () => import("./views/Login.vue") },
  { path: "/detail", component: () => import("./views/front/Detail.vue") },
  //嵌套路由
  {
    path: "/dashboard", component: () => import("./views/dashboard/Dashboard.vue"), children: [
      { path: "/dashboard/category", component: () => import("./views/dashboard/Category.vue") },
      { path: "/dashboard/article", component: () => import("./views/dashboard/Article.vue") },
    ]
  },
];
const router = createRouter({
  history: createWebHistory(),
  routes,
});

export { router, routes };
