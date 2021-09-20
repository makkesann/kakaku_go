<template>
  <div class="home">
    <b-container>
      <div class ="hover" @click="ChangeGenreID(0)">ジャンルリセット</div>
      <h2>お気に入りの商品</h2>
      <b-row>
        <b-col cols="2">
          <b-table striped hover :items="drink_genres" :fields="genre_fields">
            <template v-slot:cell(name)="{item}">
              <div class ="hover" @click="ChangeGenreID(item.ID)">{{ item.name }}</div>
            </template>
          </b-table>
        </b-col>
        <b-col cols="10">
          <p v-if="admin">商品追加</p>
          <b-table v-if="serched_favorite_drinks.length != 0" striped hover :items="serched_favorite_drinks">
          </b-table>
          <b-table striped hover :items="serched_drinks" :fields="drink_fields">
            <template v-slot:cell(name)="{item}">
              <router-link :to="{name:'drink-id',params:{id: item.ID}}">{{ item.name }}</router-link><p v-if="admin">削除</p>
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
      drink_fields: ["name", "ID", "DrinkGenreID", "お気に入り"],
      // drinks: [],
      genre_fields: ["name"],
      genre_id: 0,
      // drink_genres: []
      favorite_drinks: [],
      // serched_drinks: []
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
  },
  methods: {
    ChangeGenreID(genre_id) {
      this.genre_id = genre_id
    },
    ReloadFavoriteDrink(favorite){
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
        axios.post('http://localhost:8082/favorite_drink/add', params)
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
        axios.post('http://localhost:8082/favorite_drink/delete', params)
        .catch(error => {
          // handle error
          console.log(error)
        })
      }
    }
  },
}
</script>
