import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";

Vue.use(VueRouter);

const unathenticatedRouteNames = ["home", "login"];

const routes = [
  {
    path: "/",
    name: "home",
    component: Home
  },
  {
    path: "/login",
    name: "login",
    component: () =>
      import(/* webpackChunkName: "login" */ "../views/Login.vue")
  },
  {
    path: "/books",
    name: "books",
    component: () =>
      import(/* webpackChunkName: "books" */ "../views/Books.vue")
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export { router, unathenticatedRouteNames };
