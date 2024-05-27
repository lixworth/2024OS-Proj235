<script setup>
import GlobalComponent from "./components/GlobalComponent.vue";
import { firstLoadingRef } from '@/store/first_loading';
</script>

<template>
  <div>
    <global-component>
      <n-layout embedded position="absolute">
        <app-header></app-header>
        <n-layout position="absolute" style="top: 64px; min-height: calc(100vh - 64px)" v-bind:native-scrollbar="isMobile">
          <n-spin v-bind:show="firstLoadingRef" style="padding-top: 10px">
            <div class="container">
              <router-view></router-view>
              <n-back-top :right="100" />
            </div>
          </n-spin>
        </n-layout>
      </n-layout>
    </global-component>
  </div>
</template>

<style scoped>

</style>
