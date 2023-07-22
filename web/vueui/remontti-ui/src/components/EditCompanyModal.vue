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

    <InputModal
        v-bind:value="companyAdminName"
        @changed="updateCompanyAdminNameValue"
        :placeholder="getTranslations('CompanyAdminName')"/>

    <InputModal
        v-bind:value="companyAdminPassword"
        v-bind:type="'password'"
        @changed="updateCompanyPasswordValue"
        :placeholder="getTranslations('CompanyAdminPassword')"/>

    <ActionButton18 :title="getTranslations('AddCompany')" class="modal-horizontal-center-button" :action="addCompany"/>

  </div>

</template>

<script>

import {getTranslations} from "../../scripts/translations.js";
import {post} from "../../scripts/requests.js";
import {createSuccessBlock} from "../../scripts/success.js";
import {createErrorBlock} from "../../scripts/errors.js";
import ActionButton24 from "./ActionButton24.vue";
import ActionButton18 from "./ActionButton18.vue";
import InputModal from "./InputModal.vue";
import Companies from "./CompaniesContainer.vue";

export default {
  data() {
    return {
      companyName: "",
      companyHost: "",
      companyAdminName: "",
      companyAdminPassword: ""
    }
  },
  name: "EditCompanyModal",
  components: {InputModal, ActionButton18, ActionButton24},
  props: ['translate', 'closeAction'],
  methods: {
    getTranslations(label) {
      return getTranslations(this.translate, label)
    },
    addCompany: async function (event) {
      const companyData = {
        'company_name': this.companyName,
        'company_host': this.companyHost,
        'admin_name': this.companyAdminName,
        'admin_password': this.companyAdminPassword
      }

      let response = await post("/api/v1/companies", companyData)

      if (response.error == null) {
        if (response.data.status === "error") {
          this.errorMessage = response.data.message
        } else {
          document
              .getElementById('error-popup-block')
              .append(createSuccessBlock("Готово!", "Новая компания добавлена"))
          this.$parent.fetchCompanies();
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
    },
    updateCompanyAdminNameValue(value) {
      this.companyAdminName = value;
    },
    updateCompanyPasswordValue(value) {
      this.companyAdminPassword = value;
    }
  }
}
</script>

<style scoped>

</style>