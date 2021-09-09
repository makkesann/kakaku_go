<template>
  <div class="home">
    <b-container>
      <b-row>
        <b-table striped hover :items="drink" :fields="drink_fields"></b-table>
      </b-row>
      <b-row>
        <b-table striped hover :items="prices">
        </b-table>
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
      drink_fields: ["name", "ID", "DrinkGenreID"],
      // drinks: [],
      prices: []
    }
  },

  // インスタンス作成時の処理
  created: function() {
  //     this.doFetchDrink()
      this.doFetchPrices()
  },

  computed: {
    // drinks() {
    //   return this.$store.getters["drink/getDrinks"];
    // },
    drink() {
      return this.$store.getters["drink/getDrink"](this.$route.params.id);
    },
    // drink_genres() {
    //   return this.$store.getters["drink/getDrinkGenres"];
    // }
  },

  methods: {
    doFetchPrices() {
      let drink_id = this.$route.params.id
        axios.get('http://localhost:8082/drinks/' + drink_id + '/prices')
        .then(response => {
            if (response.status != 200) {
                throw new Error('レスポンスエラー')
            } else {
                var resultDrinkPrices = response.data

                // サーバから取得した商品情報をdataに設定する
                this.prices = resultDrinkPrices
            }
        })
    },
  }
}
</script>
