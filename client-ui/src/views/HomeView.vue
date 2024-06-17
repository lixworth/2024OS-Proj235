<script setup lang="ts">
import { onMounted } from 'vue'
import {useLoadingBar} from "naive-ui";

const loadingBar = useLoadingBar();
onMounted(() => {
  loadingBar.finish;
});
</script>

<template>Home Page</template>
