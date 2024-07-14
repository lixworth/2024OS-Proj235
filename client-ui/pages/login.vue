<script setup lang="ts">
const state = reactive({
  // username: 'root',
  password: 'adminpassword',
})
const loading = ref(true)
const auth = useAuth()
function onSubmit() {
  auth.value = true
  useRouter().push({
    path: '/',
  })
}
const pwd_visiable = ref(false)

onMounted(() => {
  loading.value = false
})
</script>

<template>
  <div class="h-screen flex items-center justify-center overlay">
    <UCard class="max-w-sm w-full bg-white/75 dark:bg-white/5 backdrop-blur">
      <div class="w-full max-w-sm space-y-6">
        <div class="text-center">
          <div class="mb-2 pointer-events-none">
            <span class="i-heroicons-wrench-screwdriver w-8 h-8 flex-shrink-0 text-gray-900 dark:text-white" />
          </div>
          <p class="text-2xl text-gray-900 dark:text-white font-bold">
            System Updater
          </p>
          <span class="text-gray-500 dark:text-gray-400 mt-1">
            登陆到OTA版本升级系统
          </span>
        </div>
        <UForm :state="state" class="space-y-4">
          <UFormGroup name="password" label="密码" class="pb-2">
            <UInput
              v-model="state.password"
              placeholder="请输入操作密码"
              icon="i-heroicons-lock-closed"
              :type="pwd_visiable ? 'text' : 'password'"
              :ui="{ icon: { trailing: { pointer: '' } } }"
            >
              <template #trailing>
                <UButton
                  color="gray"
                  variant="link"
                  :disabled="loading"
                  :icon="pwd_visiable ? 'i-heroicons-eye' : 'i-heroicons-eye-slash'"
                  :padded="false"
                  @click="pwd_visiable = !pwd_visiable"
                />
              </template>
            </UInput>
          </UFormGroup>
          <UButton label="登陆" block :loading="loading" :disabled="loading" @click="onSubmit">
            <template #trailing>
              <UIcon name="i-heroicons-arrow-right-20-solid" class="w-5 h-5" />
            </template>
          </UButton>
        </UForm>
      </div>
      <template #footer>
        <span class="text-gray-500 dark:text-gray-400 text-xs">
          初始密码请查看初次运行时的日志输出 <br> 重置密码请终端运行 <code class="text-primary font-medium">sys-upgrader resetpwd</code>
        </span>
      </template>
    </UCard>
  </div>
</template>
