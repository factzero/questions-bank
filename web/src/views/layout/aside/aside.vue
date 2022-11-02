<template>
  <div :style="{ background: '#191a23' }">
    <el-scrollbar style="height: calc(100vh - 60px)">
      <transition
        :duration="{ enter: 800, leave: 100 }"
        mode="out-in"
        name="el-fade-in-linear"
      >
        <el-menu
          class="el-menu-vertical"
          background-color="theme.background"
          :text-color="theme.normalText"
          :active-text-color="theme.activeText"
        >
          <el-sub-menu index="1">
            <template #title>
              <el-icon><user /></el-icon>
              <span>超级管理员</span>
            </template>
            <el-menu-item index="1-1">
              <el-icon><coordinate /></el-icon>
              <template #title>用户管理</template>
            </el-menu-item>
            <el-menu-item index="1-1">
              <el-icon><pie-chart /></el-icon>
              <template #title>操作历史</template>
            </el-menu-item>
          </el-sub-menu>
          <el-menu-item index="2">
            <el-icon><cloudy /></el-icon>
            <template #title>服务器状态</template>
          </el-menu-item>
        </el-menu>
      </transition>
    </el-scrollbar>
  </div>
</template>

<script>
export default {
  name: "Aside",
};
</script>

<script setup>
import { ref } from "vue";

const theme = ref({});
const getTheme = () => {
  switch ("#191a23") {
    case "#fff":
      theme.value = {
        background: "#fff",
        activeBackground: "#4D70FF",
        activeText: "#fff",
        normalText: "#333",
        hoverBackground: "rgba(64, 158, 255, 0.08)",
        hoverText: "#333",
      };
      break;
    case "#191a23":
      theme.value = {
        background: "#191a23",
        activeBackground: "#4D70FF",
        activeText: "#fff",
        normalText: "#fff",
        hoverBackground: "rgba(64, 158, 255, 0.08)",
        hoverText: "#fff",
      };
      break;
  }
};

getTheme();
</script>

<style lang="scss">
.el-sub-menu__title:hover,
.el-menu-item:hover {
  background: transparent;
}

.el-scrollbar {
  .el-scrollbar__view {
    height: 100%;
  }
}
.menu-info {
  .menu-contorl {
    line-height: 52px;
    font-size: 20px;
    display: table-cell;
    vertical-align: middle;
  }
}
</style>
