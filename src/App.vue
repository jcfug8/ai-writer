<template>
  <div id="app">
    <div id="content">
      <div class="shadow" id="nav">
        <div class="container">
          <router-link to="/" id="logo">AI Writer</router-link>
          <router-link to="/" exact>Home</router-link>
          <router-link v-if="this.$root.$data.state.isLoggedIn" to="/books" exact>Books</router-link>
          <router-link v-if="!this.$root.$data.state.isLoggedIn" to="/login">Login</router-link>
          <span class="dropdown" v-if="this.$root.$data.state.isLoggedIn">
            <input id="accountDropdown" class="dropdown-checkbox" type="checkbox" />
            <label
              class="dropdown-trigger"
              for="accountDropdown"
            >{{ this.$root.$data.state.userData.firstname }}</label>
            <div class="dropdown-menu">
              <router-link to="/account">Account</router-link>
              <a v-on:click="this.$root.$data.logout">Logout</a>
            </div>
          </span>
        </div>
      </div>
      <div id="routerContent">
        <router-view />
      </div>
    </div>
    <div id="footer">
      <div class="container">AI Writer</div>
    </div>
  </div>
</template>

<style>
body,
html {
  font-family: "Courier New", Courier, monospace;
  height: 100%;
  margin: 0;
}

#app {
  height: 100%;
  display: flex;
  flex-direction: column;
}

#content {
  flex: 1 0 auto;
}

#routerContent {
  height: 100%;
  position: relative;
}

#routerContent > div {
  padding-top: 10px;
  padding-bottom: 10px;
}

#footer {
  padding: 5px;
  flex-shrink: 0;
  background-color: #e6e6e6;
}

.shadow {
  box-shadow: 1px 2px 3px 0px #bbbbbb;
}

#nav {
  text-align: right;
  background-color: #e6e6e6;
}

#logo {
  float: left;
  margin-left: 0px !important;
  padding-left: 0px !important;
}

#nav > div > * {
  margin-left: 10px;
  padding: 5px;
  display: inline-block;
  text-decoration: none;
  cursor: pointer;
  color: black;
}

#nav > div > .router-link-active:not(#logo) {
  background-color: grey;
}

#nav > div > *:visited {
  color: initial;
}

.container {
  max-width: 1000px;
  margin: auto;
}

.dropdown {
  position: relative;
}

.dropdown-trigger {
  cursor: pointer;
}

.dropdown-checkbox {
  display: none;
}

.dropdown-menu {
  z-index: 1;
}

.dropdown-menu > * {
  z-index: 1;
  color: black;
  text-decoration: none;
  padding: 5px;
  display: block;
}

.dropdown-menu > *:visited {
  background-color: initial;
}

.dropdown-checkbox:not(:checked) ~ .dropdown-menu {
  display: none;
}

.dropdown:hover .dropdown-menu,
.dropdown-checkbox:checked ~ .dropdown-menu {
  display: block;
  position: absolute;
  background-color: #e6e6e6;
  right: 0;
}

.dropdown-checkbox:checked ~ .dropdown-trigger {
}

.clear-contain::after {
  content: "";
  clear: both;
  display: table;
}

.float-left {
  float: left;
}
</style>
