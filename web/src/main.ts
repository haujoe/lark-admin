// core
import { createApp } from 'vue'
import store from '@/stores'
import router from '@/router'
import App from '@/App.vue'

// css
import 'virtual:uno.css'
import '@/assets/main.css'

const app = createApp(App)

app.use(store).use(router)

router.isReady().then(() => {
  app.mount('#app')
})
