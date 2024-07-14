import antfu from '@antfu/eslint-config'

export default antfu({
  ignores: [
    '.nuxt',
    '.output',
  ],
  rules: {
    "no-undef": "off"
  }
})
