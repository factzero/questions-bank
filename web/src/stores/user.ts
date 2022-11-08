import { defineStore } from "pinia";

import { login } from "@/api/login/user";
import { localStorage } from "@/utils/storage";

const useUserStore = defineStore("user", {
  state: () => ({
    token: localStorage.get("token") || "",
    flag: "",
  }),
  actions: {
    LogIn(loginData) {
      return new Promise((resolve, reject) => {
        login(loginData)
          .then((respone) => {
            localStorage.set("token", respone.data.token);
            this.token = respone.data.token;
            resolve(respone.data.token);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    LogOut() {
      return new Promise((resolve, reject) => {
        localStorage.remove("token");
      });
    },
  },
});

export default useUserStore;
