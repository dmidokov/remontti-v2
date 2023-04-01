<template>
  <div>
    <input
        v-bind:value="value"
        v-on:input="$emit('changed', $event.target.value)"
        :type="type"
        :placeholder="placeholder">
  </div>
</template>

<script>
export default {
  name: "InputModal",
  props: ['placeholder', 'value', 'type'],
  data() {
    return {
      type: this.type? this.type : "text"
    }
  }
}
</script>

<style scoped>
div {
  display: flex;
  justify-content: center;
}
</style>