import { useSessionStorage } from '#imports'

const auth = ref(false)

export function useAuthSession() {
  return useSessionStorage('AuthToken', 'null')
}

export function useAuth() {
  if (useAuthSession().value === 'null') {
    auth.value = false
  }
  else {
    auth.value = true
  }
  return auth
}

export function useLogout() {
  useAuthSession().value = 'null'
  auth.value = false
}
