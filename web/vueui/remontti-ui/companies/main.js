import { createApp } from 'vue'
import '../src/assets/variables.css'
import '../src/assets/style.css'
import App from './App.vue'
import router from './../router'

createApp(App)
    .use(router)
    .mount('#app')