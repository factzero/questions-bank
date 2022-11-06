import request from "@/utils/request";

export const getMenu = () => {
  return request({
    url: "/api/v1/menu/getMenu",
    method: "post",
  });
};
