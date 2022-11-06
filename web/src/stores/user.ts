import { defineStore } from "pinia";

import { login } from "@/api/login/user";
import { localStorage } from "@/utils/storage";
import router from "@/router";
import useRouterStore from "./router";

const useUserStore = defineStore("user", {
  actions: {
    LoginIn(loginData) {
      return new Promise((resolve, reject) => {
        login(loginData)
          .then((respone) => {
            localStorage.set("token", respone.data.token);
            const routerStore = useRouterStore();
            routerStore.generateRoutes().then((res) => {
              res.forEach((route: any) => {
                router.addRoute(route);
              });
              router.push({ name: "dashboard" });
            });
            resolve(respone.data.token);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
  },
});

export default useUserStore;
