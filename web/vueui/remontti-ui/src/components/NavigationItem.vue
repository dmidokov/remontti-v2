<template>
  <div>
    <a :href="link">{{ getTranslate(label) }}</a>
  </div>
</template>

<script>
import {getTranslations} from "../../scripts/translations.js";

export default {
  name: "NavigationItem",
  props: ['label', 'link', 'translate'],
  methods: {
    getTranslate(label) {
      return getTranslations(this.translate, label)
    }
  },
}
</script>

<style scoped>
a {
  color: inherit;
  font-weight: inherit;
}

a:hover {
  color : var(--link-hover-color);
}

div {
  margin-left: 10px;
  margin-right: 10px;
}
</style>