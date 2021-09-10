import axios from 'axios'

export default {
  namespaced: true, // 追加
  state: {
    id: 1,
    favorite_drink: [],
    favorite_shop: []
  },
  getters: {
    getFavoriteDrink(state) {
      return state.favorite_drink
    },
    getFavoriteShop(state) {
      return state.favorite_shop
    },
    getID(state) {
      return state.id
    },
  },
  actions: {
    // ユーザー情報の取得
    doFetchFavoriteDrinks(context, user_id) {
      return axios.get('http://localhost:8082/user/'+ user_id + '/favorite_drink')
      .then(response => {
          if (response.status != 200) {
              throw new Error('レスポンスエラー')
          } else {
              context.commit('setFavoriteDrinks', response.data)
          }
      })
    },
    doFetchFavoriteShops(context, user_id) {
      return axios.get('http://localhost:8082/user/'+ user_id + '/favorite_shop')
      .then(response => {
          if (response.status != 200) {
              throw new Error('レスポンスエラー')
          } else {
              context.commit('setFavoriteShops', response.data)
          }
      })
    },
    idplus(context) {
      context.commit('idplus')
    }
  },
  mutations: {
    setFavoriteDrinks: (state, favorite_drinks) => (state.favorite_drink = favorite_drinks),
    setFavoriteShops: (state, favorite_shops) => (state.favorite_shop = favorite_shops),
    setID: (state, id) => (state.id = id),
    idplus(state){
      state.id++
    }
  }
}
