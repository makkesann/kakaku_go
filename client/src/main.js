import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'


Vue.config.productionTip = false
Vue.prototype.$axios = axios

new Vue({
  render: h => h(App),
}).$mount('#app')

module.exports = {
  devServer: {
      proxy: 'http://localhost:8081/',
      disableHostCheck: true,
     }
   };