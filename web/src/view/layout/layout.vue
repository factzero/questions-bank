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
        <transition
          :duration="{ enter: 800, leave: 100 }"
          mode="out-in"
          name="el-fade-in-linear"
        >
          <div :style="{ width: `calc(100% - ${'220px'})` }" class="topfix">
            <el-row>
              <el-col>
                <el-header class="header-cont">
                  <el-row class="pd-0">
                    <el-col
                      :xs="2"
                      :lg="1"
                      :md="1"
                      :sm="1"
                      :xl="1"
                      style="z-index: 100"
                    >
                      <div class="menu-total">
                        <div class="gvaIcon gvaIcon-arrow-double-left" />
                      </div>
                    </el-col>
                    <el-col
                      :xs="10"
                      :lg="14"
                      :md="14"
                      :sm="9"
                      :xl="14"
                      :pull="1"
                    >
                      <el-breadcrumb class="breadcrumb">
                        <el-breadcrumb-item
                          >el-breadcrumb-item</el-breadcrumb-item
                        >
                      </el-breadcrumb>
                    </el-col>
                    <el-col :xs="12" :lg="9" :md="9" :sm="14" :xl="9">
                      <div class="right-box">
                        <el-dropdown>
                          <div
                            class="dp-flex justify-content-center align-items height-full width-full"
                          >
                            <span class="header-avatar" style="cursor: pointer">
                              <span style="margin-left: 5px">nickName</span>
                              <el-icon><arrow-down /></el-icon>
                            </span>
                          </div>
                          <template #dropdown>
                            <el-dropdown-menu class="dropdown-group">
                              <el-dropdown-item icon="avatar"
                                >个人信息</el-dropdown-item
                              >
                              <el-dropdown-item icon="reading-lamp"
                                >登 出</el-dropdown-item
                              >
                            </el-dropdown-menu>
                          </template>
                        </el-dropdown>
                      </div>
                    </el-col>
                  </el-row>
                </el-header>
              </el-col>
            </el-row>
          </div>
        </transition>
        <router-view
          v-slot="{ Component }"
          element-loading-text="正在加载中"
          class="admin-box"
        >
          <div>
            <transition mode="out-in" name="el-fade-in-linear">
              <keep-alive>
                <component :is="Component" />
              </keep-alive>
            </transition>
          </div>
        </router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { computed } from "vue";
import Aside from "@/view/layout/aside/aside.vue";

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
