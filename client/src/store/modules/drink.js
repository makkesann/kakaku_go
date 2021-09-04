import axios from 'axios'

export default {
  namespaced: true, // 追加
  state: {
    drink: []
  },
  getters: {
    getDrink: (state) => {
      return state.drink
   } 
  },
  actions: {
    // 全ての商品情報を取得する
    async doFetchAllDrink() {
      await axios.get('http://localhost:8082/drinks')
      .then(response => {
        if (response.status != 200) {
            throw new Error('レスポンスエラー')
        } else {
            var resultDrinks = response.data
            return resultDrinks
          }
      })
      
    },
  },
  mutations: {
    SetDrinks(state, payload) {
      state.drink = payload
    }
  }
}
