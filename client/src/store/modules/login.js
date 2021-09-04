
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
    // 全ての商品情報を取得する
    async doFetchAllDrink({ commit }) {
      await axios.get('https://....com/todos')
      .then(response => {
        if (response.status != 200) {
            throw new Error('レスポンスエラー')
        } 
      })
      commit('SetDrinks', response.data)
    },
  },
  mutations: {
    SetDrinks: (state, drinks) => (state.drink = drinks)
  }
}
