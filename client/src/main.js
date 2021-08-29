import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import router from './router/index.js'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import BootstrapVue from 'bootstrap-vue'
Vue.use(BootstrapVue)


Vue.config.productionTip = true
Vue.prototype.$axios = axios

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')

module.exports = {
  devServer: {
      proxy: 'http://localhost:8081/',
      disableHostCheck: true,
     }
   };