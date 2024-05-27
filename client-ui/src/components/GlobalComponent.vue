<script setup lang="ts">
import {
  NConfigProvider,
  NMessageProvider,
  NLoadingBarProvider,
  NDialogProvider
} from 'naive-ui';
import ComponentInit from './ComponentInit.vue';
</script>
<template>
  <n-config-provider>
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