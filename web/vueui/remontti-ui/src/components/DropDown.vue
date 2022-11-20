<template>
  <div class="dropdown">
    <DropDownLine
        v-for="item in items"
        v-bind:title="item.title"
        v-bind:link="item.link"
        v-bind:action="item.action"
        v-bind:dataid="item.dataid"
    />
  </div>
</template>

<script>

import DropDownLine from "./DropDownLine.vue"

export default {
  name: "DropDown",
  components: {DropDownLine},
  props: ['items'],
  methods: {}
}
</script>

<style scoped>
.dropdown {
  display: flex;
  flex-direction: column;
  background: var(--drop-down-background);
  color: var(--drop-down-text-color);
  font-size: 18px;
  position: absolute;
  top: 115%;
  right: 0px;
  width: max-content;
  box-shadow: var(--box-shadow-offset-0);
  z-index: 1;
}


</style>