<template>
  <div class="modal">
    <ActionButton24 :title="'X'" class="close-modal-button" :action="this.closeAction"/>

  </div>

</template>

<script>

import ActionButton24 from "./ActionButton24.vue";


export default {
  name: "EditCompanyModal",
  components: {ActionButton24},
  props: ['translate', 'closeAction'],
}
</script>

<style scoped>

</style>