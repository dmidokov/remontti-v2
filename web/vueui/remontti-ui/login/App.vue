<script>

import * as requests from '../scripts/requests.js'
import {getTranslations} from '../scripts/translations.js'

export default {
  name: "Login",
  data() {
    return {
      translations: [],
      login: "",
      password: "",
      errorMessage: ""
    }
  },
  async beforeCreate() {
    this.translations = (await requests.get("/api/v1/translations/loginpage")).data
  },
  methods: {
    async auth(loginField, passwordField) {
      if (loginField === "") {
        this.errorMessage = getTranslations(this.translations, "LoginIsEmpty")
        return
      }
      if (passwordField == "" && passwordField.length < 5) {
        this.errorMessage = getTranslations(this.translations,"PasswordIsEmpty")
        return
      }

      const user = {
        login: loginField, password: passwordField
      }

      let response = await requests.post("/login", user)

      if (response.error == null) {
        if (response.data.status === "error") {
          this.errorMessage = response.data.message
        } else {
          window.location.href = "/"
        }
      }
    },
    getTranslations(t,l) {
      return getTranslations(t,l)
    }

  }

}

</script>

<template>
  <div>
    <div class="modal">
      <div>
        <input type="text" id="login" v-model="login" :placeholder='getTranslations(this.translations, "LoginFieldHeader")'>
      </div>
      <div>
        <input @keydown.enter="auth(login, password)" type="password" id="password" v-model="password" :placeholder='getTranslations(this.translations,"PasswordFieldHeader")'>
      </div>
      <div class="sign-in-button">
        <button @click="auth(login, password)">
          {{ getTranslations(this.translations, "SignIn") }}
        </button>
        <div class="error-message-login">
          {{ errorMessage }}
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

* {
  outline: none;
}

div {
  display: flex;
  justify-content: center;
  display: flex;
}

.error-message-login {
  position: absolute;
  top: 100%;
}

.sign-in-button {
  position: relative;
}
</style>