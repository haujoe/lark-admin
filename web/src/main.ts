// core
import { createApp } from 'vue'
import store from '@/stores'
import router from '@/router'
import App from '@/App.vue'
import '@/router/permission'

// load
import { loadDirectives } from '@/directives'

// css
import 'virtual:uno.css'
import '@/styles/index.scss'

const app: any = createApp(App)

/** 加载自定义指令 */
loadDirectives(app)

app.use(store).use(router)

router.isReady().then(() => {
  app.mount('#app')
})
