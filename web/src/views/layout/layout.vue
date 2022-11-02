<template>
  <el-container class="layout-cont">
    <el-container>
      <el-aside class="main-cont main-left gva-aside">
        <div class="tilte" :style="{ background: backgroundColor }">
          <div class="tit-text">试题系统</div>
        </div>
        <Aside class="aside" />
      </el-aside>
      <el-main class="main-cont main-right">
        <div>xxx</div>
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { computed } from "vue";
import Aside from "@/views/layout/aside/aside.vue";

const backgroundColor = computed(() => {
  return "#191a23";
});
</script>

<style lang="scss">
@import "@/style/mobile.scss";

.dark {
  background-color: #191a23 !important;
  color: #fff !important;
}
.light {
  background-color: #fff !important;
  color: #000 !important;
}
</style>
