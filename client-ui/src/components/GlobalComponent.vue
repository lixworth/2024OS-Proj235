<script setup lang="ts">
import {
  lightTheme,
  NConfigProvider,
  NMessageProvider,
  NLoadingBarProvider,
  NDialogProvider,
  type GlobalTheme,
  darkTheme
} from 'naive-ui';
import { ref } from 'vue';
import ComponentInit from '@/components/ComponentInit.vue';

const theme = ref<GlobalTheme | null>(null);
theme.value = lightTheme;
const handleTheme = (dark: boolean) => {
  if (dark) {
    theme.value = darkTheme;
  } else {
    theme.value = lightTheme;
  }
};
const dark = window.matchMedia('(prefers-color-scheme: dark)');
handleTheme(dark.matches);
dark.addEventListener('change', (e) => {
  handleTheme(e.matches);
});
</script>
<template>
  <n-config-provider :theme="theme">
    <n-loading-bar-provider>
      <n-message-provider>
        <n-dialog-provider>
          <component-init>
            <slot></slot>
          </component-init>
        </n-dialog-provider>
      </n-message-provider>
    </n-loading-bar-provider>
  </n-config-provider>
</template>
