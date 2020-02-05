<template>
  <div id="bookList">
    <h1>Book List</h1>
    <SimpleButton v-on:click="$router.push('/books/edit')">Add Book</SimpleButton>
  </div>
</template>

<script>
// @ is an alias to /src
import SimpleButton from "@/components/SimpleButton.vue";

export default {
  name: "book-list",
  props: {},
  components: {
    SimpleButton
  }
};
</script>
