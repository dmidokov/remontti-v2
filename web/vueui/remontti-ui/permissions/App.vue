<script>
import * as requests from "../scripts/requests.js";
import Header from '../src/components/Header.vue'
import PermissionsContainer from "../src/components/PermissionsContainer.vue";

export default {
  name: "Permissions",
  components: {Header, PermissionsContainer},
  data() {
    return {
      translations: {},
    }
  },
  async beforeCreate() {
    this.translations = (await requests.get("/api/v1/translations/permissions")).data
  }
}
</script>

<template>
  <Header/>
  <PermissionsContainer/>
</template>

<style scoped>

</style>