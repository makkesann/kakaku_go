import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'
import drink from '@/store/modules/drink.js'


Vue.use(Vuex)

// ストアの定義
export default new Vuex.Store({
  // ステート、ミューテーションなどをここに記載

  modules: {
    drink: drink,
 },
  plugins: [
    createPersistedState({
      paths: ['login']
  })],
})
