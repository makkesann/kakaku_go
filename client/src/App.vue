<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>

export default {
  name: 'App',
  beforeCreate(){
    this.$store.dispatch('drink/doFetchAllDrink')
    this.$store.dispatch('drink/doFetchAllDrinkGenre')
    if (this.login_id != 0){
      this.$store.dispatch('login/doFetchFavoriteDrinks',this.$store.state.login.id)
      this.$store.dispatch('login/doFetchFavoriteShops',this.$store.state.login.id)
    }
  },
  computed: {
    login_id() {
      return this.$store.getters["login/getID"]
    },

  },
  watch: {
    login_id(new_id, old_id) {
      this.$store.dispatch('login/doLogout',{ old_id:old_id, new_id:new_id })
      this.$store.dispatch('login/doFetchFavoriteDrinks',new_id)
      this.$store.dispatch('login/doFetchFavoriteShops',new_id)
    }
  }
}
</script>

<style>
#app {
  /* font-family: Avenir, Helvetica, Arial, sans-serif; */
  /* font-family: 'Hiragino Kaku Gothic ProN','Hiragino Sans',Meiryo,sans-serif; */
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}
</style>
