<script>

import LeftHeader from "./LeftHeader.vue"
import CenterHeader from "./CenterHeader.vue";
import RightHeader from "./RightHeader.vue";
import * as requests from "../../scripts/requests.js";

export default {
  name: "Header",
  components: {RightHeader, CenterHeader, LeftHeader},
  data() {
    return {
      translations: {},
      navigation: {}
    }
  },
  async beforeCreate() {
    this.navigation = (await requests.get("/api/v1/navigation")).data
    this.translations = (await requests.get("/api/v1/translations/navigation")).data
  }
}
</script>

<template>
  <div class="header">
    <LeftHeader/>
    <CenterHeader
        v-bind:translations="translations" v-bind:navigation="navigation"/>
    <RightHeader v-bind:translations="translations"/>
  </div>
</template>

<style scoped>

.header {
  border-bottom: 2px solid var(--light-gray);
  min-height: 50px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-part {
  min-width: 100px;
}

</style>