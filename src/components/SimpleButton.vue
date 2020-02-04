<template>
  <button @click="$emit('click')">
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
