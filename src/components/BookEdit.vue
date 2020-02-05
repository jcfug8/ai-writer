<template>
  <div id="bookEdit">
    <h1>Book Edit</h1>
    <SimpleButton v-on:click="$router.push('/books')">Back</SimpleButton>
  </div>
</template>

<script>
// @ is an alias to /src
import SimpleButton from "@/components/SimpleButton.vue";

export default {
  name: "book-edit",
  props: {},
  components: {
    SimpleButton
  }
};
</script>
