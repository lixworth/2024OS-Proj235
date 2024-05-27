<script setup>
import {onMounted} from "vue";
import {useLoadingBar} from "naive-ui";
const loadingBar = useLoadingBar();

onMounted(() => {
  loadingBar.finish();
})
</script>

<template>
  Login Page
</template>

<style scoped>

</style>