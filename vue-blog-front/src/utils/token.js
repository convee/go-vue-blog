import { lStorage } from './cache'

const TOKEN_CODE = 'access_token'
const DURATION = 6 * 60 * 60

export function getToken() {
  return lStorage.get(TOKEN_CODE)
}

export function setToken(token) {
  lStorage.set(TOKEN_CODE, token, DURATION)
}

export function removeToken() {
  lStorage.remove(TOKEN_CODE)
}
