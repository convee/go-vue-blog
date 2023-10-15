import axios from 'axios';
import { AdminStore } from "@/stores/AdminStore";
import { createDiscreteApi } from "naive-ui";
import { router } from ".";



const { message } = createDiscreteApi([
  "message",
]);

const BASE_URL = '/'; // 你的 API 基础路径，可以根据需求修改

const instance = axios.create({
  baseURL: import.meta.env.VITE_BASE_URL,
  timeout: 5000, // 请求超时时间，根据需求修改
});

// 请求拦截器，可以在发送请求之前做一些全局操作，如添加认证信息等
instance.interceptors.request.use(
  (config) => {
    // 在请求头中添加认证信息，如果需要的话
    config.headers["Authorization"] = `Bearer ${localStorage.getItem('token')}` || null;
    
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器，可以在接收响应之后做一些全局操作，如处理错误信息等
instance.interceptors.response.use(
  (response) => {
    // 在这里处理响应数据，根据需求做相应的操作
    // 获取res.code 如果是401则跳转到登录页面 并且清除本地存储中的token
    if (response.data.code === 401) {
      const adminStore = AdminStore();
      adminStore.delToken();
      message.error("登录失效，请重新登录");
      // 刷新页面
      router.go(0);

    }
    return response.data;
  },
  (error) => {
    // 在这里处理错误信息，根据需求做相应的操作
    return Promise.reject(error);
  }
);

// 导出封装后的 Axios 实例
export default instance;
