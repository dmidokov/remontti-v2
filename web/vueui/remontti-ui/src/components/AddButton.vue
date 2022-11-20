<template>
  <button @click="handler">
    +
  </button>
</template>

<script>
export default {
  name: "AddButton",
  props: ['action'],
  methods: {
    handler(event) {
      this.action(event)
    }
  }
}
</script>

<style scoped>
button {
  background: var(--button-color-1);
  color: var(--button-text-color-1);
  border: 0px solid transparent;
  font-size: 24px;
  padding: 1px 10px;
}

button:hover {
  background: var(--button-color-1-light);
}

button:active {
  box-shadow: var(--box-shadow-offset-0)
}

</style>