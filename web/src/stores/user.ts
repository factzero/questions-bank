import { defineStore } from "pinia";

import { login, getUserInfo } from "@/api/login/user";
import { localStorage } from "@/utils/storage";

const useUserStore = defineStore("user", {
  state: () => ({
    token: localStorage.get("token") || "",
    flag: "",
    userInfo: {
      uuid: "",
      nickName: "",
      headerImg: "",
      authority: {},
      sideMode: "dark",
      activeColor: "#4D70FF",
      baseColor: "#fff",
    },
  }),
  actions: {
    setUserInfo(val: any) {
      const {
        uuid,
        nickName,
        headerImg,
        authority,
        sideMode,
        activeColor,
        baseColor,
      } = val;
      this.userInfo.uuid = uuid;
      this.userInfo.nickName = nickName;
      this.userInfo.headerImg = headerImg;
      this.userInfo.authority = authority;
      this.userInfo.sideMode = sideMode;
      this.userInfo.activeColor = activeColor;
      this.userInfo.baseColor = baseColor;
    },
    LogIn(loginData: any) {
      return new Promise((resolve, reject) => {
        login(loginData)
          .then((respone) => {
            localStorage.set("token", respone.data.token);
            this.token = respone.data.token;
            this.setUserInfo(respone.data.user);
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
    /* 获取用户信息 */
    GetUserInfo() {
      return new Promise((resolve, reject) => {
        getUserInfo()
          .then((respone) => {
            this.setUserInfo(respone.data.userInfo);
            resolve(respone.data.userInfo);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    /* 清理数据 */
    ClearStorage() {
      this.token = "";
      sessionStorage.clear();
      localStorage.clear();
    },
  },
});

export default useUserStore;
