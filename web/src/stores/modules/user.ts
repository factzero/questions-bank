import { defineStore } from "pinia";
import type { LoginFormData } from '@/types/api/system/login';

const useUserStore = defineStore("user", {
  actions: {
    login(loginData: LoginFormData) {
      console.log("pinia useUserStore login():", loginData)
      const { username, password, code, uuid } = loginData;
      return new Promise((resolve, reject) => {
        if(username === "admin") {
          resolve(123)
        } else {
          reject(new Error("Something awful happened"));
        }
      });
    },
  },
});

export default useUserStore;