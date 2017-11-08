import Vue from 'vue'
import Router from 'vue-router'
import Vaccinations from '../components/Vaccinations.vue'
import Login from '../components/Login.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'login',
      component: Login
    },
    {
      path: '/vaccinations',
      name: 'vaccinations',
      component: Vaccinations
    }
  ]
})
