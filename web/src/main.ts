import { createApp } from "vue";
import { createPinia } from "pinia";

import 'element-plus/theme-chalk/index.css';

import App from "./App.vue";
import router from "./router";

// 引入svg注册脚本
import "virtual:svg-icons-register";

// 自定义样式
import "@/styles/index.scss";

const app = createApp(App);

app.use(createPinia());
app.use(router);

app.mount("#app");
