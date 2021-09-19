import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import router from './router/index.js'
import store from './store/index.js'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import { BootstrapVue, BootstrapVueIcons } from 'bootstrap-vue'
import "./assets/style/global.css"

Vue.use(BootstrapVue)
Vue.use(BootstrapVueIcons)

Vue.config.productionTip = true
Vue.prototype.$axios = axios
// axios.defaults.baseURL = 'http://54.65.204.164.164:8082'

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')

module.exports = {
  devServer: {
      proxy: 'http://54.65.204.164.164:8081/',
      disableHostCheck: true,
  }
}