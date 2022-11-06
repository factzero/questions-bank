<template>
  <component :is="menuComponent" :router-info="routerInfo" :theme="theme">
    <template v-if="routerInfo.children && routerInfo.children.length">
      <AsideComponent
        v-for="item in routerInfo.children"
        :router-info="item"
        :theme="theme"
      />
    </template>
  </component>
</template>

<script>
export default {
  name: "AsideComponent",
};
</script>

<script setup>
import { computed } from "vue";
import MenuItem from "./menuItem.vue";
import SubMenu from "./subMenu.vue";

const props = defineProps({
  routerInfo: {
    type: Object,
    default: () => null,
  },
  isCollapse: {
    default: function () {
      return false;
    },
    type: Boolean,
  },
  theme: {
    default: function () {
      return {};
    },
    type: Object,
  },
});

const menuComponent = computed(() => {
  //   console.log("asideComponent:", props.routerInfo);
  if (
    props.routerInfo.children &&
    props.routerInfo.children.filter((item) => !item.hidden).length
  ) {
    return SubMenu;
  } else {
    return MenuItem;
  }
});
</script>
