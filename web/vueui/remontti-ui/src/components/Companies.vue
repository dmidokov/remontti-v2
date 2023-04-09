<template>
  <div class="container">
    <div class="companies-table-header">
      {{ getTranslate("CompanyTableTitle") }}
    </div>
    <div class="companies-table">
      <CompanyItem
          v-bind:translate="translations"
          v-for="item in companies"
          v-bind:name="item.company"
          v-bind:host="item.host"
          v-bind:id="item.id"/>
    </div>
    <div class="companies-table-footer-line">
      <AddButton v-bind:action="showAddCompanyModal"/>
    </div>
    <AddCompanyModal
        v-show="addCompanyModalToggle"
        v-bind:translate="translations"
        v-bind:closeAction="closeAddCompanyModal"/>
    <EditCompanyModal
      v-show="editCompanyModalToggle"
      v-bind:translate="translations"
      v-bind:closeAction="closeEditCompanyModal"/>
  </div>
  <ErrorMessagePopupBlock :id="'error-popup-block'"/>
</template>

<script>
import * as requests from "../../scripts/requests.js";
import {getTranslations} from "../../scripts/translations.js";
import CompanyItem from "./CompanyItem.vue";
import AddButton from "./AddButton.vue";
import AddCompanyModal from "./AddCompanyModal.vue";
import ErrorMessagePopupBlock from "./ErrorMessagePopupBlock.vue";
import EditCompanyModal from "./EditCompanyModal.vue";

export default {
  name: "Companies",
  components: {
    AddCompanyModal,
    CompanyItem,
    AddButton,
    ErrorMessagePopupBlock,
    EditCompanyModal
  },
  data() {
    return {
      translations: {},
      companies: {},
      addCompanyModalToggle: false,
      editCompanyModalToggle: false
    }
  },
  async created() {
    await this.fetchCompanies()
    await this.fetchCompaniesTranslations()
  },
  methods: {
    getTranslate(label) {
      return getTranslations(this.translations, label)
    },
    showAddCompanyModal() {
      this.addCompanyModalToggle = true;
    },
    closeAddCompanyModal() {
      this.addCompanyModalToggle = false;
    },
    async fetchCompanies() {
      this.companies = (await requests.get("/api/v1/companies/get")).data;
    },
    async fetchCompaniesTranslations() {
      this.translations = (await requests.get("/api/v1/translations/get?pages=companies")).data;
    },
    closeEditCompanyModal(){
      this.editCompanyModalToggle = false
    },
    showEditCompanyModal() {
      this.editCompanyModalToggle = true;
    },
  },
}
</script>

<style scoped>

.companies-table {
  display: flex;
  flex-direction: column;
  min-height: 100px;
  width: 90vw;
  margin: auto;
}

.companies-table-header {
  width: 90vw;
  font-size: 24px;
  margin: 15px auto;
}

.container {
  margin-top: 10px;
}

.companies-table-footer-line {
  display: flex;
  width: 90vw;
  flex-direction: row-reverse;
  margin: auto;
}

</style>