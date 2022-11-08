<template>
  <div id="userLayout">
    <div class="login_panel">
      <div class="login_panel_form">
        <div class="login_panel_form_title">
          <p class="login_panel_form_title_p">试题管理系统</p>
        </div>
        <el-form
          ref="loginForm"
          :model="loginFormData"
          :rules="rules"
          :validate-on-rule-change="false"
          @keyup.enter="submitForm"
        >
          <el-form-item prop="username">
            <el-input
              v-model="loginFormData.username"
              placeholder="请输入用户名"
            >
              <template #suffix>
                <span class="input-icon">
                  <el-icon><user /></el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="loginFormData.password"
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
            >
              <template #suffix>
                <span class="input-icon">
                  <el-icon>
                    <component :is="lock" @click="changeLock" />
                  </el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="captcha">
            <div class="vPicBox">
              <el-input
                v-model="loginFormData.captcha"
                placeholder="请输入验证码"
                style="width: 60%"
              />
              <div class="vPic">
                <img
                  v-if="picPath"
                  :src="picPath"
                  alt="请输入验证码"
                  @click="loginVerify()"
                />
              </div>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              size="large"
              style="width: 100%"
              @click="submitForm"
              >登 录</el-button
            >
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import { ElMessage } from "element-plus";
import { getCaptcha } from "@/api/login/user";

// 状态管理
import useUserStore from "@/stores/user";

import router from "@/router";

// 控制是否显示密码
const lock = ref("lock");
const changeLock = () => {
  lock.value = lock.value === "lock" ? "unlock" : "lock";
};

// 用户名和密码验证函数
const checkUsername = (rule, value, callback) => {
  if (value.length < 5) {
    return callback(new Error("请输入正确的用户名"));
  } else {
    callback();
  }
};
const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error("请输入正确的密码"));
  } else {
    callback();
  }
};

// 获取验证码
const loginVerify = () => {
  getCaptcha().then((ele) => {
    console.log(ele);
    picPath.value = ele.data.picPath;
    loginFormData.captchaId = ele.data.captchaId;
  });
};
loginVerify();

const picPath = ref("");
const loginForm = ref(null);
const loginFormData = reactive({
  username: "admin",
  password: "123456",
  captcha: "",
  captchaId: "",
});
const rules = reactive({
  username: [{ validator: checkUsername, trigger: "blur" }],
  password: [{ validator: checkPassword, trigger: "blur" }],
  captcha: [
    {
      message: "验证码格式不正确",
      trigger: "blur",
    },
  ],
});

const userStore = useUserStore();
const submitForm = () => {
  loginForm.value.validate((v) => {
    if (v) {
      userStore
        .LogIn(loginFormData)
        .then(() => {
          router.push({ path: "/layout/dashboard" });
        })
        .catch(() => {
          loginVerify();
        });
    } else {
      ElMessage({
        type: "error",
        message: "请正确填写登录信息",
        showClose: true,
      });
      loginVerify();
      return false;
    }
  });
};
</script>

<style scoped lang="scss">
@import "@/style/newLogin.scss";
</style>
