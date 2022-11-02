<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
</script>

<template>
  <RouterView />
</template>

<style scoped></style>

<style lang="scss">
@import "@/style/main.scss";
@import "@/style/base.scss";
@import "@/style/mobile.scss";
#app {
  background: #eee;
  height: 100vh;
  overflow: hidden;
  font-weight: 400 !important;
}
.el-button {
  font-weight: 400 !important;
}
</style>
