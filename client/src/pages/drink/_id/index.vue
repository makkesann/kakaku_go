<template>
  <div class="product_detail">
    <b-container>
      <div class="text-left">
        <router-link to="/drink">商品一覧に戻る</router-link>
      </div>
      <div class="box">
        <h2 class="text-left">{{drink[0].name}}</h2><hr>
        <b-row>
          <b-col cols="4" class="py-4 px-2 out-img">
            <div id="Imgbox" class="mx-auto">
              <img v-lazy="'https://kakaku-go-product.s3.ap-northeast-1.amazonaws.com/large/' + drink[0].Image">
            </div>
          </b-col>
          <b-col cols="8" class="py-2 px-2 w-100">
            <b-row>
              <h2 v-if="prices.length ==0">価格が登録されていません</h2>
              <b-col cols="4" v-if="prices.length !=0">
                <h4 class="text-left mb-3">最安価格：{{ prices[0].Price }}円</h4>
              </b-col>
              <b-col cols="8" v-if="prices.length !=0">
                <div class="pr-2 w-60 d-inline-block">
                  <h5 class="text-left">最安ショップ：{{ prices[0].Name }}</h5>
                </div>
                <div class="w-40 d-inline-block">
                  <b-button>このショップを探す</b-button>
                </div>
              </b-col>
            </b-row>
            <a :href="rakuten.Items[0].itemUrl" target="_blank" v-if="this.rakuten!=null">
              <div class="rakuten">
                <h4 class="text-left">楽天の参考価格:</h4>
                <b-row>
                  <b-col cols="2">
                    <img :src="rakuten.Items[0].smallImageUrls[0]">
                  </b-col>
                  <b-col cols="10">
                    <h6 class="text-left">{{ rakuten.Items[0].itemName }}</h6>
                    <h4 class="text-left">{{ rakuten.Items[0].itemPrice }}円</h4>
                  </b-col>
                </b-row>
              </div>
            </a>
          </b-col>
        </b-row>

      </div>
      <b-row>
        <h2 v-if="favorite_shops.length != 0">お気に入りの店</h2>
        <b-table v-if="favorite_shops.length != 0" striped hover :items="favorite_shops">
        </b-table>
        <router-link :to="{path: this.$route.path +'/add'}">価格の追加</router-link>
        <b-table striped hover :items="prices">
          <template v-slot:cell(Name)="{item}">
            {{ item.Name }}
            <!-- <p v-if="admin">削除</p> -->
            <b-icon v-if="item.favorite" @click="doDeleteFavoriteShop(item)" icon="star-fill" aria-hidden="true" variant="warning"></b-icon>
            <b-icon v-else @click="doAddFavoriteShop(item)" icon="star" aria-hidden="true" variant="warning"></b-icon>
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
  data: function(){
    return {
      drink_fields: ["name", "ID", "DrinkGenreID"],
      // drinks: [],
      prices: [],
      // favorite_shops: [],
      rakuten: [],
      rakuten_fields:["商品名", "価格", "画像"],
      google: []
    }
  },

  // インスタンス作成時の処理
  created: function() {
    this.doFetchPrices()
    this.rakutenapi()
    // this.googleapi()
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
        if (String(drink.Jan).length == 8|| String(drink.Jan).length == 13 ){
          jan_code = drink.Jan
          axios.get('https://app.rakuten.co.jp/services/api/IchibaItem/Search/20170706?applicationId=1053675093689591225&formatVersion=2&sort=%2BitemPrice&hits=1&keyword=' + jan_code)
          .then(response => {
            // console.log(response.data)
            this.rakuten = response.data
          })
          .catch(error => {
            // handle error
            console.log(error)
          })
        }else{
          this.rakuten = null
        }
        // console.log(jan_code)
        // console.log("わろた")

      }
    },
    googleapi(){
      axios.get('https://maps.googleapis.com/maps/api/place/nearbysearch/json?key=AIzaSyB8JS5diZ8zIEUkoapu9qp_fVAVihF1C_M&location=35.6987769,139.76471&radius=300&language=ja&keyword=公園OR広場OR駅')
      .then(response => {
        // console.log(response.data)
        this.google = response.data
      })
      .catch(error => {
        // handle error
        console.log(error)
      })
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
    },

  },
}
</script>
