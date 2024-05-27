<script setup lang="ts">
import { NIcon, NLayoutHeader, NMenu, NPopover, NText, NDropdown, NButton, useDialog } from 'naive-ui';
import { ref, computed, onMounted, nextTick } from 'vue';
import { MenuOutline, HelpCircle } from '@vicons/ionicons5';
import { useIsMobile, useIsTablet } from '@/utils/composables';
import { useRoute } from 'vue-router';
import { useUserInfoStore } from '@/store/user';
import { firstLoadingRef } from '@/store/first_loading';
import { getAfterLoginMenu, beforeLoginMenu, userActionsOptions } from '@/components/header';

const route = useRoute();
const user_info = useUserInfoStore();
const dialog = useDialog();
const activeKey = computed<any>({
  get() {
    if (route.name !== undefined) {
      return route.name;
    } else {
      return null;
    }
  },
  set(newValue) {
    console.log(newValue);
  }
});
const isMobileRef = useIsMobile();
const mobilePopoverRef = ref<any>(null);

const handleUpdateMobileMenu = (value: string, { path }: any) => {
  mobilePopoverRef.value.setShow(false);
};
const handleClickTitle = () => {
  window.location.href = '/';
};
const afterLoginMenu = computed<any>({
  get() {
    return getAfterLoginMenu();
  },
  set() {}
});
onMounted(() => {});
</script>

<template>
  <n-layout-header bordered class="nav" position="static">
    <div class="drawer" v-if="isMobileRef">
      <n-popover
          ref="mobilePopoverRef"
          placement="bottom-end"
          display-directive="show"
          trigger="click"
          style="padding: 0; min-width: 240px; z-index: 100"
      >
        <template v-slot:trigger>
          <n-icon size="24" style="margin-top: 5px">
            <menu-outline />
          </n-icon>
        </template>
        <div style="overflow: auto; max-height: 79vh">
          <n-menu
              v-if="user_info.check()"
              size="small"
              v-bind:options="afterLoginMenu"
              v-model:value="activeKey"
              @update:value="handleUpdateMobileMenu"
              default-expand-all
          />
          <n-menu
              v-else
              size="small"
              v-bind:options="beforeLoginMenu"
              v-model:value="activeKey"
              @update:value="handleUpdateMobileMenu"
              default-expand-all
          ></n-menu>
        </div>
      </n-popover>
    </div>
    <n-text tag="div" class="siteName" v-on:click="handleClickTitle">
      <img src="@/assets/img.png" alt="mfuns union" />
      <span><b>BlackBE</b> Impact</span>
    </n-text>
    <div v-if="!isMobileRef">
      <template v-if="!firstLoadingRef">
        <n-menu
            v-if="user_info.check()"
            class="menuList"
            v-model:value="activeKey"
            mode="horizontal"
            :options="afterLoginMenu"
        />
        <n-menu v-else class="menuList" v-model:value="activeKey" mode="horizontal" :options="beforeLoginMenu" />
      </template>
    </div>
    <div class="userActions">
      <template v-if="user_info.check()">
        <n-dropdown trigger="hover" v-bind:options="userActionsOptions">
          <n-text style="display: flex; align-items: center">
            <span v-if="!isMobileRef">{{ user_info.info?.account }}</span>
            <img class="avatar" v-bind:src="user_info.info?.avatar" alt="avatar" />
          </n-text>
        </n-dropdown>
      </template>
      <template v-else>
        <n-button
            icon-placement="left"
            :bordered="false"
            strong
            v-if="!firstLoadingRef"
            @click="
            dialog.info({
              title: '你得自己帮助自己哦亲',
              content:
                '当面对两难的抉择时，不妨丢一枚硬币吧，并非是要靠那二分之一的机运来帮你做出抉择！而是因为当硬币被抛上空中，开始旋转的那一瞬间你会突然明白，自己想要的！',
              positiveText: '关闭'
            })
          "
        >
          <template #icon>
            <n-icon>
              <help-circle></help-circle>
            </n-icon>
          </template>
          帮助
        </n-button>
      </template>
    </div>
  </n-layout-header>
</template>

<style scoped lang="scss">
.nav {
  padding: 10px 16px;
  user-select: none;
  text-align: center;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.nav .menuList {
  justify-content: center;
  flex-grow: 1;
}
.nav .userActions {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  box-sizing: border-box;
  line-height: 1.75;
  font-size: 14px;
  cursor: pointer;
}
.nav .siteName {
  white-space: nowrap;
  text-align: left;
  flex: 1;
  cursor: pointer;
  display: flex;
  align-items: center;
  font-size: 18px;
}
.nav .siteName > img {
  margin-right: 5px;
  height: 28px;
  width: 28px;
}
.nav .drawer {
  justify-content: start;
  text-align: left;
  align-content: center;
  flex: 1;
}
.nav .userActions .avatar {
  margin-left: 8px;
  border-radius: 10%;
  width: 35px;
  height: 35px;
}

.nav .userActions > span:hover {
  color: #18a058; // todo primary theme
}
</style>