import axios from 'axios'

export default {
  namespaced: true, // 追加
  state: {
    id: 0,
    // user: ["test"]
  },
  getters: {
    getUser(state) {
      return state.user
    },
    getID(state) {
      return state.id
    },
  },
  actions: {
    // ユーザー情報の取得
    doFetchUser(context) {
      return axios.get('http://localhost:8082/user/'+ this.state.id)
      .then(response => {
          if (response.status != 200) {
              throw new Error('レスポンスエラー')
          } else {
              context.commit('setUser', response.data)
          }
      })
    },
  },
  mutations: {
    SetDrinks: (state, drinks) => (state.drink = drinks)
  }
}
