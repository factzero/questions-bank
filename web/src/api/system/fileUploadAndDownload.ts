import request from "@/utils/request";

export const uploadFileApi = (data: any) => {
  return request({
    url: "/api/v1/fileUploadAndDownload/upload",
    method: "post",
    data: data,
  });
};
