<script lang="ts" setup>
import { useRoute } from '#imports'

const route = useRoute()

const links = [{
  label: 'Installation',
  icon: 'i-heroicons-home',
  to: '/getting-started/installation',
}, {
  label: 'Horizontal Navigation',
  icon: 'i-heroicons-chart-bar',
  to: `${route.path.startsWith('/dev') ? '/dev' : ''}/components/horizontal-navigation`,
}, {
  label: 'Command Palette',
  icon: 'i-heroicons-command-line',
  to: '/components/command-palette',
}, {
  label: 'Profile',
  avatar: {
    src: 'https://avatars.githubusercontent.com/u/739984?v=4',
  },
  badge: 100,
}]
</script>

<template>
  <UHorizontalNavigation :links="links" />
</template>
