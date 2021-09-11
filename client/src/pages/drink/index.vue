<template>
  <div class="home">
    <b-container>
      <b-button pill v-on:click="idplus">くさ</b-button>
      <b-row>
        <b-col cols="2">
          <b-table striped hover :items="drink_genres" :fields="genre_fields">
            <template v-slot:cell(name)="{item}">
              <div class ="hover" @click="ChangeGenreID(item.ID)">{{ item.name }}</div>
            </template>
          </b-table>
        </b-col>
        <b-col cols="10">
          <!-- <h2>お気に入りの商品</h2> -->
          <b-table striped hover :items="favorite_drinks">
          </b-table>
          <b-table striped hover :items="serched_drinks" :fields="drink_fields">
            <template v-slot:cell(name)="{item}">
              <router-link :to="{name:'drink-id',params:{id: item.ID}}">{{ item.name }}</router-link>
            </template>
          </b-table>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
// @ is an alias to /src
export default {
  name: 'home',
  components: {
  },
  data(){
    return {
      drink_fields: ["name", "ID", "DrinkGenreID"],
      // drinks: [],
      genre_fields: ["name"],
      genre_id: 0,
      // drink_genres: []
      favorite_drinks: [],
      favorite_drink_triger: 0,
      // serched_drinks: []
    }
  },
  computed: {
    drinks() {
      return this.$store.getters["drink/getDrinks"]
    },
    drink_genres() {
      return this.$store.getters["drink/getDrinkGenres"]
    },
    serched_drinks(){
      if (this.genre_id == 0){
        return this.drinks
      } else {
        return this.drinks.filter((drink) => drink.ID == this.genre_id)
      }
    },
    favorite_drinks_id() {
      // this.favorite_drink_triger=1;
      return this.$store.getters["login/getFavoriteDrink"]
    },

    favorite_shops_id() {
      return this.$store.getters["login/getFavoriteShop"]
    },
  //   favorite_drinks() {
  //     let result = []
  //     // this.favorite_drink_triger = 0
  //     const favorite = this.favorite_drinks_id
  //     for (const i in favorite){
  //       let name = this.drinks.filter((drink) => drink.ID == favorite[i].DrinkID)[0].name
  //       // console.log(name)
  //       // console.log(favorite[i].DrinkID)
  //       result.push({id: favorite[i].DrinkID, name: name})
  //       console.log(result)
  //     }
  //     return result
  //   }
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
    idplus() {
      this.$store.dispatch('login/idplus')
    },
    ChangeGenreID(genre_id) {
      this.genre_id = genre_id
    },
    ReloadFavoriteDrink(favorite){
      let result = []
      // const favorite = this.favorite_drinks_id
      console.log(favorite)
      for (const i in favorite){
        console.log("いえい")
        console.log(favorite[i].DrinkID)
        console.log(this.drinks)
        console.log(this.drinks.filter((drink) => drink.ID == favorite[i].DrinkID)[0])
        if (this.drinks.some((drink) => drink.ID == favorite[i].DrinkID)){
          let name = this.drinks.filter((drink) => drink.ID == favorite[i].DrinkID)[0].name
          // console.log(name)
          // console.log(favorite[i].DrinkID)
          result.push({id: favorite[i].DrinkID, name: name})
          console.log("くそわろた" + JSON.stringify(result))
        }

      }
      this.favorite_drinks = result
      console.log("ぶちのめす")
      console.log(result)
    }
  },
}
</script>
