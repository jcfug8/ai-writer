<template>
  <div id="login">
    <div id="login-container" class="container">
      <div v-if="isLoggingIn" id="formLogin">
        <h2>Login</h2>
        <a v-on:click="toggle">Sign Up</a>
        <br />
        <input placeholder="Email" type="text" v-model="loginData.email" />
        <br />
        <input placeholder="Password" type="password" v-model="loginData.password" />
        <br />
        <SimpleButton v-on:click="login">Submit</SimpleButton>
      </div>
      <div v-else id="formNewUser">
        <h2>New User</h2>
        <br />
        <a v-on:click="toggle">Login</a>
        <br />
        <input placeholder="Firstname" type="text" v-model="signupData.firstname" />
        <br />
        <input placeholder="Last" type="text" v-model="signupData.lastname" />
        <br />
        <br />
        <input placeholder="Email" type="text" v-model="signupData.email" />
        <br />
        <br />
        <input placeholder="Password" type="password" v-model="signupData.password" />
        <br />
        <input placeholder="Repeat Password" type="password" v-model="signupData.password" />
        <br />
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

<style>
#login-container {
  margin-top: 20px;
  text-align: center;
}
#login-container > div {
  display: inline-block;
  margin: auto;
  text-align: center;
  border: 1px solid #f1f1f1;
  box-shadow: 1px 1px 2px 0px #e5e5e5;
  box-sizing: border-box;
  border-radius: 5px;
  padding: 0 20px 20px 20px;
}
input[type="text"],
input[type="password"] {
  border-radius: 3px;
  border: 1px solid #dfdfdf;
  padding: 5px;
  border-style: inherit;
  margin-top: 5px;
}
#login-container h2 {
  margin-bottom: 5px;
}
#login-container a {
  font-size: 0.8rem;
  text-decoration: underline;
  display: inline-block;
  margin-bottom: 10px;
  cursor: pointer;
}
#login-container button {
  margin-top: 10px;
}
</style>