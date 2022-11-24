<template>
  <div class="gva-table-box">
    <el-upload
      action="#"
      accept=".pdf, .excel"
      :http-request="uploadFile"
      :on-exceed="handleExceed"
      :show-file-list="false"
    >
      <el-button type="primary">Click to upload</el-button>
      <template #tip>
        <div class="el-upload__tip">仅支持excel、pdf文件</div>
      </template>
    </el-upload>

    <el-table :data="tableData">
      <el-table-column align="left" label="日期" prop="UpdatedAt" width="180">
        <template #default="scope">
          <div>{{ formatDate(scope.row.UpdatedAt) }}</div>
        </template>
      </el-table-column>
      <el-table-column align="left" label="文件名/备注" prop="name" width="180">
        <template #default="scope">
          <div>{{ scope.row.name }}</div>
        </template>
      </el-table-column>
      <el-table-column align="left" label="链接" prop="url" min-width="300" />
      <el-table-column align="left" label="标签" prop="tag" width="100">
        <template #default="scope">
          <el-tag
            :type="scope.row.tag === 'jpg' ? 'primary' : 'success'"
            disable-transitions
            >{{ scope.row.tag }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column align="left" label="操作" width="160">
        <template #default="scope">
          <el-button
            size="small"
            icon="download"
            type="primary"
            link
            @click="downloadFile(scope.row)"
            >下载</el-button
          >
          <el-button
            size="small"
            icon="delete"
            type="primary"
            link
            @click="deleteFileFunc(scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <div class="gva-pagination">
      <el-pagination
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10, 30, 50, 100]"
        :style="{ float: 'right', padding: '20px' }"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import type { UploadProps } from "element-plus";

import {
  uploadFileApi,
  getFileList,
  deleteFile,
} from "@/api/system/fileUploadAndDownload";
import { formatDate } from "@/utils/format";

const uploadFile = (params: any) => {
  const file = params.file;
  var form = new FormData();
  form.append("file", file);
  uploadFileApi(form)
    .then((res) => {
      getTableData();
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

const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const search = ref({});
const tableData = ref([]);

// 查询
const getTableData = async () => {
  const table = await getFileList({
    page: page.value,
    pageSize: pageSize.value,
    ...search.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};
getTableData();

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

const deleteFileFunc = async (row) => {
  ElMessageBox.confirm("此操作将永久删除文件, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      const res = await deleteFile(row);
      if (res.code === 0) {
        ElMessage({
          type: "success",
          message: "删除成功!",
        });
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--;
        }
        getTableData();
      }
    })
    .catch(() => {
      ElMessage({
        type: "info",
        message: "已取消删除",
      });
    });
};
</script>
