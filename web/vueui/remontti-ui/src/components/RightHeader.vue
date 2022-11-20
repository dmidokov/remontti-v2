<template>
  <div class="right-header header-part">
    <IconDropdown
        v-bind:src="settingsIcon.src"
        v-bind:items="settingsIcon.items"/>

    <IconDropdown
        v-bind:src="accountIcon.src"
        v-bind:items="accountIcon.items"/>
  </div>
</template>

<script>
import IconDropdown from "./IconDropdown.vue";
import {getTranslations} from "../../scripts/translations.js";
import {computed} from "vue";

export default {
  name: "RightHeader",
  components: {IconDropdown},
  props: [
    'translations'
  ],
  data() {
    return {
      accountIcon: {
        src: new URL("../assets/account.png", import.meta.url).href,
        items: [
          {
            "title": computed(() => getTranslations(this.translations, "Account"))
          },
          {
            "title": computed(() => getTranslations(this.translations, "Logout")),
            "link": "/logout"
          }
        ]
      },
      settingsIcon: {
        src: new URL("../assets/settings.png", import.meta.url).href,
        items: [
          {
            "title": computed(() => getTranslations(this.translations, "Settings"))
          },
        ]
      },
    }
  }
}
</script>

<style scoped>
.right-header {
  display: flex;
  flex-direction: row;
}

.right-header div {
  margin-left: 10px;
}
</style>