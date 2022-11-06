import { createApp } from "vue";
import { createPinia } from "pinia";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";

import App from "./App.vue";
import router from "./router";
import useStore from "@/stores/stores";

import "element-plus/theme-chalk/index.css";

const app = createApp(App);
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}

app.use(createPinia());
app.use(router);

app.mount("#app");

// router.beforeEach(async (to, from, next) => {
//   const { routerStore } = useStore();
//   const accessRoutes: any = await routerStore.generateRoutes();
//   console.log("router.beforeEach: ", accessRoutes);
//   accessRoutes.forEach((route: any) => {
//     router.addRoute(route);
//   });
//   console.log("router: ", router);
//   next();
// });
