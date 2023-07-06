import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
    meta: {
      keepAlive: true,
    }
  },
  
  {
    path: '/about',
    name: 'about',
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')    
  },
  {
    path: '/person',
    name:'person',
    component: () => import(/* webpackChunkName: "about" */ '../views/PersonalView.vue')  
  },
  {
    path: '/post',
    name:'post',
    component: () => import(/* webpackChunkName: "about" */ '../views/PhotoPublish.vue')  
  },
  {
    path: '/otherPerson',
    name: 'otherPerson',
    component: () => import(/* webpackChunkName: "about" */ '../views/otherPersonal.vue')
  },
]

const router = new VueRouter({
  routes
})

export default router
