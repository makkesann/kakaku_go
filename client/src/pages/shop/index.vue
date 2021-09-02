<template>
  <div class="home">
    <b-container>
      <b-row>
        <b-col cols="2">
          <b-table striped hover :items="drink_genres" :fields="genre_fields"></b-table>
        </b-col>
        <b-col cols="10">
          <b-table striped hover :items="drinks" :fields="drink_fields"></b-table>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
// @ is an alias to /src
import axios from 'axios'
export default {
  name: 'home',
  components: {
  },
  data: function(){
    return {
      drink_fields: ["ID", "name", "DrinkGenreID"],
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
