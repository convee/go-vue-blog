import {createApp} from "vue";

import "@/styles/style.css";
import 'uno.css'
import App from "@/App.vue";
import naive, {createDiscreteApi} from "naive-ui";
import {createPinia} from "pinia"; // 引入pinia
import axios from "axios";
import {router} from "@/router/index.js";
import {createStore} from "vuex";

// 服务端地址
axios.defaults.baseURL = "http://localhost:8080";
// 独立API
const {message, notification, dialog} = createDiscreteApi([
    "message",
    "dialog",
    "notification",
]);
//创建并挂载根实例
const app = createApp(App);
const store = createStore({});


// 全局提供属性
app.provide("axios", axios); //注册网络库
app.provide("message", message);
app.provide("notification", notification);
app.provide("dialog", dialog);
app.provide("server_url", axios.defaults.baseURL);

app.use(naive); // 注册 naive ui
app.use(createPinia()); // 注册使用 pinia
app.use(router); // 注册路由
app.use(store) // 将 store 实例作为插件安装

// router.beforeEach(async (to, from, next) => {
//     // 判断是否登录
//     let isLogin = await store.dispatch('validate')
//     let needLogin = to.matched.some(match => match.meta.needLogin)
//     if (needLogin) {
//         if (isLogin) {
//             next()
//         } else {
//             next("/login")
//         }
//     } else {
//         if (isLogin && to.path === "/login") {
//             next("/")
//         } else {
//             next()
//         }
//     }
// })

// axios拦截器
axios.interceptors.request.use((config) => {
    //每次请求都在headers中添加token
    config.headers.Authorization = "Bearer " + localStorage.getItem("token");
    return config;
});
app.mount("#app"); //挂载到app
