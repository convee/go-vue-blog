import vue from '@vitejs/plugin-vue'

import Unocss from 'unocss/vite'

import unplugin from './unplugin'

export function createVitePlugins(viteEnv, isBuild) {
  const plugins = [vue(), ...unplugin, Unocss()]
  return plugins
}
