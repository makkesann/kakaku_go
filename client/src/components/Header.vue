<template>
  <div id="header">
    <h1>飲み物の価格</h1>
    <router-link to="/user/login" v-if="id == 0">ログイン</router-link>
    <div v-else>
      <router-link :to="{name:'user-id',params:{id: id}}">マイページ</router-link>
      <div class="hover" @click="doLogOut">ログアウト</div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Header',
  data(){
  return {
    }
  },
  computed: {
    id() {
      return this.$store.getters["login/getID"]
    },
  },
  methods: {
    doLogOut(){
      this.$store.dispatch('login/doSetID',0)
    }
  }
}
</script>