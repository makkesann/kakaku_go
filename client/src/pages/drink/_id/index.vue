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
          <template v-slot:cell(Name)="{item}">
            {{ item.Name }}
            <p v-if="admin">削除</p>
            <b-icon v-if="item.favorite" @click="doDeleteFavoriteShop(item)" icon="star-fill" aria-hidden="true" variant="warning"></b-icon>
            <b-icon v-else @click="doAddFavoriteShop(item)" icon="star" aria-hidden="true" variant="warning"></b-icon>
          </template>
        </b-table>
      </b-row>
      <b-row>
        <h2>楽天での価格</h2>
        <b-table striped hover :items="rakuten.Items" :fields="rakuten_fields">
            <template v-slot:cell(商品名)="{item}">
              {{ item.itemName }}
            </template>
            <template v-slot:cell(価格)="{item}">
              {{ item.itemPrice }}円
            </template>
            <template v-slot:cell(画像)="{item}">
              <img :src="item.smallImageUrls[0]" alt="">
            </template>
        </b-table>
      </b-row>
      <img :src="rakuten.Items[0].smallImageUrls[0]" alt="">
    </b-container>
  </div>
</template>

<script>
// @ is an alias to /src
import axios from 'axios'
export default {
  name: 'home',
  data: function(){
    return {
      drink_fields: ["name", "ID", "DrinkGenreID"],
      // drinks: [],
      prices: [],
      // favorite_shops: [],
      rakuten: [],
      rakuten_fields:["商品名", "価格", "画像"]
    }
  },

  // インスタンス作成時の処理
  created: function() {
    this.doFetchPrices()
    this.rakutenapi()
  },

  computed: {
    drink() {
      return this.$store.getters["drink/getDrink"](this.$route.params.id)
    },
    favorite_shops_id() {
      return this.$store.getters["login/getFavoriteShop"]
    },
    favorite_shops(){
      return this.prices.filter((price) => price.favorite == true)
    },
    admin(){
      return this.$store.getters["login/getAdmin"]
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
      immediate: true,
      // deep: true,
      handler(){
        this.ReloadFavoriteShop(this.favorite_shops_id)
      }
    },
    drink:{
      // deep: true,
      immediate: true,
      handler(drink){
        this.rakutenapi(drink[0])
      }
    },
  },

  methods: {
    rakutenapi(drink){
      if (drink != undefined){
        let jan_code = 0
        if (drink.Jan == '0'){
          jan_code = this.drink[0].name
        }else{
          jan_code = this.drink[0].Jan
        }
        // console.log(jan_code)
        // console.log("わろた")
        axios.get('https://app.rakuten.co.jp/services/api/IchibaItem/Search/20170706?applicationId=1053675093689591225&formatVersion=2&sort=%2BitemPrice&hits=1&keyword=' + jan_code)
        .then(response => {
          // console.log(response.data)
          this.rakuten = response.data
        })
        .catch(error => {
          // handle error
          console.log(error)
        })
      }
    },
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
      for (const j in this.prices){
        this.$set(this.prices[j], 'favorite', false)
      }
      for (const i in favorite){
        if (this.prices.some((price) => price.ShopID == favorite[i].ShopID)){
          let arr = this.prices.filter((price) => price.ShopID == favorite[i].ShopID)
          for (const j in arr){
            this.$set(arr[j], 'favorite', true)
          }
        }
      }
      result.sort(function(a,b){
        if(a.price < b.price) return -1;
        if(a.price > b.price) return 1;
        return 0;
      });
    },
    doAddFavoriteShop(item){
      // item.favorite = true
      this.$store.dispatch('login/doAddFavoriteShop',item)
      if (this.$store.state.login.id != 0){
        // サーバへ送信するパラメータ
        const params = new URLSearchParams()
        params.append('shop_id', item.ShopID)
        params.append('user_id', this.$store.state.login.id)
        axios.post('http://54.65.204.164:8082/favorite_shop/add', params)
        .catch(error => {
          // handle error
          console.log(error)
        })
      }

    },
    doDeleteFavoriteShop(item){
      // item.favorite = false
      this.$store.dispatch('login/doDeleteFavoriteShop',item)
      if (this.$store.state.login.id != 0){
        const params = new URLSearchParams()
        params.append('shop_id', item.ShopID)
        params.append('user_id', this.$store.state.login.id)
        axios.post('http://54.65.204.164:8082/favorite_shop/delete', params)
        .catch(error => {
          // handle error
          console.log(error)
        })
      }
    }
  },
}
</script>
