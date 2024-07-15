import antfu from '@antfu/eslint-config'

export default antfu({
  typescript: true,
  vue: true,
  ignores: [
    '.nuxt',
    '.output',
  ],
  rules: {
    'no-undef': 'off',
  },
})
