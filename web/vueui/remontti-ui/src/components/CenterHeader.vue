<template>
  <div class="header-part header-center">
    <NavigationItem
        v-bind:translate="translations"
        v-for="item in navigation"
        v-bind:label="item.label"
        v-bind:link="item.link"
    />
  </div>
</template>

<script>
import NavigationItem from "./NavigationItem.vue";
import {computed} from "vue";

export default {
  name: "CenterHeaderBlock",
  components: {NavigationItem},
  data() {
    return {
      translations: computed(() => this.translations),
    }
  },
  props: [
    'navigation',
    'translations'
  ],
}
</script>

<style scoped>

.header-center {
  display: flex;
  justify-content: center;
}

</style>