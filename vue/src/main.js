import starportLibrary from '@starport/vue'
import { createApp } from 'vue'
import VueApexCharts from "vue3-apexcharts"
import axios from 'axios'
import VueAxios from 'vue-axios'

import App from './App.vue'
import router from './router'
import store from './store'

const app = createApp(App)

app.use(store)
app.use(VueApexCharts)
app.use(starportLibrary)
app.use(router)
app.use(VueAxios, axios)
app.mount('#app')
