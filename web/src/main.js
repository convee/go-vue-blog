import { createApp } from "vue";

import "./style.css";
import App from "./App.vue";
import naive from "naive-ui";
import { createDiscreteApi } from "naive-ui";
import { createPinia } from "pinia"; // 引入pinia
import axios from "axios";
import { router } from "./router.js";
import { AdminStore } from "./stores/AdminStore";

// 服务端地址
axios.defaults.baseURL = "http://localhost:8000";
// 独立API
const { message, notification, dialog } = createDiscreteApi([
  "message",
  "dialog",
  "notification",
]);
//创建并挂载根实例
const app = createApp(App);

// 全局提供属性
app.provide("axios", axios); //注册网络库
app.provide("message", message);
app.provide("notification", notification);
app.provide("dialog", dialog);
app.provide("server_url", axios.defaults.baseURL);

app.use(naive); //注册 naive ui
app.use(createPinia()); //注册使用pinia
app.use(router); //注册路由

const adminStore = AdminStore();
// axios拦截器
axios.interceptors.request.use((config) => {
  //每次请求都在headers中添加token
  config.headers.Authorization = "Bearer " + adminStore.token;
  return config;
});
app.mount("#app"); //挂载到app
