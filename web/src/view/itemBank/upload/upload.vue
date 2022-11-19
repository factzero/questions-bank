<template>
  <div class="gva-table-box">
    <el-upload
      action="#"
      accept=".pdf, .excel"
      :http-request="uploadFile"
      :limit="3"
      :on-exceed="handleExceed"
    >
      <el-button type="primary">Click to upload</el-button>
      <template #tip>
        <div class="el-upload__tip">仅支持excel、pdf文件</div>
      </template>
    </el-upload>
  </div>
</template>

<script lang="ts" setup>
import { ElMessage } from "element-plus";
import type { UploadProps } from "element-plus";

import { uploadFileApi } from "@/api/system/fileUploadAndDownload";

const uploadFile = (params: any) => {
  const file = params.file;
  var form = new FormData();
  form.append("file", file);
  uploadFileApi(form)
    .then((res) => {
      console.log("uploadFileApi res ", res);
    })
    .catch((err) => {
      console.log("catch ", err);
    });
};

const handleExceed: UploadProps["onExceed"] = (files, uploadFiles) => {
  ElMessage.warning(
    `The limit is 3, you selected ${files.length} files this time, add up to ${
      files.length + uploadFiles.length
    } totally`
  );
};
</script>
