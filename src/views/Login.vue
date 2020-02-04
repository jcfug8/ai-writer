<template>
  <div id="login">
    <div class="container">
      <div v-if="isLoggingIn" id="formLogin">
        <h2>Login</h2>
        <SimpleButton v-on:click="toggle">Sign Up</SimpleButton>
        <input v-model="loginData.email" />
        <input v-model="loginData.password" />
        <SimpleButton v-on:click="login">Submit</SimpleButton>
      </div>
      <div v-else id="formNewUser">
        <h2>New User</h2>
        <SimpleButton v-on:click="toggle">Login</SimpleButton>
        <input v-model="signupData.firstname" />
        <input v-model="signupData.lastname" />
        <input v-model="signupData.email" />
        <input v-model="signupData.password" />
        <SimpleButton v-on:click="singup">Submit</SimpleButton>
      </div>
    </div>
  </div>
</template>

<script>
import SimpleButton from "@/components/SimpleButton.vue";

const FORM_STATE_LOGIN = 0;
const FORM_STATE_SIGN_UP = 1;

export default {
  name: "Login",
  components: {
    SimpleButton
  },
  data: function() {
    return {
      form_state: FORM_STATE_LOGIN,
      loginData: {
        email: "",
        password: ""
      },
      signupData: {
        firstname: "",
        lastname: "",
        email: "",
        password: ""
      }
    };
  },
  props: {},
  computed: {
    isLoggingIn: function() {
      return this.form_state === FORM_STATE_LOGIN;
    }
  },
  methods: {
    toggle: function() {
      if (this.form_state === FORM_STATE_LOGIN) {
        this.form_state = FORM_STATE_SIGN_UP;
      } else {
        this.form_state = FORM_STATE_LOGIN;
      }
    },
    login: async function() {
      let data;
      let res;
      try {
        res = await fetch(
          `http://${this.$root.$data.state.baseURL}/api/session`,
          {
            method: "POST",
            mode: "cors",
            credentials: "include", // include, *same-origin, omit
            headers: {
              "Content-Type": "application/json"
            },
            body: JSON.stringify({
              email: this.loginData.email,
              password: this.loginData.password
            })
          }
        );
        data = await res.json();
        console.log(data);
      } catch (err) {
        console.log(await err.json());
      }
      if (res.status == 201) {
        this.$root.$data.login(data);
      }
    },
    singup: async function() {
      let data;
      let res;
      try {
        res = await fetch(`http://${this.$root.$data.state.baseURL}/api/user`, {
          method: "POST",
          mode: "cors",
          credentials: "include", // include, *same-origin, omit
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify({
            email: this.signupData.email,
            password: this.signupData.password,
            firstname: this.signupData.firstname,
            lastname: this.signupData.lastname
          })
        });
        data = await res.json();
        console.log(data);
      } catch (err) {
        console.log(await err.json());
      }
      if (res.status == 201) {
        this.toggle();
      }
    }
  }
};
</script>
