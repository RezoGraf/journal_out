import Vue from 'vue'
import Router from 'vue-router'
import Vaccinations from '../components/Vaccinations.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'vaccinations',
      component: Vaccinations
    }
  ]
})
