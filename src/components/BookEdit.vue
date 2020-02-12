<template>
  <div id="bookEdit">
    <div v-if="book != null">
      <SimpleButton v-on:click="$router.push('/books')">Back</SimpleButton>
      <SavingIndicator v-bind:state="state" />
      <input placeholder="name" id="name" v-model="book.name" />
      <textarea
        placeholder="description"
        v-on:input="autoGrow"
        id="description"
        v-model="book.description"
      />
      <textarea placeholder="body" v-on:input="autoGrow" id="body" v-model="book.body" />
    </div>
    <LoadingIndicator v-else />
  </div>
</template>

<script>
// @ is an alias to /src
import SimpleButton from "@/components/SimpleButton.vue";
import LoadingIndicator from "@/components/LoadingIndicator.vue";
import SavingIndicator from "@/components/SavingIndicator.vue";

export default {
  name: "book-edit",
  props: ["id"],
  data: function() {
    return {
      state: 0,
      book: null,
      autoSaveTimeout: null
    };
  },
  components: {
    SimpleButton,
    LoadingIndicator,
    SavingIndicator
  },
  created() {
    this.getBook();
  },
  watch: {
    book: {
      deep: true,
      handler: function(newValue, oldValue) {
        if (oldValue == null) {
          return;
        }
        clearTimeout(this.autoSaveTimeout);
        this.state = 1;
        // console.log("set");
        this.autoSaveTimeout = setTimeout(async () => {
          let res;
          try {
            res = await fetch(
              `http://${this.$root.$data.state.baseURL}/api/book`,
              {
                method: "PUT",
                mode: "cors",
                credentials: "include", // include, *same-origin, omit
                headers: {
                  "Content-Type": "application/json"
                },
                body: JSON.stringify(this.book)
              }
            );
            if (res.status != 200) {
              console.log("non 200 status saving book", res);
              this.state = 3;
            } else {
              this.state = 2;
            }
          } catch (err) {
            console.log("error saving book", err);
            this.state = 3;
          }
          clearTimeout(this.autoSaveTimeout);
        }, 1000);
      }
    }
  },
  methods: {
    getBook: async function() {
      let data;
      let res;
      try {
        res = await fetch(
          `http://${this.$root.$data.state.baseURL}/api/book/${this.id}`,
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
        console.log(err);
      }
      if (res.status == 200) {
        console.log(data);
        this.book = data;
      }
    },
    autoGrow: function(event) {
      // console.log(event);
      event.target.style.height = "5px";
      event.target.style.height = event.target.scrollHeight + "px";
    }
  }
};
</script>

<style>
textarea:focus,
input:focus {
  outline: none;
}

#name {
  border-style: solid;
  border-width: thin;
  border-radius: 3px;
  border-color: #e1e1e1;
  box-shadow: 1px 1px 3px 1px #e5e5e5;
  font-size: 2rem;
  margin-top: 10px;
  display: block;
  width: 100%;
  padding: 5px 10px;
  box-sizing: border-box;
}

textarea {
  resize: none;
  border-radius: 3px;
  border-color: #e1e1e1;
  box-shadow: 1px 1px 3px 1px #e5e5e5;
  margin-top: 10px;
  display: block;
  width: 100%;
  padding: 10px;
  box-sizing: border-box;
  overflow: hidden;
}

#description {
  min-height: 50px;
}

#body {
  min-height: 300px;
}
</style>