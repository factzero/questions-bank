import { defineStore } from "pinia";
import { getMenu } from "@/api/system/menu";
import { asyncRouterHandle } from "@/utils/asyncRouter";

const baseRouter = [
  {
    path: "/layout",
    name: "layout",
    component: "view/layout/layout.vue",
    meta: {
      title: "底层layout",
    },
    children: [],
  },
];

const useRouterStore = defineStore("router", {
  state: () => ({
    baseRouter,
  }),
  actions: {
    generateRoutes() {
      return new Promise((resolve, reject) => {
        getMenu()
          .then((respone) => {
            baseRouter[0].children = respone.data.menus;
            console.log("1:", baseRouter);
            asyncRouterHandle(baseRouter);
            // console.log("2:", baseRouter);
            resolve(baseRouter);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
  },
});

export default useRouterStore;
