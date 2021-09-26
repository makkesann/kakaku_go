<template>
  <div id="header">
    <h1>飲み物の価格</h1>
    <div v-if="this.$store.state.login.admin">
      <p @click="doAdminLogIn(false)">管理者ログアウト</p>
      <router-link to="/admin">管理者ページ</router-link>
    </div>
    <div v-else>
      <p @click="doAdminLogIn(true)">管理者ログイン</p>
    </div>
    <router-link class="mr-2" to="/user/login" v-if="id == 0">ログイン</router-link>
    <router-link to="/user/new" v-if="id == 0">アカウント作成</router-link>
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
      this.$store.dispatch('login/doLogout', {old_id:this.id, new_id:0 })
    },
    doAdminLogIn(ad_state){
      this.$store.dispatch('login/doAdminLogIn',ad_state)
    }
  }
}
</script>