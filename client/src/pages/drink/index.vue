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
      // serched_drinks: []
    }
  },
  computed: {
    drinks() {
      return this.$store.getters["drink/getDrinks"];
    },
    drink_genres() {
      return this.$store.getters["drink/getDrinkGenres"];
    },
    serched_drinks(){
      if (this.genre_id == 0){
        return this.drinks
      } else {
        return this.drinks.filter((drink) => drink.ID == this.genre_id)
      }
      
    }
  },
  methods: {
    idplus() {
      this.$store.dispatch('login/idplus')
    },
    ChangeGenreID(genre_id) {
      this.genre_id = genre_id
    }
  }
}
</script>
