import '../src/assets/variables.css'
import '../src/assets/style.css'
import App from './App.vue'
import router from './../router'
import {createApp} from "vue";
import '../scripts/errors.js'

createApp(App)
    .use(router)
    .mount('#app')
