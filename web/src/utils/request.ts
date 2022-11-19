import axios from "axios";
import { ElMessage } from "element-plus";
import router from "@/router";
import useStore from "@/stores/stores";

// 创建 axios 实例
const service = axios.create({
  baseURL: import.meta.env.VITE_APP_BASE_API,
  timeout: 50000,
  headers: { "Content-Type": "application/json;charset=utf-8" },
});

// 添加请求拦截器
service.interceptors.request.use(
  function (config) {
    const { userStore } = useStore();
    var contentType = "application/json";
    if (config.url === "/api/v1/fileUploadAndDownload/upload") {
      contentType = "multipart/form-data";
    }
    config.headers = {
      ...config.headers,
      "Content-Type": contentType,
      "x-token": userStore.token,
      "x-user-id": userStore.userInfo.uuid,
    };
    return config;
  },
  function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
  }
);

// 添加响应拦截器
service.interceptors.response.use(
  function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    if (response.data.code === 0) {
      return response.data;
    } else {
      ElMessage({
        message: response.data.msg,
        type: "error",
      });
      const { userStore } = useStore();
      userStore.token = "";
      localStorage.clear();
      router.push({ name: "Login", replace: true });
      return Promise.reject(new Error(response.data.msg || "Error"));
    }
  },
  function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    return Promise.reject(error);
  }
);

// 导出 axios 实例
export default service;
