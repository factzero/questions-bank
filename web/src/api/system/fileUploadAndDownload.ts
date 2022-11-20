import request from "@/utils/request";

export const uploadFileApi = (data: any) => {
  return request({
    url: "/api/v1/fileUploadAndDownload/upload",
    method: "post",
    data: data,
  });
};

export const getFileList = (data: any) => {
  return request({
    url: "/api/v1/fileUploadAndDownload/getFileList",
    method: "post",
    data: data,
  });
};

export const deleteFile = (data: any) => {
  return request({
    url: "/api/v1/fileUploadAndDownload/deleteFile",
    method: "post",
    data: data,
  });
};

export const editFileName = (data: any) => {
  return request({
    url: "/api/v1/fileUploadAndDownload/editFileName",
    method: "post",
    data: data,
  });
};
