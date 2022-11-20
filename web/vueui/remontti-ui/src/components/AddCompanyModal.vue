<template>
  <div class="modal">
    <ActionButton24 :title="'X'" class="close-modal-button" :action="this.closeAction"/>

    <InputModal
        v-bind:value="companyName"
        @changed="updateCompanyNameValue"
        :placeholder="getTranslations('CompanyName')"/>

    <InputModal
        v-bind:value="companyHost"
        @changed="updateCompanyHostValue"
        :placeholder="getTranslations('CompanyHost')"/>

    <ActionButton18 :title="getTranslations('AddCompany')" class="modal-horizontal-center-button" :action="addCompany"/>
  </div>
</template>

<script>
import ActionButton24 from "./ActionButton24.vue";
import ActionButton18 from "./ActionButton18.vue";
import InputModal from "./InputModal.vue";
import {getTranslations} from "../../scripts/translations.js";
import {post} from "../../scripts/requests.js";
import {compile, createApp, createVNode, render} from "vue";
import {createErrorBlock} from "../../scripts/errors.js";

export default {
  data() {
    return {
      companyName: "",
      companyHost: ""
    }
  },
  name: "AddCompanyModal",
  props: ['translate', 'closeAction'],
  components: {InputModal, ActionButton18, ActionButton24},
  methods: {
    getTranslations(label) {
      return getTranslations(this.translate, label)
    },
    async addCompany(event) {

      const companyData = {
        'companyName': this.companyName,
        'companyHost': this.companyHost
      }

      let response = await post("/api/v1/companies/add", companyData)

      console.log(response)

      if (response.error == null) {
        if (response.data.status === "error") {
          this.errorMessage = response.data.message
          console.log("ddasdasd")
        } else {
          console.log("ok")
        }
      } else {

        document
            .getElementById('error-popup-block')
            .append(createErrorBlock("Ошибка", "В ходе выполнения запроса произошла ошибка, попробуйте еще раз"))

      }
    },
    updateCompanyNameValue(value) {
      this.companyName = value;
    },
    updateCompanyHostValue(value) {
      this.companyHost = value;
    }
  }
}
</script>

<style scored>

</style>