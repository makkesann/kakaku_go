import axios from 'axios'
export default {
  namespaced: true,
  // 初期値
  state: {
    drinks: [],
    drink_genres: []
  },
  getters: {
    getDrinks(state) {
      return state.drinks
    },
    getDrink: (state) => (id) => {
      return state.drinks.filter((drink) => drink.ID == id)
    },
    // getDrink(state){
    //   return state.drinks.filter(drink => drink.ID == 1)
    // },
    getDrinkGenres(state) {
      return state.drink_genres
    }
  },

  actions: {
    doFetchAllDrink(context) {
      return axios.get('http://localhost:8082/drinks')
      .then(response => {
          if (response.status != 200) {
              throw new Error('レスポンスエラー')
          } else {
              // let resultDrinks = response.data
              context.commit('setDrinks', response.data)
          }
      })
    },
    // 全てのジャンル情報を取得する
    doFetchAllDrinkGenre(context) {
      axios.get('http://localhost:8082/drink/genres')
      .then(response => {
          if (response.status != 200) {
              throw new Error('レスポンスエラー')
          } else {
             context.commit('setDrinkGenre', response.data)
          }
      })
    },
  },

  mutations: {
    // 初期化処理
    setDrinks(state, resultDrinks) {
      state.drinks = resultDrinks
    },
    setDrinkGenre(state, resultDrinks) {
      state.drink_genres = resultDrinks
    }
  }
}
