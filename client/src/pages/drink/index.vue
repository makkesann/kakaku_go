<template>
  <div class="drink">
    <b-container>
      <div class="search-box">
        <h2>どのような飲み物をお探しですか</h2>
        <b-form-input v-model="keyword" placeholder="商品名を入れてください"></b-form-input>
      </div>

      <h2 v-if="serched_favorite_drinks.length != 0">お気に入りの商品</h2>
      <b-row class="pc">
        <b-col cols="2">
          <div class="text-left genre_box mb-4">
            <h6 class="mb-0">ジャンルで絞る</h6><hr class="mb-3 mt-2">
            <p class="hover mb-2" @click="ChangeGenreID(0)">リセット</p>
            <div v-for="genre in drink_genres" v-bind:key="genre.id">
              <p @click="ChangeGenreID(genre.ID)" class="hover mb-2">{{ genre.name }}</p>
            </div>
          </div>
          <div class="text-left genre_box">
            <h6 class="mb-0">内容量で絞る</h6><hr class="mb-3 mt-2">
            <p class="hover mb-2" @click="ChangeQuantity(0)">リセット</p>
            <p class="hover mb-2" @click="ChangeQuantity(1)">~200ml</p>
            <p class="hover mb-2" @click="ChangeQuantity(2)">200ml~700ml</p>
            <p class="hover mb-2" @click="ChangeQuantity(3)">700ml~1L</p>
            <p class="hover mb-2" @click="ChangeQuantity(4)">1L~</p>
          </div>
        </b-col>
        <b-col cols="10">
          <div v-if="keyword != null && keyword !=''">
            <h2>検索結果</h2>
            <b-table v-if="keyworddrinks.length != 0" striped hover :items="keyworddrinks" :fields="drink_fields" class="drinks-table">
              <template v-slot:cell(画像)="{item}">
                <div class="test2 mx-auto">
                  <img v-lazy="'https://kakaku-go-product.s3.ap-northeast-1.amazonaws.com/small/' + item.Image">
                </div>
              </template>
              <template v-slot:cell(商品名)="{item}">
                <router-link :to="{name:'drink-id',params:{id: item.ID}}">{{ item.name }}</router-link>
                <!-- <p v-if="admin">削除</p> -->
              </template>
              <template v-slot:cell(内容量) ="{item}">
                <p v-if="item.Quantity!=0">{{ item.Quantity }}ml</p>
                <p v-else>- ml</p>
              </template>
              <template v-slot:cell(お気に入り)="{item}">
                <b-icon v-if="item.favorite" @click="doDeleteFavoriteDrink(item)" icon="star-fill" aria-hidden="true" variant="warning"></b-icon>
                <b-icon v-else @click="doAddFavoriteDrink(item)" icon="star" aria-hidden="true" variant="warning"></b-icon>
              </template>
            </b-table>
            <h3 v-else>該当する商品がありません</h3>
          </div>
          <!-- <p v-if="admin">商品追加</p> -->
          <b-table v-if="serched_favorite_drinks.length != 0" hover :items="serched_favorite_drinks" :fields="drink_fields" class="drinks-table">
            <template v-slot:cell(画像)="{item}">
              <div class="test2 mx-auto">
                <img v-lazy="'https://kakaku-go-product.s3.ap-northeast-1.amazonaws.com/small/' + item.Image">
              </div>
            </template>
            <template v-slot:cell(商品名)="{item}">
              <router-link :to="{name:'drink-id',params:{id: item.ID}}">{{ item.name }}</router-link>
              <!-- <p v-if="admin">削除</p> -->
            </template>
            <template v-slot:cell(内容量) ="{item}">
              <p v-if="item.Quantity!=0">{{ item.Quantity }}ml</p>
              <p v-else>- ml</p>
            </template>
            <template v-slot:cell(お気に入り)="{item}">
              <b-icon v-if="item.favorite" @click="doDeleteFavoriteDrink(item)" icon="star-fill" aria-hidden="true" variant="warning"></b-icon>
              <b-icon v-else @click="doAddFavoriteDrink(item)" icon="star" aria-hidden="true" variant="warning"></b-icon>
            </template>
          </b-table>
          <b-table hover :items="getItems" :fields="drink_fields" class="drinks-table">
            <template v-slot:cell(画像)="{item}">
              <div class="test2 mx-auto">
                <img v-lazy="'https://kakaku-go-product.s3.ap-northeast-1.amazonaws.com/small/' + item.Image">
              </div>
            </template>
            <template v-slot:cell(商品名)="{item}">
              <router-link :to="{name:'drink-id',params:{id: item.ID}}">{{ item.name }}</router-link>
              <!-- <p v-if="admin">削除</p> -->
            </template>
            <template v-slot:cell(内容量) ="{item}">
              <p v-if="item.Quantity!=0">{{ item.Quantity }}ml</p>
              <p v-else>- ml</p>
            </template>
            <template v-slot:cell(お気に入り)="{item}">
              <b-icon v-if="item.favorite" @click="doDeleteFavoriteDrink(item)" icon="star-fill" aria-hidden="true" variant="warning"></b-icon>
              <b-icon v-else @click="doAddFavoriteDrink(item)" icon="star" aria-hidden="true" variant="warning"></b-icon>
            </template>
          </b-table>
        </b-col>
      </b-row>
      <b-row class="sp drinks-table">
        <b-table v-if="serched_favorite_drinks.length != 0" striped hover :items="serched_favorite_drinks">
        </b-table>
        <b-table hover thead-class="d-none" :items="serched_drinks" :fields="drink_fields" class="drinks-table">
          <template v-slot:cell(画像)="{item}">
            <div class="test2 mx-auto">
              <img v-lazy="'https://kakaku-go-product.s3.ap-northeast-1.amazonaws.com/small/' + item.Image">
            </div>
          </template>
          <template v-slot:cell(商品名)="{item}">
            <router-link :to="{name:'drink-id',params:{id: item.ID}}">{{ item.name }}</router-link>
            <!-- <p v-if="admin">削除</p> -->
          </template>
          <template v-slot:cell(内容量) ="{item}">
            <p v-if="item.Quantity!=0">{{ item.Quantity }}ml</p>
            <p v-else>- ml</p>
          </template>
          <template v-slot:cell(お気に入り)="{item}">
            <b-icon v-if="item.favorite" @click="doDeleteFavoriteDrink(item)" icon="star-fill" aria-hidden="true" variant="warning"></b-icon>
            <b-icon v-else @click="doAddFavoriteDrink(item)" icon="star" aria-hidden="true" variant="warning"></b-icon>
          </template>
        </b-table>
      </b-row>
      <b-row>
        <paginate
          class="mx-auto page-nation"
          :page-count="getPageCount"
          :page-range="3"
          :margin-pages="2"
          :click-handler="clickCallback"
          :prev-text="'<'"
          :next-text="'>'"
          :container-class="'pagination'"
          :page-class="'page-item'"
          :page-link-class="'page-link'"
          :prev-class="'page-item'"
          :prev-link-class="'page-link'"
          :next-class="'page-item'"
          :next-link-class="'page-link'"
          :first-last-button="true"
          :first-button-text="'<<'"
          :last-button-text="'>>'">
        </paginate>
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
  data(){
    return {
      keyword: null,
      drink_fields: ["画像", "商品名", "内容量", "お気に入り"],
      genre_id: 0,
      favorite_drinks: [],
      quantity: 0,
      parPage: 10,
      currentPage: 1,
      // drinks:[]
    }
  },
  computed: {
    drinks() {
      return this.$store.getters["drink/getDrinks"]
    },
    admin(){
      return this.$store.getters["login/getAdmin"]
    },
    drink_genres() {
      return this.$store.getters["drink/getDrinkGenres"]
    },
    serched_drinks(){
      var result = this.drinks
      if (this.genre_id != 0){
        result = result.filter((drink) => drink.DrinkGenreID == this.genre_id)
      }
      if (this.quantity == 0){
        return result
      }
      else if (this.quantity == 1){
        return result.filter((drink) => (drink.Quantity >0 && drink.Quantity <= 200))
      }
      else if (this.quantity == 2){
        return result.filter((drink) => (drink.Quantity >200 && drink.Quantity <=700))
      }
      else if (this.quantity == 3){
        return result.filter((drink) => (drink.Quantity >700 && drink.Quantity <=1000))
      }
      else if (this.quantity == 4){
        return result.filter((drink) => drink.Quantity >1000)
      }
      else{
        return result
      }

    },
    favorite_drinks_id() {
      return this.$store.getters["login/getFavoriteDrink"]
    },
    serched_favorite_drinks(){
      var result = this.favorite_drinks
      if (this.genre_id != 0){
        result = result.filter((drink) => drink.DrinkGenreID == this.genre_id)
      }
      if (this.quantity == 0){
        return result
      }
      else if (this.quantity == 1){
        return result.filter((drink) => (drink.Quantity >0 && drink.Quantity <= 200))
      }
      else if (this.quantity == 2){
        return result.filter((drink) => (drink.Quantity >200 && drink.Quantity <=700))
      }
      else if (this.quantity == 3){
        return result.filter((drink) => (drink.Quantity >700 && drink.Quantity <=1000))
      }
      else if (this.quantity == 4){
        return result.filter((drink) => drink.Quantity >1000)
      }
      else{
        return result
      }
    },
    keyworddrinks() {
      var drinks = []; 
      if (this.keyword != null && this.keyword != ""){
        for(var i in this.drinks) {
            var drink = this.drinks[i];
            if(drink.name.indexOf(this.keyword) !== -1) {
                drinks.push(drink);
            }
          }
      }
      return drinks;
    },
    getItems: function() {
      let current = this.currentPage * this.parPage;
      let start = current - this.parPage;
      return this.serched_drinks.slice(start, current);
    },
    getPageCount: function() {
      return Math.ceil(this.serched_drinks.length / this.parPage);
    },
  },

  watch: {
    favorite_drinks_id:{
      immediate: true,
      // deep: true,
      handler(favorite){
        this.ReloadFavoriteDrink(favorite)
      }
    },
    drinks:{
      // deep: true,
      immediate: true,
      handler(){
        this.ReloadFavoriteDrink(this.favorite_drinks_id)
      }
    },
  },
  methods: {
    ChangeGenreID(genre_id) {
      this.currentPage = 1
      this.genre_id = genre_id
    },
    ChangeQuantity(i) {
      this.currentPage = 1
      this.quantity = i
    },
    ReloadFavoriteDrink(favorite){
      if (this.drinks != null){        
        for (const i in this.drinks){
          this.$set(this.drinks[i], "favorite" , false)
        }
        let result = []
        for (const i in favorite){
          const drink = (drink) => drink.ID == favorite[i].DrinkID
          if (this.drinks.some(drink)){
            console.log( this.drinks.find(drink))
            let name = this.drinks.find(drink).name
            let quantity = this.drinks.find(drink).Quantity
            let image = this.drinks.find(drink).Image
            result.push({ID: favorite[i].DrinkID, name: name, Image: image, Quantity: quantity, favorite: true})
            this.$set(this.drinks.find(drink), "favorite" , true)
          }
        }
        this.favorite_drinks = result
      }
    },
    doAddFavoriteDrink(item){
      // item.favorite = true
      this.$store.dispatch('login/doAddFavoriteDrink',item)
      if (this.$store.state.login.id != 0){
        // サーバへ送信するパラメータ
        const params = new URLSearchParams()
        params.append('drink_id', item.ID)
        params.append('user_id', this.$store.state.login.id)
        axios.post('https://kakaku-real-store.tk:8082/favorite_drink/add', params)
        .catch(error => {
          // handle error
          console.log(error)
        })
      }

    },
    doDeleteFavoriteDrink(item){
      // item.favorite = false
      this.$store.dispatch('login/doDeleteFavoriteDrink',item)
      if (this.$store.state.login.id != 0){
        const params = new URLSearchParams()
        params.append('drink_id', item.ID)
        params.append('user_id', this.$store.state.login.id)
        axios.post('https://kakaku-real-store.tk:8082/favorite_drink/delete', params)
        .catch(error => {
          // handle error
          console.log(error)
        })
      }
    },
    clickCallback(pageNum) {
       this.currentPage = Number(pageNum);
    },
    noImage(element){
      element.target.src = 'https://placehold.jp/600x300.png'
    },
  },
}
</script>
