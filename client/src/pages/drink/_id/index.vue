<template>
  <div class="product_detail">
    <b-container class="pc">
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
              <b-col cols="8" v-if="prices.length !=0" class="cheapest d-table">
                <div class="pr-2 d-table-cell">
                  <h5 class="text-left">最安ショップ：{{ prices[0].Name }}</h5>
                </div>
                <div class="d-table-cell" id="serch-shop">
                  <b-button @click="googleapi(prices[0].Name)" href="#map-box">このショップを探す</b-button>
                </div>
              </b-col>
            </b-row>
            <a :href="rakuten.Items[0].itemUrl" target="_blank" v-if="this.rakuten!=null">
              <div class="rakuten">
                <h4 class="text-left">楽天の参考価格</h4>
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
        <b-table v-if="favorite_shops.length != 0" striped hover :items="favorite_shops" :fields="price_fields" class="prices-table">
          <template v-slot:cell(お店)="{item}">
            <span>{{ item.Name }}&nbsp;&nbsp;</span>
           <b-icon v-if="item.favorite" @click="doDeleteFavoriteShop(item)" icon="star-fill" aria-hidden="true" variant="warning"></b-icon>
            <b-icon v-else @click="doAddFavoriteShop(item)" icon="star" aria-hidden="true" variant="warning"></b-icon>
            <!-- <p v-if="admin">削除</p> -->
          </template>
          <template v-slot:cell(価格)="{item}" class="table-price">
            <span>{{ item.Price }}円</span>
            <!-- <p v-if="admin">削除</p> -->
          </template>
          <template v-slot:cell(お店を探す)="{item}">
            <b-button @click="googleapi(item.Name)" href="#map-box">このショップを探す</b-button>
          </template>
        </b-table>
        <router-link :to="{path: this.$route.path +'/add'}">価格の追加</router-link>
        <b-table striped hover :items="prices" :fields="price_fields" class="prices-table">
          <template v-slot:cell(お店)="{item}">
            <span>{{ item.Name }}&nbsp;&nbsp;</span>
           <b-icon v-if="item.favorite" @click="doDeleteFavoriteShop(item)" icon="star-fill" aria-hidden="true" variant="warning"></b-icon>
            <b-icon v-else @click="doAddFavoriteShop(item)" icon="star" aria-hidden="true" variant="warning"></b-icon>
            <!-- <p v-if="admin">削除</p> -->
          </template>
          <template v-slot:cell(価格)="{item}" class="table-price">
            <span>{{ item.Price }}円</span>
            <!-- <p v-if="admin">削除</p> -->
          </template>
          <template v-slot:cell(お店を探す)="{item}">
            <b-button @click="googleapi(item.Name)" href="#map-box">このショップを探す</b-button>
          </template>
        </b-table>
      </b-row>
    </b-container>
    <b-container class="sp">
      <div class="text-left">
        <router-link to="/drink">商品一覧に戻る</router-link>
      </div>
      <div class="my-3">
        <h3 class="text-left">{{drink[0].name}}</h3><hr class="m-0">
        <b-row>
          <b-col cols="5" class="py-4 px-2 out-img">
            <div id="Imgbox" class="mx-auto">
              <img v-lazy="'https://kakaku-go-product.s3.ap-northeast-1.amazonaws.com/large/' + drink[0].Image">
            </div>
          </b-col>
          <b-col cols="7" class="py-2 px-2 w-100">
            <p v-if="prices.length ==0">価格が登録されていません</p>
            <p v-if="prices.length !=0" class="text-left mb-1">最安価格：{{ prices[0].Price }}円</p>
            <div v-if="prices.length !=0">
              <p class="text-left mb-3">最安ショップ：<br>{{ prices[0].Name }}</p>
            </div>
            <div  id="serch-shop">
              <b-button @click="googleapi(prices[0].Name)" href="#map-box">このショップを探す</b-button>
            </div>
          </b-col>
        </b-row>
        <b-row>          
          <a :href="rakuten.Items[0].itemUrl" target="_blank" v-if="this.rakuten!=null">
            <div class="rakuten">
              <h5 class="text-left">楽天の参考価格</h5>
              <b-row>
                <b-col cols="4" class="p-0">
                  <img :src="rakuten.Items[0].smallImageUrls[0]">
                </b-col>
                <b-col cols="8" class="p-0">
                  <p class="text-left mb-1">{{ rakuten.Items[0].itemName }}</p>
                  <p class="text-left mb-0">{{ rakuten.Items[0].itemPrice }}円</p>
                </b-col>
              </b-row>
            </div>
          </a>
        </b-row>
      </div>

      <b-row>
        <h2 v-if="favorite_shops.length != 0">お気に入りの店</h2>
        <b-table v-if="favorite_shops.length != 0" striped hover :items="favorite_shops" :fields="price_fields" class="prices-table">
          <template v-slot:cell(お店)="{item}">
            <span>{{ item.Name }}&nbsp;&nbsp;</span>
           <b-icon v-if="item.favorite" @click="doDeleteFavoriteShop(item)" icon="star-fill" aria-hidden="true" variant="warning"></b-icon>
            <b-icon v-else @click="doAddFavoriteShop(item)" icon="star" aria-hidden="true" variant="warning"></b-icon>
            <!-- <p v-if="admin">削除</p> -->
          </template>
          <template v-slot:cell(価格)="{item}" class="table-price">
            <span>{{ item.Price }}円</span>
            <!-- <p v-if="admin">削除</p> -->
          </template>
          <template v-slot:cell(お店を探す)="{item}">
            <b-button @click="googleapi(item.Name)" href="#map-box">このショップを探す</b-button>
          </template>
        </b-table>
        <router-link :to="{path: this.$route.path +'/add'}">価格の追加</router-link>
        <b-table striped hover :items="prices" :fields="price_fields" class="prices-table">
          <template v-slot:cell(お店)="{item}">
            <span>{{ item.Name }}&nbsp;&nbsp;</span>
           <b-icon v-if="item.favorite" @click="doDeleteFavoriteShop(item)" icon="star-fill" aria-hidden="true" variant="warning"></b-icon>
            <b-icon v-else @click="doAddFavoriteShop(item)" icon="star" aria-hidden="true" variant="warning"></b-icon>
            <!-- <p v-if="admin">削除</p> -->
          </template>
          <template v-slot:cell(価格)="{item}" class="table-price">
            <span>{{ item.Price }}円</span>
            <!-- <p v-if="admin">削除</p> -->
          </template>
          <template v-slot:cell(お店を探す)="{item}">
            <b-button @click="googleapi(item.Name)" href="#map-box">このショップを探す</b-button>
          </template>
        </b-table>
      </b-row>
    </b-container>
    <b-container>
      <div id="map-box" v-if="map_show">
        <b-alert v-if="map_error" variant="danger" show>
          {{map_error}}
        </b-alert>
        <div id="map"></div>
      </div>
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
      price_fields: ["お店", "価格", "お店を探す"],
      // drinks: [],
      prices: [],
      // favorite_shops: [],
      rakuten: [],
      rakuten_fields:["商品名", "価格", "画像"],
      map_error: null,
      map_show: false,
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
            this.rakuten = response.data
          })
          .catch(error => {
            // handle error
            console.log(error)
          })
        }else{
          this.rakuten = null
        }
      }
    },
    googleapi(shop_name){
      this.map_show = true
      var self = this
      var google = window.google
      // Geolocation APIに対応している
      if (navigator.geolocation) {
        // 現在地を取得
        navigator.geolocation.getCurrentPosition(
          // 取得成功した場合
          function(position) {
            var mapLatLng = new google.maps.LatLng(position.coords.latitude, position.coords.longitude)
            // マップオプションを変数に格納
            var mapOptions = {
              zoom : 14,          // 拡大倍率
              center : mapLatLng  // 緯度・経度
            };
            // マップオブジェクト作成
            var map = new google.maps.Map(
              document.getElementById("map"), // マップを表示する要素
              mapOptions         // マップオプション
            );
            // Places APIのnearbySearchを使用する。
            let placeService = new google.maps.places.PlacesService(map)
            placeService.nearbySearch(
              {
                location: new google.maps.LatLng(position.coords.latitude, position.coords.longitude),
                keyword: shop_name,
                rankby: "distance",
                radius: 5000
              },
              function(results, status) {
                if (results.length ==0){
                  self.map_error = "付近にこのお店は存在しません"
                } else{
                  if (status == google.maps.places.PlacesServiceStatus.OK) {
                    var marker = new google.maps.Marker({
                      map: map,
                      position: results[0].geometry.location,
                      id: results[0].place_id
                    })
                    var infoWindow = new google.maps.InfoWindow({ // 吹き出しの追加
                      content: '<div class="map_balloon">' + results[0].name + '</div>' // 吹き出しに表示する内容
                    });
                    marker.addListener('click', function() { // マーカーをクリックしたとき
                      infoWindow.open(map, marker); // 吹き出しの表示
                    });
                }
                }
              }
            )
          },
          // 取得失敗した場合
          function(error) {
            // エラーメッセージを表示
            switch(error.code) {
              case 1: // PERMISSION_DENIED
                alert("位置情報の利用が許可されていません");
                break;
              case 2: // POSITION_UNAVAILABLE
                alert("現在位置が取得できませんでした");
                break;
              case 3: // TIMEOUT
                alert("タイムアウトになりました");
                break;
              default:
                alert("その他のエラー(エラーコード:"+error.code+")");
                break;
            }
          }
        );
      // Geolocation APIに対応していない
      } else {
        alert("この端末では位置情報が取得できません");
      }
    },
    doFetchPrices() {
      let drink_id = this.$route.params.id
      axios.get('https://kakaku-real-store.tk:8082/drinks/' + drink_id + '/prices')
      .then(response => {
        const resultDrinkPrices = response.data
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
        axios.post('https://kakaku-real-store.tk:8082/favorite_shop/add', params)
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
        axios.post('https://kakaku-real-store.tk:8082/favorite_shop/delete', params)
        .catch(error => {
          // handle error
          console.log(error)
        })
      }
    },

  },
}
</script>
