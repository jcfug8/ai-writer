<template>
  <div id="bookList">
    <div class="page-title-contain clear-contain">
      <h1 class="page-title float-left">Book List</h1>
      <SimpleButton class="float-left" v-on:click="createBook">Add Book</SimpleButton>
    </div>
    <div class="book-list-item clear-contain" v-for="book in books" v-bind:key="book.id">
      <h2>{{ book.name }}</h2>
      <SimpleButton v-on:click="deleteBook(book.id)">Delete</SimpleButton>
      <SimpleButton v-on:click="$router.push(`books/edit/${book.id}`)">Edit</SimpleButton>
      <div class="book-desc">{{ book.description }}</div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import SimpleButton from "@/components/SimpleButton.vue";

export default {
  name: "book-list",
  data: function() {
    return {
      books: null
    };
  },
  props: {},
  components: {
    SimpleButton
  },
  created() {
    // fetch the data when the view is created and the data is
    // already being observed
    this.listBooks();
  },
  methods: {
    listBooks: async function() {
      let data;
      let res;
      try {
        res = await fetch(
          `http://${this.$root.$data.state.baseURL}/api/books`,
          {
            method: "GET",
            mode: "cors",
            credentials: "include", // include, *same-origin, omit
            headers: {
              "Content-Type": "application/json"
            }
          }
        );
        data = await res.json();
        console.log(data);
      } catch (err) {
        console.log(await err.json());
      }
      if (res.status == 200) {
        console.log(data);
        this.books = data.books;
      }
    },
    createBook: async function() {
      let data;
      let res;
      try {
        res = await fetch(`http://${this.$root.$data.state.baseURL}/api/book`, {
          method: "POST",
          mode: "cors",
          credentials: "include", // include, *same-origin, omit
          headers: {
            "Content-Type": "application/json"
          }
        });
        data = await res.json();
        console.log(data);
      } catch (err) {
        console.log(await err.json());
      }
      if (res.status == 201) {
        console.log(data);
        this.$router.push(`/books/edit/${data.id}`);
      }
    },
    deleteBook: async function(bookId) {
      let data;
      let res;
      try {
        res = await fetch(`http://${this.$root.$data.state.baseURL}/api/book`, {
          method: "DELETE",
          mode: "cors",
          credentials: "include", // include, *same-origin, omit
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify({
            id: bookId
          })
        });
        data = await res.json();
        console.log(data);
      } catch (err) {
        console.log(await err.json());
      }
      if (res.status == 200) {
        this.listBooks();
      }
    }
  }
};
</script>

<style scoped>
.page-title-contain {
  margin: 15px 0 10px;
}
.page-title {
  margin: 0;
  margin: -6px 10px 0px 0;
}
.book-list-item {
  margin: 0px 10px 10px;
  padding: 10px;
  border: 1px solid #f1f1f1;
  box-shadow: 1px 1px 2px 0px #e5e5e5;
  box-sizing: border-box;
  border-radius: 5px;
}
.book-list-item h2 {
  margin: 0;
  display: inline-block;
}
.book-list-item button {
  float: right;
  margin-left: 5px;
}
.book-desc {
  font-size: 0.8rem;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
