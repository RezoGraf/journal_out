import Vue from 'vue'
import Router from 'vue-router'
import Vaccinations from '../components/Vaccinations.vue'
import Login from '../components/Login.vue'

Vue.use(Router)

var Auth = {
  loggedIn: false,
  login: function () {
    this.loggedIn = true
    router.push('vaccinations')
  },
  logout: function () { this.loggedIn = false }
}

const router = new Router({
  routes: [
    {
      path: '/',
      component: Vaccinations,
      meta: {requiresAuth: true}
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/vaccinations',
      name: 'vaccinations',
      component: Vaccinations,
      meta: {requiresAuth: true}
    }
  ]
})

export default router
export {Auth}

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // этот путь требует авторизации, проверяем залогинен ли
    // пользователь, и если нет, перенаправляем на страницу логина
    if (!Auth.loggedIn) {
      next({
        path: '/login',
        query: { redirect: to.fullPath }
      })
    } else {
      next()
    }
  } else {
    next() // всегда так или иначе нужно вызвать next()!
  }
})
