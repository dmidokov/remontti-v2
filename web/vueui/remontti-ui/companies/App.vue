<script setup>
import Header from '../src/components/Header.vue'
import { HollowDotsSpinner } from 'epic-spinners'
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
  <hollow-dots-spinner
      :animation-duration="1000"
      :dot-size="15"
      :dots-num="3"
      color="#ff1d5e"
  />
</template>

<style scoped>

</style>
