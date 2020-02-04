import Vue from "vue";
import App from "./App.vue";
import { router, unathenticatedRouteNames } from "./router";

Vue.config.productionTip = false;

var store = {
  state: {
    baseURL: process.env.VUE_APP_BASE_URL,
    isLoggedIn: false,
    userData: null
  },
  login(userData) {
    console.log("logging in", userData);
    store.state.isLoggedIn = true;
    store.state.userData = userData;
    router.push("/");
  },
  logout() {
    console.log("logging out");
    store.state.isLoggedIn = false;
    router.push("/login");
  }
};

router.beforeEach((to, from, next) => {
  if (!store.state.isLoggedIn && !unathenticatedRouteNames.includes(to.name)) {
    next("/login");
  } else {
    next();
  }
});

new Vue({
  router,
  render: h => h(App),
  data: store
}).$mount("#app");
