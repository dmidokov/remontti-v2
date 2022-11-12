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
    this.translations = (await requests.get("/api/v1/translations/get?pages=loginpage")).data
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
          window.location.href = window.location.href
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
        <input type="password" id="password" v-model="password" :placeholder='getTranslations(this.translations,"PasswordFieldHeader")'>
      </div>
      <div class="sign-in-button">
        <button @click="auth(login, password)">
          {{ getTranslations(this.translations, "SignIn") }}
        </button>
        <div class="error-message">
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

.modal {
  width: 570px;
  height: 380px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  box-shadow: 0px 0px 16px 2px rgba(34, 60, 80, 0.24);
  border-radius: 10px;
}

.modal div {
  margin: 10px 10px;
}

.modal input {
  font-size: 18px;
  border-radius: 10px;
  outline: none;
  padding: 10px 15px;
  width: 200px;
  border: 1px solid rgba(0, 0, 0, 0.2);
  color: rgba(0, 0, 0, 0.7);
  background-color: var(--input-background);
}

button {
  background: var(--button-color-1);
  color: var(--button-text-color-1);
  border-radius: 10px;
  border: 0px solid transparent;
  outline: none;
}

button:hover {
  background: var(--button-color-1-light);
}

.error-message {
  position: absolute;
  top: 100%;
}

.sign-in-button {
  position: relative;
}
</style>