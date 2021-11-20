import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../components/Home.vue'
import Redirect from '../components/Redirect.vue'

// let url = window.location.href

Vue.use(VueRouter)

const routes = [
  {
    path: '/home',
    name: 'Home',
    component: Home
  },
  {
    path: '*',
    name: 'Redirect',
    component: Redirect
  },
]

const router = new VueRouter({
  mode: 'history',
  routes
})

export default router
