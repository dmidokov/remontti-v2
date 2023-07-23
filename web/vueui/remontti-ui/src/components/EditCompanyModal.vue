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

    <ActionButton18 :title="getTranslations('UpdateCompany')" class="modal-horizontal-center-button"
                    :action="updateCompany" :dataid="companyId"/>

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
import {computed} from "vue";

export default {
  data() {
    return {
      companyName: "",
      companyHost: "",
      companyAdminName: "",
      companyAdminPassword: "",
      companyId: ""
    }
  },
  computed: {
    companyName: function () {
      let value = Object.hasOwn(this.companyData, 'CompanyName') ? this.companyData.CompanyName : ''
      this.companyName = value
      return value
    },
    companyHost: function () {
      let value = Object.hasOwn(this.companyData, 'HostName') ? this.companyData.HostName : ''
      this.companyHost = value
      return value
    },
    companyId: function () {
      return Object.hasOwn(this.companyData, 'CompanyId') ? this.companyData.CompanyId : ''
    },

  },
  name: "EditCompanyModal",
  components: {InputModal, ActionButton18, ActionButton24},
  props: ['translate', 'closeAction', 'companyData'],
  methods: {
    getTranslations(label) {
      return getTranslations(this.translate, label)
    },
    updateCompany: async function (event) {
      const companyData = {
        'company_name': this.companyName,
        'company_host': this.companyHost,
      }

      let response = await post("/api/v1/companies/" + event.target.getAttribute('data-id'), companyData)

      if (response.error == null) {
        if (response.data.status === "error") {
          this.errorMessage = response.data.message
        } else {
          document
              .getElementById('error-popup-block')
              .append(createSuccessBlock("Готово!", "Данные компании обновлены"))
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