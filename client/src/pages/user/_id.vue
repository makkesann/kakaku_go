<template>
  <div class="home">
    <b-container>
      <b-row>
        <b-table striped hover :items="favorite_drinks"></b-table>
      </b-row>
      <b-row>
        <b-table striped hover :items="favorite_shops"></b-table>
      </b-row>
    </b-container>
  </div>
</template>
<script>
// @ is an alias to /src
export default {
  name: 'home',
  beforeCreate() {
    if(this.$store.state.login.id != this.$route.params.id){
      this.$router.push('/drink')
    }
  },
  computed: {
    favorite_drinks() {
      return this.$store.getters["login/getFavoriteDrink"]
    },
    favorite_shops() {
      return this.$store.getters["login/getFavoriteShop"]
    }
  },
}
</script>
