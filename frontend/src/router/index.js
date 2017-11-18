import Vue from 'vue'
import Router from 'vue-router'
import Vaccinations from '../components/Vaccinations.vue'
import Login from '../components/Login.vue'
import axios from 'axios'

Vue.use(Router)

var Auth = {
  loggedIn: false,
  errauth: null,
  checkToken: function () {
    var ss = this
    console.log(localStorage.getItem('token_vaccinations'))
    var token = localStorage.getItem('token_vaccinations')
    console.log(token)
    axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token_vaccinations')
    axios.post('http://localhost:8084/jwt/checktoken')
      .then(function (response) {
        console.log(response.status)
        if (response.status === 200) {
          ss.loggedIn = true
          router.push('vaccinations')
        }
      })
      .catch(function (error) {
        console.log(error)
      })
  },
  login: function (name, pass) {
    var ss = this
    const dataform = new FormData()
    dataform.append('name', name)
    dataform.append('pass', pass)
    console.log('!!!!!!!', dataform)
    axios.post('http://localhost:8084/login', dataform, {responeType: 'application/json'})
      .then(function (response) {
        if (response.status === 200) {
          var data = JSON.parse(JSON.stringify(response.data))
          console.log(data.token)
          localStorage.setItem('token_vaccinations', data.token)
          ss.loggedIn = true
          router.replace('vaccinations')
        } else {
          ss.errauth = JSON.parse(JSON.stringify(response.data))
        }
      })
      .catch(function (error) {
        console.log(error)
      })
  },
  logout: function () {
    this.loggedIn = false
    localStorage.removeItem('token_vaccinations')
    router.push('login')
  },
  refresh: function () {
    router.refresh('vaccinations')
  }
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
