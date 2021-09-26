<template>
  <div class="drink">
    <b-container>
      
      <h2>お気に入りの商品</h2>
      <b-row>
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
            <p class="hover mb-2" @click="ChangeGenreID(0)">リセット</p>
            <p class="hover mb-2">~200ml</p>
            <p class="hover mb-2">200ml~700ml</p>
            <p class="hover mb-2">700ml~1L</p>
            <p class="hover mb-2">1L~</p>
          </div>
        </b-col>
        <b-col cols="10">
          <!-- <p v-if="admin">商品追加</p> -->
          <b-table v-if="serched_favorite_drinks.length != 0" striped hover :items="serched_favorite_drinks">
          </b-table>
          <b-table hover :items="serched_drinks" :fields="drink_fields">
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
      drink_fields: ["画像", "商品名", "内容量", "お気に入り"],
      genre_id: 0,
      favorite_drinks: [],
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
      if (this.genre_id == 0){
        return this.drinks
      } else {
        return this.drinks.filter((drink) => drink.DrinkGenreID == this.genre_id)
      }
    },
    favorite_drinks_id() {
      return this.$store.getters["login/getFavoriteDrink"]
    },
    serched_favorite_drinks(){
      if (this.genre_id == 0){
        return this.favorite_drinks
      } else {
        return this.favorite_drinks.filter((drink) => drink.DrinkGenreID == this.genre_id)
      }
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
    // login_id:{
    //   immediate: true,
    //   // deep: true,
    //   handler(new_id, old_id){
    //     this.LogOutFavoriteDrink(new_id, old_id)
    //   }
    // },
  },
  methods: {
    ChangeGenreID(genre_id) {
      this.genre_id = genre_id
    },
    ReloadFavoriteDrink(favorite){
      for (const i in this.drinks){
        this.$set(this.drinks[i], "favorite" , false)
      }
      let result = []
      for (const i in favorite){
        const drink = (drink) => drink.ID == favorite[i].DrinkID
        if (this.drinks.some(drink)){
          let name = this.drinks.find(drink).name
          result.push({id: favorite[i].DrinkID, name: name})
          this.$set(this.drinks.find(drink), "favorite" , true)
        }
      }
      this.favorite_drinks = result
    },
    doAddFavoriteDrink(item){
      // item.favorite = true
      this.$store.dispatch('login/doAddFavoriteDrink',item)
      if (this.$store.state.login.id != 0){
        // サーバへ送信するパラメータ
        const params = new URLSearchParams()
        params.append('drink_id', item.ID)
        params.append('user_id', this.$store.state.login.id)
        axios.post('http://54.65.204.164:8082/favorite_drink/add', params)
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
        axios.post('http://54.65.204.164:8082/favorite_drink/delete', params)
        .catch(error => {
          // handle error
          console.log(error)
        })
      }
    },
    noImage(element){
      element.target.src = 'https://placehold.jp/600x300.png'
    },
  },
}
</script>
