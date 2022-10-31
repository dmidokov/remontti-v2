import {createApp} from 'vue'
// TODO: вынести стили тоже
import '../src/assets/style.css'
import '../src/assets/variables.css'
import App from './App.vue'
import router from './../router'


createApp(App)
    .use(router)
    .mount('#app')

// TODO: вынести кусок с роутером в отдельный файл, так как в нескольких местах