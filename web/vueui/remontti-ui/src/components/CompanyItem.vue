<template>
  <div class="company-line" v-bind:data-id="id">
    <div class="company-name">
      {{ name }}
    </div>
    <div class="company-host">
      {{ host }}
    </div>
    <div class="company-actions">
      <IconDropdown
          v-bind:src="companyActionsIcon.src"
          v-bind:items="companyActionsIcon.items"/>
    </div>
  </div>
</template>

<script>
import IconDropdown from "./IconDropdown.vue";
import {computed} from "vue";
import {getTranslations} from "../../scripts/translations.js";
import {del} from "../../scripts/requests.js";
import {createSuccessBlock} from "../../scripts/success.js";
import {createErrorBlock} from "../../scripts/errors.js";

export default {
  name: "CompanyItem",
  props: ['name', 'host', 'id', 'translate'],
  components: {IconDropdown},
  data() {
    return {
      companyActionsIcon: {
        src: new URL("../assets/account.png", import.meta.url).href,
        items: [
          {
            "title": computed(() => getTranslations(this.translate, "EditCompany")),
            "action": this.editCompany,
            "dataid": this.id
          },
          {
            "title": computed(() => getTranslations(this.translate, "DeleteCompany")),
            "action":this.deleteCompany,
            "dataid": this.id
          },
        ]
      },
    }
  },
  methods: {
    editCompany(event) {
      this.$parent.showEditCompanyModal(event.target.getAttribute("data-id"))
      console.log("EDIT COMPANY:: ", event.target.getAttribute("data-id"))
    },
    async deleteCompany(event) {
      const companyData = {
        'company_id': Number(event.target.getAttribute("data-id")),
      }

      let response = await del("/api/v1/companies/delete", companyData)

      if (response.error == null) {
        if (response.data.status === "error") {
          this.errorMessage = response.data.message
          document
              .getElementById('error-popup-block')
              .append(createErrorBlock("Ошибка", response.data.message))
        } else {
          document
              .getElementById('error-popup-block')
              .append(createSuccessBlock("Готово!", "Компания удалена"))
          this.$parent.fetchCompanies()
        }
      } else {
        document
            .getElementById('error-popup-block')
            .append(createErrorBlock("Ошибка", "В ходе выполнения запроса произошла ошибка, попробуйте еще раз"))
      }

      console.log("DELETE COMPANY:: ", event.target.getAttribute("data-id"))
    }
  }
}
</script>

<style scoped>
.company-line {
  display: flex;
  text-align: center;
  align-items: center;
  padding: 10px 0px 10px 10px;
}

.company-line:nth-child(1) {
  border-top: 2px solid var(--table-first-line-top-border-color);
}

.company-line:nth-child(even) {
  background: var(--table-even-line-bg);
  color: var(--table-even-line-color);
}

.company-line:nth-child(even):hover {
  background: var(--table-even-line-bg-hover);
  color: var(--table-even-line-color-hover);
}

.company-line:nth-child(odd) {
  background: var(--table-odd-line-bg);
  color: var(--table-odd-line-color)
}

.company-line:nth-child(odd):hover {
  background: var(--table-odd-line-bg-hover);
  color: var(--table-odd-line-color-hover)
}

.company-line:hover {
  background: rgba(10, 10, 10, 0.2)
}

.company-name {
  display: flex;
  width: 40vw;

}

.company-host {
  display: flex;
  width: 40vw;

}

.company-actions {
  width: 20vw;
  position: relative;
}

</style>