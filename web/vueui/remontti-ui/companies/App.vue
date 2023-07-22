<script setup>
import Header from '../src/components/Header.vue'
import CompaniesContainer from "../src/components/CompaniesContainer.vue";
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
    this.translations = (await requests.get("/api/v1/translations/companies")).data
  }
}

</script>

<template>
  <Header/>
  <CompaniesContainer/>
</template>

<style scoped>

</style>

<!--<hollow-dots-spinner-->
<!--    :animation-duration="1000"-->
<!--    :dot-size="15"-->
<!--    :dots-num="3"-->
<!--    color="#ff1d5e"-->
<!--/>-->