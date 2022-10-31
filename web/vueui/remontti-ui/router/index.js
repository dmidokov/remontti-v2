import {createRouter, createWebHistory} from 'vue-router'

import Home from '../src/App.vue'
import Login from '../login/App.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            component: Home,
        },
        {
            path: "/login",
            component: Login
        }
    ]
})

export default router