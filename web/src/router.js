import { createRouter, createWebHistory } from "vue-router"; //引入路由

// 路由配置
let routes = [
  //每个路由都需要映射到一个组件。
  { path: "/login", component: () => import("./views/Login.vue") },
  { path: "/", component: () => import("./views/front/Home.vue") },
  { path: "/detail", component: () => import("./views/front/Detail.vue") },

  {
    //嵌套路由
    path: "/main",
    component: () => import("./views/Main.vue"),
    children: [
      { path: "/main/page1", component: () => import("./views/Page1.vue") },
      { path: "/main/page2", component: () => import("./views/Page2.vue") },
    ],
  },
];
const router = createRouter({
  history: createWebHistory(),
  routes,
});

export { router, routes };
