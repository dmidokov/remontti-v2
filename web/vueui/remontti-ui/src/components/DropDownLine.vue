<template>
  <div @click.stop class="dropdown-line">
    <a v-if="link !== undefined" :href=link>{{ title }}</a>
    <span v-else>{{ title }}</span>
  </div>
</template>

<script>
export default {
  name: "DropDownLine",
  props: ['title', 'link']
}
</script>

<style scoped>
.dropdown-line {
  font-size: 18px;
  padding: 6px 20px;
  cursor: pointer;
  color: var(--drop-down-text-color);
  display: flex;
  justify-content: center;
  text-align: right;
}

.dropdown-line:hover {
  font-size: 18px;
  padding: 6px 20px;
  background: var(--drop-down-item-hover-color);
  color: var(--drop-down-text-hover-color);
}

a {
  color: inherit;
  font-weight: inherit;
}


</style>