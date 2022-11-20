<script setup>
import Header from '../src/components/Header.vue'
import Companies from "../src/components/Companies.vue";
</script>
<script>

import * as requests from "../scripts/requests.js";

export default {
  name: "Companies",
  data() {
    return {
      translations: {},
    }
  },
  async beforeCreate() {
    this.translations = (await requests.get("/api/v1/translations/get?pages=companies")).data
  }
}

</script>

<template>
  <Header v-bind:translations="translations"/>
  <Companies/>
</template>

<style scoped>

</style>

<!--<hollow-dots-spinner-->
<!--    :animation-duration="1000"-->
<!--    :dot-size="15"-->
<!--    :dots-num="3"-->
<!--    color="#ff1d5e"-->
<!--/>-->