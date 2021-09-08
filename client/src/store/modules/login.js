import axios from 'axios'

export default {
  namespaced: true, // 追加
  state: {
    user: []
  },
  getters: {
    user(state) {
      return state.user
    }
  },
  actions: {
    // ユーザー情報の取得
    doFetchUser(context, id) {
      return axios.get('http://localhost:8082/user/'+ id)
      .then(response => {
          if (response.status != 200) {
              throw new Error('レスポンスエラー')
          } else {
              // let resultDrinks = response.data
              context.commit('setDrinks', response.data)
          }
      })
    },
  },
  mutations: {
    SetDrinks: (state, drinks) => (state.drink = drinks)
  }
}
