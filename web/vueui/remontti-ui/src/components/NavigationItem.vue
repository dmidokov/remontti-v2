<template>
  <div>
    <a :class="{ 'current-page':isCurrentPage}" :href="link">{{ getTranslate(label) }}</a>
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
  data() {
    return {
      isCurrentPage: window.location.pathname == this.link ? true : false,
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
  color: var(--link-hover-color);
}

div {
  margin-left: 10px;
  margin-right: 10px;
}

.current-page {
  font-weight: bold;
  color: var(--link-hover-color);
}
</style>