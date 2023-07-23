<template>
  <div>
    <button @click="handler" v-bind:data-id="dataid">
      {{ title }}
    </button>
  </div>
</template>

<script>
export default {
  name: "ActionButton18",
  props: ['action', 'title', 'dataid'],
  methods: {
    handler(event) {
      this.action(event)
    }
  }
}
</script>

<style scoped>
div {
  display: flex;
}

button {
  font-size : 18px;
}

button:hover {
  background: var(--button-color-1-light);
}

button:active {
  box-shadow: var(--box-shadow-offset-0)
}

</style>