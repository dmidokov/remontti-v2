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
import * as requests from "../../scripts/requests.js";

export default {
  name: "CenterHeaderBlock",
  components: {NavigationItem},
  data() {
    return {
      translations : {}
    }
  },
  props: [
    'navigation'
  ],
  async beforeCreate() {
    this.translations = (await requests.get("/api/v1/translations/get?pages=navigation")).data
  }
}
</script>

<style scoped>

.header-center {
  display: flex;
  justify-content: center;
}

</style>