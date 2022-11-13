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

// @Summary 用户注册
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/resige [post]
export const register = (data) => {
  return request({
    url: "/api/v1/user/admin_register",
    method: "post",
    data: data,
  });
};

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "设置用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserInfo [put]
export const setUserInfo = (data) => {
  return request({
    url: "/api/v1/user/setUserInfo",
    method: "put",
    data: data,
  });
};

export const getUserInfo = () => {
  return request({
    url: "/api/v1/user/getUserInfo",
    method: "get",
  });
};

// @Tags User
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
export const getUserList = (data) => {
  return request({
    url: "/api/v1/user/getUserList",
    method: "post",
    data: data,
  });
};

// @Tags User
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.SetUserAuth true "设置用户权限"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
export const setUserAuthority = (data) => {
  return request({
    url: "/api/v1/user/setUserAuthority",
    method: "post",
    data: data,
  });
};

// @Tags User
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.setUserAuthorities true "设置用户权限"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthorities [post]
export const setUserAuthorities = (data) => {
  return request({
    url: "/api/v1/user/setUserAuthorities",
    method: "post",
    data: data,
  });
};

// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/deleteUser [delete]
export const deleteUser = (data) => {
  return request({
    url: "/api/v1/user/deleteUser",
    method: "delete",
    data: data,
  });
};

export const resetPassword = (data) => {
  return request({
    url: "/api/v1/user/resetPassword",
    method: "post",
    data: data,
  });
};
