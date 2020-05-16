import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../views/Home.vue';
import Login from '../views/auth/Login.vue';
import Categories from '../views/Categories.vue';
import store from '../store';

Vue.use(VueRouter);

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/home',
    name: 'Home',
    component: Home,
    meta: {
      private: true,
    },
  },
  {
    path: '/categories',
    name: 'Categories',
    component: Categories,
    meta: {
      private: true,
    },
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.matched.some((record) => record.meta.private)) {
    if (!store.state.auth.authenticated) {
      return next('/login');
    }
  }
  if (to.fullPath === '/login') {
    if (store.state.auth.authenticated) {
      return next('/home');
    }
  }
  next();
});

export default router;
