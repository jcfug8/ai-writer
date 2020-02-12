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
  logout: async function() {
    console.log("logging out");
    store.state.isLoggedIn = false;
    try {
      await fetch(`http://${store.state.baseURL}/api/session`, {
        method: "DELETE",
        mode: "cors",
        credentials: "include", // include, *same-origin, omit
        headers: {
          "Content-Type": "application/json"
        }
      });
    } catch (err) {
      console.log("error logging out ", err);
    }
    router.push("/login");
  }
};

let checkLoggedIn = async function(to) {
  if (!store.state.isLoggedIn) {
    let res;
    let data;
    try {
      res = await fetch(`http://${store.state.baseURL}/api/session`, {
        method: "GET",
        mode: "cors",
        credentials: "include", // include, *same-origin, omit
        headers: {
          "Content-Type": "application/json"
        }
      });
      data = await res.json();
    } catch (err) {
      console.log("error checking if logged in", err);
    }

    if (res.status == 200 && data) {
      store.state.isLoggedIn = true;
      store.state.userData = data;
    }
  }

  if (!store.state.isLoggedIn && !unathenticatedRouteNames.includes(to.name)) {
    return "/login";
  } else if (store.state.isLoggedIn && to.name == "login") {
    return "/";
  } else {
    return null;
  }
};

router.beforeEach(async (to, from, next) => {
  let path = await checkLoggedIn(to, from, next);
  if (path == null) {
    next();
  } else {
    next(path);
  }
});

new Vue({
  router,
  render: h => h(App),
  data: store
}).$mount("#app");
