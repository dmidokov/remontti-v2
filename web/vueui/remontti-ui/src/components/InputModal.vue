<template>
  <div>
    <input
        v-bind:value="value"
        v-on:input="$emit('changed', $event.target.value)"
        type="text"
        :placeholder="placeholder">
  </div>
</template>

<script>
export default {
  name: "InputModal",
  props: ['placeholder', 'value']
}
</script>

<style scoped>
div {
  display: flex;
  justify-content: center;
}
</style>