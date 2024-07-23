<script lang="ts" setup>
const state = reactive({
  server: 'http://cscc.kokomi.ltd/api',
  auto_update: true,
})
const loading = ref(true)
async function onSubmit() {

}
</script>

<template>
  <div>
    <UForm :state="state" class="space-y-4" @submit="onSubmit">
      <UFormGroup
        label="更新服务器" class="pb-2"
        help="请保证以 &quot;https://&quot; 或 &quot;http://&quot; 开头。"
      >
        <UInput
          v-model="state.server"
          :disabled="loading"
          placeholder="例如: https://mirror.lix.moe/"
          icon="i-heroicons-globe-alt"
        />
      </UFormGroup>
      <UFormGroup>
        <UCheckbox v-model="state.auto_update" label="自动检查并更新" />
      </UFormGroup>
      <UButton label="保存设置" block :loading="loading" :disabled="loading" @click="onSubmit" />
    </UForm>
  </div>
</template>
