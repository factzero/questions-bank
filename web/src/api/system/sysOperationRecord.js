import request from "@/utils/request";

export const deleteSysOperationRecord = (data) => {
  return request({
    url: "/api/v1/sysOperationRecord/deleteSysOperationRecord",
    method: "delete",
    data,
  });
};

export const deleteSysOperationRecordByIds = (data) => {
  return request({
    url: "/api/v1/sysOperationRecord/deleteSysOperationRecordByIds",
    method: "delete",
    data,
  });
};

export const getSysOperationRecordList = (params) => {
  return request({
    url: "/api/v1/sysOperationRecord/getSysOperationRecordList",
    method: "get",
    params,
  });
};
