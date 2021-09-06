<template>
  <div class="home">
    <b-container>
      <b-row>
        <b-col cols="2">
          <b-table striped hover :items="drink_genres" :fields="genre_fields"></b-table>
        </b-col>
        <b-col cols="10">
          <b-table striped hover :items="drinks" :fields="drink_fields">
            <template v-slot:cell(name)="{item}">
              <a :href="'drink/' + item.ID" >{{ item.name }}</a>
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
      // drink_genres: []
    }
  },


  computed: {
    drinks() {
      return this.$store.getters["drink/getDrinks"];
    },
    drink_genres() {
      return this.$store.getters["drink/getDrinkGenres"];
    }
  },
}
</script>
