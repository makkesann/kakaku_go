<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'App',
  data: function(){
    return {
      drink_fields: ["name", "ID", "DrinkGenreID"],
      drinks: [],
      genre_fields: ["name"],
      drink_genres: []
    }
  },

  // インスタンス作成時の処理
  created: function() {
      this.doFetchAllDrink()
      this.doFetchAllDrinkGenre()
  },

  methods: {
    // 全ての商品情報を取得する
    doFetchAllDrink() {
        axios.get('http://localhost:8082/drinks')
        .then(response => {
            if (response.status != 200) {
                throw new Error('レスポンスエラー')
            } else {
                var resultDrinks = response.data

                // サーバから取得した商品情報をdataに設定する
                this.drinks = resultDrinks
            }
        })
    },
    // 全てのジャンル情報を取得する
    doFetchAllDrinkGenre() {
        axios.get('http://localhost:8082/drink/genres')
        .then(response => {
            if (response.status != 200) {
                throw new Error('レスポンスエラー')
            } else {
                var resultDrinkGenres = response.data

                // サーバから取得した商品情報をdataに設定する
                this.drink_genres = resultDrinkGenres
            }
        })
    },
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
