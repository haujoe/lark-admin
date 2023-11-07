import { permission } from '@/directives/permission'
import type { App } from 'vue'

/** 挂载自定义指令 */
export function loadDirectives(app: App) {
  app.directive('permission', permission)
}
