import axios from 'axios'

export default {
  namespaced: true, // 追加
  state: {
    id: 0,
    favorite_drink: [],
    favorite_shop: []
  },
  getters: {
    getFavoriteDrink(state) {
      return state.favorite_drink
    },
    getFavoriteDrinkTriger(state) {
      return state.favorite_drink_triger
    },
    getFavoriteShop(state) {
      return state.favorite_shop
    },
    getID(state) {
      return state.id
    },
  },
  actions: {
    doSetID(context, user_id) {
      context.commit('setID', user_id)
    },
    // ユーザー情報の取得
    doFetchFavoriteDrinks(context, user_id) {
      return axios.get('http://localhost:8082/user/'+ user_id + '/favorite_drink')
      .then(response => {
              context.commit('setFavoriteDrinks', response.data)
      })
      .catch(function (error) {
        // handle error
        console.log(error)
      })
    },
    doFetchFavoriteShops(context, user_id) {
      return axios.get('http://localhost:8082/user/'+ user_id + '/favorite_shop')
      .then(response => {
        context.commit('setFavoriteShops', response.data)
      })
      .catch(function (error) {
        // handle error
        console.log(error)
      })
    },
    idplus(context) {
      context.commit('idplus')
    }
  },
  mutations: {
    setFavoriteDrinks: function(state, favorite_drinks) {
      state.favorite_drink = favorite_drinks
    },
    setFavoriteShops: (state, favorite_shops) => (state.favorite_shop = favorite_shops),
    setID: (state, id) => (state.id = id),
    idplus(state){
      state.id++
    }
  }
}
