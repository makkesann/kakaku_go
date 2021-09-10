<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>

export default {
  name: 'App',
  created(){
    this.$store.dispatch('drink/doFetchAllDrink')
    this.$store.dispatch('drink/doFetchAllDrinkGenre')
    if (this.login_id != 0){
      this.$store.dispatch('login/doFetchFavoriteDrinks',this.login_id)
      this.$store.dispatch('login/doFetchFavoriteShops',this.login_id)
    }
  },
  computed: {
    login_id() {
      return this.$store.getters["login/getID"]
    },

  },
  watch: {
    login_id(id) {
      this.$store.dispatch('login/doFetchFavoriteDrinks',id)
      this.$store.dispatch('login/doFetchFavoriteShops',id)
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
