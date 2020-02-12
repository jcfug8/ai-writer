<template>
  <button class="simple-button" @click="$emit('click')">
    <slot></slot>
  </button>
</template>

<script>
// @ is an alias to /src

export default {
  name: "simple-button",
  props: {
    link: Boolean
  },
  components: {}
};
</script>

<style scoped>
.simple-button {
  background-color: #e6e6e6;
  border: 1px solid #c4c4c4;
  border-radius: 5px;
  padding: 2px 6px;
}
.simple-button:hover {
  background-color: #adadad;
  color: white;
}
</style>
