import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import router from './router/index.js'
import store from './store/index.js'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import Paginate from 'vuejs-paginate'
import { BootstrapVue, BootstrapVueIcons } from 'bootstrap-vue'
import "@/assets/style/global.scss"
import VueLazyload from 'vue-lazyload'
import ja from 'vee-validate/dist/locale/ja.json'
import {
  ValidationObserver,
  ValidationProvider,
  extend,
  localize,
} from "vee-validate"
import * as rules from "vee-validate/dist/rules";
import * as VueGoogleMaps from 'vue2-google-maps'

Vue.component('paginate', Paginate)
localize({
  ja: { messages: { ...ja.messages } },
});
localize('ja');

extend('my_alpha_dash', {
  // エラーメッセージを設定する
  message: '半角英数字とハイフン(-)、アンダーバー(_)のみ使用することが出来ます',
  validate(value) {
      if (value.match(/^[A-Za-z0-9_-]+$/)) {
          return true;
      }
  }
});

Object.keys(rules).forEach(rule => {
  extend(rule, rules[rule]);
});
Vue.component("ValidationObserver", ValidationObserver)
Vue.component("ValidationProvider", ValidationProvider)


Vue.use(BootstrapVue)
Vue.use(BootstrapVueIcons)

Vue.config.productionTip = true
Vue.prototype.$axios = axios
Vue.use(VueLazyload, {
  preLoad: 1.3,
  error: 'https://dummyimage.com/64x64/ccc/999.png&text=Not+Found',
  loading: 'https://dummyimage.com/64x64/dcdcdc/999.png&text=Now loading',
  attempt: 1
})

Vue.use(VueGoogleMaps, {
  load: {
    key: '{GoogleMapのAPIキー}',
    libraries: 'places', 
  },
})

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')

module.exports = {
  devServer: {
      proxy: 'https://kakaku-real-store.tk:8082/',
      disableHostCheck: true,
  }
}