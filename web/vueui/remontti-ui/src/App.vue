<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
import Header from './components/Header.vue'

</script>
<script>

import * as requests from "../scripts/requests.js";

export default {
  name: "Main",
  data() {
    return {
      translations: {},
    }
  },
  async beforeCreate() {
    this.translations = (await requests.get("/api/v1/translations/get?pages=mainpage")).data
  }
}

</script>

<template>
  <Header v-bind:translations="translations"/>
</template>

<style scoped>

</style>
