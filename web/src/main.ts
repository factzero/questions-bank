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

/**
 * @description 导入加载进度条，防止首屏加载时间过长，用户等待
 *
 * */
import Nprogress from "nprogress";
import "nprogress/nprogress.css";
Nprogress.configure({ showSpinner: false, ease: "ease", speed: 500 });
Nprogress.start();

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
let asyncRouterFlag = 0;

const getRouter = async () => {
  const { userStore, routerStore } = useStore();
  const accessRoutes: any = await routerStore.generateRoutes();
  accessRoutes.forEach((route: any) => {
    router.addRoute(route);
  });
  await userStore.GetUserInfo();
};

router.beforeEach(async (to, from, next) => {
  Nprogress.start();
  console.log("router.beforeEach: to ", to);
  to.meta.matched = [...to.matched];
  const { userStore } = useStore();
  const hasToken = userStore.token;
  if (hasToken) {
    if (!asyncRouterFlag && whiteList.indexOf(from.name) < 0) {
      asyncRouterFlag++;
      await getRouter();
      if (userStore.token) {
        next({ ...to, replace: true });
      } else {
        return { path: "/login" };
      }
    } else {
      next();
    }
  } else {
    if (whiteList.indexOf(to.path) !== -1) {
      next();
    } else {
      next(`/login?redirect=${to.path}`);
    }
  }
});

router.afterEach(() => {
  // 路由加载完成后关闭进度条
  Nprogress.done();
});

router.onError(() => {
  // 路由发生错误后销毁进度条
  Nprogress.remove();
});
