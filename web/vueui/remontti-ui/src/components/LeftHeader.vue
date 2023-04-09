<template>
  <div class="header-part">
    {{ companyName }}
  </div>
</template>

<script>
import * as requests from "../../scripts/requests.js";

export default {
  name: "LeftHeaderBlock",
  data() {
    return {
      companyName : ''
    }
  },
  async beforeCreate() {
    this.companyName = (await requests.get("/api/v1/companies/get?act=getCurrentCompanyName")).data
  }
}
</script>

<style scoped>

</style>