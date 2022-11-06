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
          :background-color="theme.background"
          :text-color="theme.normalText"
          :active-text-color="theme.activeText"
          unique-opened
          @select="selectMenuItem"
        >
          <template v-for="item in routerStore.baseRouter[0].children">
            <aside-component
              v-if="!item.hidden"
              :router-info="item"
              :theme="theme"
            />
          </template>
        </el-menu>
      </transition>
    </el-scrollbar>
  </div>
</template>

<script lang="ts">
export default {
  name: "Aside",
};
</script>

<script lang="ts" setup>
import { ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import useStore from "@/stores/stores";
import AsideComponent from "./asideComponent/asideComponent.vue";

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

const route = useRoute();
const router = useRouter();
const { userStore, routerStore } = useStore();

const selectMenuItem = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
  console.log(route.name);
  const query = {};
  const params = {};
  router.push({ name: key, query, params });
};
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
