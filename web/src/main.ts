import { createApp } from "vue";
import { createPinia } from "pinia";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";
import ElementPlus from "element-plus";
import zhCn from "element-plus/es/locale/lang/zh-cn";

import App from "./App.vue";
import router from "./router";
import useStore from "@/stores/stores";

import "element-plus/theme-chalk/index.css";
import "@/style/element_visiable.scss";

const app = createApp(App);
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}

app.use(createPinia());
app.use(router);
app.use(ElementPlus, { locale: zhCn });

app.mount("#app");

// 白名单路由
const whiteList = ["/login"];

router.beforeEach(async (to, from, next) => {
  console.log("router.beforeEach: to ", to);
  const { userStore, routerStore } = useStore();
  const hasToken = userStore.token;
  if (hasToken) {
    const hasLogin = userStore.flag.length > 0;
    if (hasLogin) {
      next();
    } else {
      try {
        userStore.flag = "login";
        const accessRoutes: any = await routerStore.generateRoutes();
        accessRoutes.forEach((route: any) => {
          router.addRoute(route);
        });
        next({ ...to, replace: true });
      } catch (error) {
        console.log("beforeEach:", error);
      }
    }
  } else {
    if (whiteList.indexOf(to.path) !== -1) {
      next();
    } else {
      // next({ path: "/login" });
      next(`/login?redirect=${to.path}`);
    }
  }
});
