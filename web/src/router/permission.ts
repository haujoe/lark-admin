import router from '@/router/index'
import { useNProgress } from '@/hooks/use-nprogress'

const { start, done } = useNProgress()

router.beforeEach((to, from, next) => {
  start()
  next()
})

router.afterEach(() => {
  done()
})
