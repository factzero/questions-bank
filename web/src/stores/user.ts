import { defineStore } from "pinia";

import { login } from "@/api/login/user";
import { localStorage } from "@/utils/storage";

const useUserStore = defineStore("user", {
  state: () => ({
    token: localStorage.get("token") || "",
    flag: "",
    userName: "",
    uuid: "",
    nickName: "",
    headerImg: "",
    authority: {},
    sideMode: "dark",
    activeColor: "#4D70FF",
    baseColor: "#fff",
  }),
  actions: {
    LogIn(loginData) {
      return new Promise((resolve, reject) => {
        login(loginData)
          .then((respone) => {
            localStorage.set("token", respone.data.token);
            this.token = respone.data.token;
            this.userName = loginData.username;
            resolve(respone.data.token);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    LogOut() {
      return new Promise((resolve, reject) => {
        this.token = "";
        sessionStorage.clear();
        localStorage.clear();
        window.location.reload();
      });
    },
  },
});

export default useUserStore;
