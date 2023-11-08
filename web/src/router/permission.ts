import router from '@/router/index'
import { useNProgress } from '@/hooks/use-nprogress'
import isWhiteList from '@/config/white-list'

const { start, done } = useNProgress()

router.beforeEach((to, from, next) => {
  start()
  const token = localStorage.getItem('token')

  // 判断该用户是否已经登录
  if (!token) {
    // 如果在免登录的白名单中，则直接进入
    if (isWhiteList(to)) {
      next()
    } else {
      // 其他没有访问权限的页面将被重定向到登录页面
      done()
      next('/login')
    }
    return
  }

  // 如果已经登录，并准备进入 Login 页面，则重定向到主页
  if (to.path === '/login') {
    done()
    return next({ path: '/' })
  }

  next()
})

router.afterEach(() => {
  done()
})
