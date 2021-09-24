import axios from 'axios'

export default {
  namespaced: true, // 追加
  state: {
    id: 0,
    favorite_drink: [],
    favorite_shop: [],
    admin: false
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
    getAdmin(state) {
      return state.admin
    },
  },
  actions: {
    doSetID(context, user_id) {
      context.commit('setID', user_id)
    },
    doAdminLogIn(context, admin_state) {
      context.commit('AdminLogIn', admin_state)
    },
    // ユーザー情報の取得
    doFetchFavoriteDrinks(context, user_id) {
      if (user_id != 0){
        return axios.get('http://localhost:8082/user/'+ user_id + '/favorite_drink')
        .then(response => {
          context.commit('setFavoriteDrinks', response.data)
        })
        .catch(error => {
          // handle error
          console.log(error)
        })
      }
    },
    doFetchFavoriteShops(context, user_id) {
      if (user_id != 0){
        return axios.get('http://localhost:8082/user/'+ user_id + '/favorite_shop')
        .then(response => {
          context.commit('setFavoriteShops', response.data)
        })
        .catch(error => {
          // handle error
          console.log(error)
        })
      }
    },
    doAddFavoriteDrink(context, item){
      context.commit('addFavoriteDrink', item)
      context.commit("drink/addFavoriteDrink",item.ID,{root:true})
    },
    doDeleteFavoriteDrink(context, item){
      context.commit('deleteFavoriteDrink', item)
      context.commit("drink/deleteFavoriteDrink",item.ID,{root:true})
    },
    doAddFavoriteShop(context, item){
      context.commit('addFavoriteShop', item)
      // context.commit("shop/addFavoriteShop",item.ID,{root:true})
    },
    doDeleteFavoriteShop(context, item){
      context.commit('deleteFavoriteShop', item)
      // context.commit("shop/deleteFavoriteShop",item.ID,{root:true})
    },
  },
  mutations: {
    setFavoriteDrinks: function(state, favorite_drinks) {
      state.favorite_drink = favorite_drinks
    },
    setFavoriteShops: (state, favorite_shops) => (state.favorite_shop = favorite_shops),
    addFavoriteDrink: function(state, item){
      state.favorite_drink.push({"DrinkID": item.ID, "UserID": state.id})
    },
    deleteFavoriteDrink: function(state, item){
      state.favorite_drink = state.favorite_drink.filter(function(value){
        return value.DrinkID != item.ID || value.UserID != state.id
      })
    },
    addFavoriteShop: function(state, item){
      state.favorite_shop.push({"ShopID": item.ShopID, "UserID": state.id})
    },
    deleteFavoriteShop: function(state, item){
      state.favorite_shop = state.favorite_shop.filter(function(value){
        return value.ShopID != item.ShopID || value.UserID != state.id
      })
    },
    setID: (state, id) => (state.id = id),
    AdminLogIn: (state, admin_state) => (state.admin = admin_state),
  }
}
