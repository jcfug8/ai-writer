
<template>
  <div class="autocomplete">
    <textarea
      ref="auto-area"
      v-on:keyup="getSuggestions"
      v-on:keydown="navSuggestions"
      v-on:input="input"
      v-on:click="hideAuto"
      v-model="dataText"
    ></textarea>
    <span
      ref="auto-container"
      class="autocomplete-container"
      v-if="autoShowing"
      v-bind:style="spanStyle"
    >
      <ul v-if="!loading">
        <li
          v-for="(s, i) in suggestions"
          v-on:mouseover="active = i"
          v-bind:key="i"
          v-on:click="selectSuggestion($event, s)"
          v-bind:class="[i == active ? 'active' : '', 'suggestion']"
        >{{ s }}</li>
      </ul>
      <div v-else>Loading...</div>
    </span>
  </div>
</template>

<script>
export default {
  name: "autocomplete",
  props: ["body"],
  data() {
    return {
      dataText: this.body,
      timeout: null,
      loading: true,
      spanStyle: {
        position: "absolute",
        top: null,
        left: null
      },
      active: null,
      suggestions: [
        "Hello, how are you?",
        "I'm a little teapot short and stout.",
        "This is another suggestion. Use it wisely."
      ]
    };
  },
  methods: {
    getSuggestions: function(event) {
      console.log("get", event);
      if (
        (event.keyCode == 13 && this.autoShowing) ||
        (event.keyCode == 40 && this.autoShowing) ||
        (event.keyCode == 38 && this.autoShowing)
      ) {
        // up down or enter
        return false;
      } else if (event.keyCode == 81 && event.ctrlKey && event.shiftKey) {
        console.log("big");
        this.loadLargeSuggestion();
      } else if (event.keyCode == 81 && event.ctrlKey) {
        console.log("small");
        this.showAuto();
        this.loadSuggestions();
      } else if (event.keyCode != 17 && event.keyCode != 16) {
        this.hideAuto();
      }
    },
    navSuggestions: function(event) {
      // console.log("nav", event);
      if (event.keyCode == 13 && this.autoShowing) {
        // enter
        event.preventDefault();
        event.stopImmediatePropagation();
        let e = { target: this.$refs["auto-area"] };
        this.selectSuggestion(e);
        return false;
      } else if (event.keyCode == 40 && this.autoShowing) {
        // down
        event.preventDefault();
        event.stopImmediatePropagation();
        if (this.active == null) {
          this.active = 0;
        } else {
          this.active = (this.active + 1) % this.suggestions.length;
        }
        return false;
      } else if (event.keyCode == 38 && this.autoShowing) {
        // up
        event.preventDefault();
        event.stopImmediatePropagation();
        if (this.active == null) {
          this.active = this.suggestions.length - 1;
        } else {
          this.active =
            (this.active - 1 + this.suggestions.length) %
            this.suggestions.length;
        }
        return false;
      }
    },
    hideAuto: function() {
      this.spanStyle.top = null;
      this.spanStyle.left = null;
    },
    showAuto: function() {
      let input = this.$refs["auto-area"];
      let selectionPoint = input.selectionEnd;
      const { offsetLeft: inputX, offsetTop: inputY } = input;
      // create a dummy element that will be a clone of our input
      const div = document.createElement("div");
      // get the computed style of the input and clone it onto the dummy element
      const copyStyle = getComputedStyle(input);
      for (const prop of copyStyle) {
        div.style[prop] = copyStyle[prop];
      }
      // we need a character that will replace whitespace when filling our dummy element if it's a single line <input/>
      const swap = ".";
      const inputValue =
        input.tagName === "INPUT"
          ? input.value.replace(/ /g, swap)
          : input.value;
      // set the div content to that of the textarea up until selection
      const textContent = inputValue.substr(0, selectionPoint);
      // set the text content of the dummy element div
      div.textContent = textContent;
      if (input.tagName === "TEXTAREA") div.style.height = "auto";
      // if a single line input then the div needs to be single line and not break out like a text area
      if (input.tagName === "INPUT") div.style.width = "auto";
      // create a marker element to obtain caret position
      const span = document.createElement("span");
      // give the span the textContent of remaining content so that the recreated dummy element is as close as possible
      span.textContent = inputValue.substr(selectionPoint) || ".";
      // append the span marker to the div
      div.appendChild(span);
      // append the dummy element to the body
      input.parentElement.prepend(div);
      // get the marker position, this is the caret position top and left relative to the input
      const { offsetLeft: spanX, offsetTop: spanY } = span;
      // lastly, remove that dummy element
      // NOTE:: can comment this out for debugging purposes if you want to see where that span is rendered
      input.parentElement.removeChild(div);
      // return an object with the x and y of the caret. account for input positioning so that you don't need to wrap the input
      // console.log("inputX", inputX);
      // console.log("spanX", spanX);
      // console.log("inputY", inputY);
      // console.log("spanY", spanY);
      this.spanStyle.left = inputX + spanX + "px";
      this.spanStyle.top = inputY + spanY + 20 + "px";
      // console.log(this.spanStyle);
      setTimeout(() => {
        if (this.$refs["auto-container"]) {
          var rt =
            window.innerWidth -
            this.$refs["auto-container"].getBoundingClientRect().right;
          // console.log(rt);
          if (rt < 0) {
            this.spanStyle.left = inputX + spanX + rt + "px";
          }
        }
      }, 1);
    },
    loadSuggestions: async function() {
      this.active = null;
      this.loading = true;
      let seed_text = this.$refs["auto-area"].value.slice(-255);
      let res;
      try {
        res = await fetch(
          `http://${this.$root.$data.state.baseURL}/api/getsimple`,
          {
            method: "POST",
            mode: "cors",
            credentials: "include", // include, *same-origin, omit
            headers: {
              "Content-Type": "application/json"
            },
            body: JSON.stringify({ seed_text })
          }
        );
        if (res.status == 200) {
          console.log("200 status getting suggestions", res);
          let body = await res.json();
          console.log("suggestion body", body);
          this.suggestions = body.messages;
          this.loading = false;
        } else {
          this.hideAuto();
        }
      } catch (err) {
        this.hideAuto();
        console.log("error getting suggestions", err);
      }
    },
    loadLargeSuggestion: async function() {
      let input = this.$refs["auto-area"];
      input.disabled = true;
      let seed_text = this.$refs["auto-area"].value.slice(-255);
      let res;
      try {
        res = await fetch(
          `http://${this.$root.$data.state.baseURL}/api/getlarge`,
          {
            method: "POST",
            mode: "cors",
            credentials: "include", // include, *same-origin, omit
            headers: {
              "Content-Type": "application/json"
            },
            body: JSON.stringify({ seed_text })
          }
        );
        if (res.status == 200) {
          console.log("200 status getting suggestions", res);
          let body = await res.json();
          console.log("suggestion body", body);
          input.disabled = false;
          let index = input.selectionEnd;
          this.dataText =
            this.dataText.slice(0, index) +
            body.message +
            this.dataText.slice(index);
          // this.$emit("update:body", this.dataText);
          this.hideAuto();
          this.input();
        }
      } catch (err) {
        console.log("error getting suggestions", err);
      }
    },
    selectSuggestion: function(event) {
      let suggestion = this.suggestions[this.active];
      let input = this.$refs["auto-area"];
      if (event.stopImmediatePropagation) {
        event.stopImmediatePropagation();
      }
      let index = input.selectionEnd;
      this.dataText =
        this.dataText.slice(0, index) + suggestion + this.dataText.slice(index);
      // this.$emit("update:body", this.dataText);
      this.hideAuto();
      this.input();
    },
    input: function() {
      this.$emit("update:body", this.dataText);
      this.$emit("onInput", this.$refs["auto-area"]);
    }
  },
  computed: {
    autoShowing: function() {
      return this.spanStyle.top != null && this.spanStyle.left != null;
    }
  },
  watch: {},
  mounted() {},
  destroyed() {}
};
</script>

<style scoped>
.autocomplete {
  position: relative;
}

.autocomplete-container {
  font-size: 0.7rem;
  white-space: nowrap;
  padding: 2px 0;
  border: 1px solid lightgrey;
  border-radius: 2px;
  background-color: white;
  box-shadow: 1px 1px 3px 1px #e5e5e5;
}

.autocomplete-container ul {
  list-style-type: none;
  margin-block-start: 0;
  margin-block-end: 0;
  padding-inline-start: 0;
}

.autocomplete-container ul li:not(:last-child) {
  border-bottom: 1px #c4c4c4 dotted;
}

.autocomplete-container ul li {
  padding: 0 4px;
  display: block;
  cursor: pointer;
  white-space: pre;
}

.suggestion.active {
  background-color: lightgray;
}
</style>