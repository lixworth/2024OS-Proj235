<script setup lang="ts">

</script>

<template>
  <div class="h-screen flex pt-36 justify-center overlay">
    <div class="max-w-4xl w-full">
      <div class="absolute top-4 left-4">
        <UButton
          to="/login"
          color="gray"
          variant="ghost"
          class="inline-flex items-center" icon="i-heroicons-arrow-left" size="lg"
        >
          登出
        </UButton>
      </div>
      <UCard
        class="w-full bg-white/75 dark:bg-white/5 backdrop-blur"
      >
        <div class="px-4 sm:px-0">
          <h3 class="text-lg text-base font-semibold leading-7 text-gray-900 dark:text-white">
            系统版本升级
          </h3>
          <p class="mt-1 max-w-2xl text-sm leading-6 text-gray-500 dark:text-gray-400">
            OTA版本升级系统
          </p>
        </div>
        <UAlert
          class="mt-6"
          icon="i-heroicons-shield-exclamation"
          color="orange"
          variant="soft"
          title="新版本更新 v1.0.1 -> v1.1.0"
          description="当前系统版本 (v1.0.1) 低于最新版本 (v1.1.0) ，建议尽快更新到最新系统，保证系统安全稳定！"
        />
        <div class="mt-6 border-t border-gray-200 dark:border-gray-600">
          <dl class="divide-y divide-gray-200 dark:divide-gray-600">
            <div class="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
              <dt class="text-sm font-medium leading-6 text-gray-900 dark:text-white">
                设备名称
              </dt>
              <dd class="mt-1 text-sm leading-6 text-gray-700 dark:text-gray-300 sm:col-span-2 sm:mt-0">
                lixworth@WorthdeMacBook-Pro-2.local
              </dd>
            </div>
            <div class="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
              <dt class="text-sm font-medium leading-6 text-gray-900 dark:text-white">
                硬件信息
              </dt>
              <dd class="mt-1 text-sm leading-6 text-gray-700 dark:text-gray-300 sm:col-span-2 sm:mt-0">
                <UICon icon="i-hericons-apple">
                  Apple M1 Pro / 16384MiB
                </uicon>
              </dd>
            </div>
            <div class="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
              <dt class="text-sm font-medium leading-6 text-gray-900 dark:text-white">
                系统版本
              </dt>
              <dd class="mt-1 text-sm leading-6 text-gray-700 dark:text-gray-300 sm:col-span-2 sm:mt-0">
                <UBadge color="primary" size="lg" variant="soft">
                  v1.0.1
                </UBadge>
                (d41d8cd98f00b204e9800998ecf8427e)
              </dd>
            </div>
            <div class="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
              <dt class="text-sm font-medium leading-6 text-gray-900 dark:text-white">
                手动更新
              </dt>
              <dd class="mt-1 text-sm leading-6 text-gray-700 dark:text-gray-300 sm:col-span-2 sm:mt-0">
                <UButton
                  icon="i-heroicons-arrow-path"
                  size="md"
                >
                  检查更新
                </UButton>
                <UButton
                  class="mx-2"
                  size="md"
                  icon="i-heroicons-arrow-up-tray"
                >
                  手动上传升级包
                </UButton>
              </dd>
            </div>
            <div class="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
              <dt class="text-sm font-medium leading-6 text-gray-900 dark:text-white">
                升级设置
              </dt>
              <dd class="mt-1 text-sm leading-6 text-gray-700 dark:text-gray-300 sm:col-span-2 sm:mt-0">
                111
              </dd>
            </div>
            <div class="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
              <dt class="text-sm font-medium leading-6 text-gray-900 dark:text-white">
                升级日志
              </dt>
              <dd class="mt-2 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                <ul role="list" class="divide-y divide-gray-100 dark:divide-gray-600 rounded-md border border-gray-200 dark:border-gray-600">
                  <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm leading-6">
                    <div class="flex w-0 flex-1 items-center">
                      <svg class="h-5 w-5 flex-shrink-0 text-gray-400  dark:text-gray-300 " viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                        <path fill-rule="evenodd" d="M15.621 4.379a3 3 0 00-4.242 0l-7 7a3 3 0 004.241 4.243h.001l.497-.5a.75.75 0 011.064 1.057l-.498.501-.002.002a4.5 4.5 0 01-6.364-6.364l7-7a4.5 4.5 0 016.368 6.36l-3.455 3.553A2.625 2.625 0 119.52 9.52l3.45-3.451a.75.75 0 111.061 1.06l-3.45 3.451a1.125 1.125 0 001.587 1.595l3.454-3.553a3 3 0 000-4.242z" clip-rule="evenodd" />
                      </svg>
                      <div class="ml-4 flex min-w-0 flex-1 gap-2  dark:text-gray-300 ">
                        <span class="truncate font-medium">20240202230010010101.log</span>
                        <span class="flex-shrink-0 text-gray-400">2.4kb</span>
                      </div>
                    </div>
                    <div class="ml-4 flex-shrink-0">
                      <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500">查看</a>
                    </div>
                  </li>
                  <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm leading-6">
                    <div class="flex w-0 flex-1 items-center">
                      <svg class="h-5 w-5 flex-shrink-0 text-gray-400" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                        <path fill-rule="evenodd" d="M15.621 4.379a3 3 0 00-4.242 0l-7 7a3 3 0 004.241 4.243h.001l.497-.5a.75.75 0 011.064 1.057l-.498.501-.002.002a4.5 4.5 0 01-6.364-6.364l7-7a4.5 4.5 0 016.368 6.36l-3.455 3.553A2.625 2.625 0 119.52 9.52l3.45-3.451a.75.75 0 111.061 1.06l-3.45 3.451a1.125 1.125 0 001.587 1.595l3.454-3.553a3 3 0 000-4.242z" clip-rule="evenodd" />
                      </svg>
                      <div class="ml-4 flex min-w-0 flex-1 gap-2 dark:text-gray-300 ">
                        <span class="truncate font-medium">20240202230010010102.log</span>
                        <span class="flex-shrink-0 text-gray-400">2.1kb</span>
                      </div>
                    </div>
                    <div class="ml-4 flex-shrink-0">
                      <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500">查看</a>
                    </div>
                  </li>
                </ul>
              </dd>
            </div>
          </dl>
        </div>
      </UCard>
    </div>
  </div>
</template>
