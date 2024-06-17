<script setup>
import { useDialog, useLoadingBar, useMessage } from 'naive-ui';
import { onMounted } from 'vue';
import { loadingBarApiRef } from '../store/loading';

const loadingBar = useLoadingBar();
window.$message = useMessage();
window.$dialog = useDialog();

onMounted(() => {
  loadingBarApiRef.value = loadingBar;
  loadingBar.finish();
});
</script>

<template>
  <slot></slot>
</template>
