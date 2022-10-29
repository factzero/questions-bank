import request from "@/utils/request";

export const login = (data) => {
  return request({
    url: "/base/login",
    method: "post",
    data: data,
  });
};

/**
 * 获取图片验证码
 */
export const getCaptcha = () => {
  return request({
    url: "/base/captcha",
    method: "post",
  });
};
