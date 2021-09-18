<template>
  <div class="home">
    <b-container>
      <b-row>
        <b-table striped hover :items="drink"></b-table>
      </b-row>
      <b-row>
        <h2>お気に入りの店</h2>
        <b-table v-if="favorite_shops.length != 0" striped hover :items="favorite_shops">
        </b-table>
        <router-link :to="{path: this.$route.path +'/add'}">価格の追加</router-link>
        <b-table striped hover :items="prices">
          <template v-slot:cell(favorite)="{item}">
            <b-icon v-if="item.favorite" @click="doDeleteFavoriteDrink(item)" icon="star-fill" aria-hidden="true" variant="warning"></b-icon>
            <b-icon v-else @click="doAddFavoriteDrink(item)" icon="star" aria-hidden="true" variant="warning"></b-icon>
          </template>
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
      prices: [],
      favorite_shops: []
    }
  },

  // インスタンス作成時の処理
  created: function() {
    this.doFetchPrices()
  },

  computed: {
    drink() {
      return this.$store.getters["drink/getDrink"](this.$route.params.id)
    },
    favorite_shops_id() {
      return this.$store.getters["login/getFavoriteShop"]
    },
    serched_favorite_shops(){
      if (this.genre_id == 0){
        return this.favorite_shops
      } else {
        return this.favorite_shops.filter((shop) => shop.id == this.genre_id)
      }
    },
  },
  watch: {
    favorite_shops_id:{
      immediate: true,
      // deep: true,
      handler(favorite){
        this.ReloadFavoriteShop(favorite)
      }
    },
    prices:{
      // deep: true,
      immediate: true,
      handler(){
        this.ReloadFavoriteShop(this.favorite_shops_id)
      }
    },
  },

  methods: {

    doFetchPrices() {
      let drink_id = this.$route.params.id
      axios.get('http://54.65.204.164:8082/drinks/' + drink_id + '/prices')
      .then(response => {
        const resultDrinkPrices = response.data
        // console.log(resultDrinkPrices)
        // サーバから取得した商品情報をdataに設定する
      for (const i in resultDrinkPrices){
        this.$set(resultDrinkPrices[i], 'favorite', false)
      }
        this.prices = resultDrinkPrices
      })
      .catch(error => {
        // handle error
        console.log(error)
      })
    },
    ReloadFavoriteShop(favorite){
      let result = []
      for (const i in favorite){
        if (this.prices.some((price) => price.ShopID == favorite[i].ShopID)){
          let arr = this.prices.filter((price) => price.ShopID == favorite[i].ShopID)
          for (const j in arr){
            this.$set(arr[j], 'favorite', true)
            result.push(arr[j])
          }
        }
      }
      result.sort(function(a,b){
        if(a.price < b.price) return -1;
        if(a.price > b.price) return 1;
        return 0;
      });
      this.favorite_shops = result
    }
  }
}
</script>
