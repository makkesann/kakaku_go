import axios from 'axios'
import Vue from 'vue'
export default {
  namespaced: true,
  // 初期値
  state: {
    drinks: [],
    drink_genres: [],
  },
  getters: {
    getDrinks(state) {
      let arr = state.drinks
      return arr
    },
    getDrink: (state) => (id) => {
      return state.drinks.filter((drink) => drink.ID == id)
    },
    getDrinkGenres(state) {
      return state.drink_genres
    },
    getGenre: (state) => (drink_genre_id) => {
      return state.drink_genres.filter((drink_genre) => drink_genre.ID == drink_genre_id)
    },
  },

  actions: {
    doFetchAllDrink(context) {
      axios.get('http://localhost:8082/drinks')
      .then(response => {
        context.commit('setDrinks', response.data)
      })
      .catch(error => {
        // handle error
        console.log(error)
      })
    },
    // 全てのジャンル情報を取得する
    doFetchAllDrinkGenre(context) {
      axios.get('http://localhost:8082/drink/genres')
      .then(response => {
        context.commit('setDrinkGenre', response.data)
      })
      .catch(error => {
        // handle error
        console.log(error)
      })
    },

  },

  mutations: {
    // 初期化処理
    setDrinks(state, resultDrinks) {
      for (const i in resultDrinks){
        Vue.set(resultDrinks[i], 'favorite', false)
      }
      state.drinks = resultDrinks
    },
    setDrinkGenre(state, resultDrinks) {
      state.drink_genres = resultDrinks
    },
    addFavoriteDrink: function(state, ID){
      state.drinks.find((drink) => drink.ID == ID).favorite = true
    },
    deleteFavoriteDrink: function(state, ID){
      state.drinks.find((drink) => drink.ID == ID).favorite = false
    },
  }
}
