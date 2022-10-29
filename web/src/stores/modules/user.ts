import { defineStore } from "pinia";
import type { LoginFormData } from "@/types/api/system/login";
import { login } from "@/api/login/user";
import { localStorage } from "@/utils/storage";

const useUserStore = defineStore("user", {
  actions: {
    LoginIn(loginData: LoginFormData) {
      return new Promise((resolve, reject) => {
        login(loginData)
          .then((respone) => {
            localStorage.set("token", respone.data.token);
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