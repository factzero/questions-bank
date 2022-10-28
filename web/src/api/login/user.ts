import request from "@/utils/request";

/**
 * 获取图片验证码
 */
export function getCaptcha() {
  return request({
    url: "/base/captcha",
    method: "post",
  });
}
