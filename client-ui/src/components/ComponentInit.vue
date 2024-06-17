<script setup lang="ts">
import { useDialog, useLoadingBar, useMessage } from 'naive-ui';
import { onMounted } from 'vue';
import { loadingBarApiRef } from '@/stores/loading';

const loadingBar = useLoadingBar();
// @ts-ignore
window.$message = useMessage();
// @ts-ignore
window.$dialog = useDialog();

onMounted(() => {
  loadingBarApiRef.value = loadingBar;
  loadingBar.finish();
});
</script>

<template>
  <slot></slot>
</template>
