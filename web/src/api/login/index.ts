import request from '@/utils/request';


/**
 * 获取图片验证码
 */
 export function getCaptcha() {
    return request({
      url: '/captcha?t=' + new Date().getTime().toString(),
      method: 'get',
    });
  }