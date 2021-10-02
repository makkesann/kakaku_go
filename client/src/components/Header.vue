<template>
  <div id="header">
  <div id="header-box">
    <b-container class="pc">
      <b-row>
        <b-navbar class="w-100">

          <b-navbar-nav class="mr-auto">

              <b-nav-item to="/drink" id="nav-home">Home</b-nav-item>
              <b-nav-item class="d-flex align-items-center" @click="doAdminLogIn(false)" v-if="this.$store.state.login.admin">管理者ログアウト</b-nav-item>
              <b-nav-item class="d-flex align-items-center"  to="/admin" v-if="this.$store.state.login.admin">管理者ページ</b-nav-item>
              <b-nav-item class="d-flex align-items-center" @click="doAdminLogIn(true)" v-if="!this.$store.state.login.admin">管理者ログイン</b-nav-item>
          </b-navbar-nav>
          <b-navbar-nav>
              <b-nav-item class="d-flex align-items-center" v-if="id == 0" to="/user/login">ログイン</b-nav-item>
              <b-nav-item class="d-flex align-items-center" v-if="id == 0" to="/user/new">アカウント作成</b-nav-item>
              <b-nav-item class="d-flex align-items-center" v-if="id != 0" :to="{name:'user-id',params:{id: id}}">マイページ</b-nav-item>
              <b-nav-item class="d-flex align-items-center" v-if="id != 0" @click="doLogOut">ログアウト</b-nav-item>
          </b-navbar-nav>
        </b-navbar>
      </b-row>
    </b-container>
    <b-container class="sp">
      <b-row>
        <b-navbar type="dark" toggleable class="w-100">

          <b-nav-item to="/drink" id="nav-home">Home</b-nav-item>
          <b-navbar-toggle target="navbar-toggle-collapse">
            <template #default="{ expanded }">
              <b-icon v-if="expanded" icon="chevron-bar-up"></b-icon>
              <b-icon v-else icon="chevron-bar-down"></b-icon>
            </template>
          </b-navbar-toggle>
          <b-collapse id="navbar-toggle-collapse" is-nav>
            <b-nav-item class="d-flex align-items-center" @click="doAdminLogIn(false)" v-if="this.$store.state.login.admin">管理者ログアウト</b-nav-item>
            <b-nav-item class="d-flex align-items-center"  to="/admin" v-if="this.$store.state.login.admin">管理者ページ</b-nav-item>
            <b-nav-item class="d-flex align-items-center" @click="doAdminLogIn(true)" v-if="!this.$store.state.login.admin">管理者ログイン</b-nav-item>
            <b-nav-item class="d-flex align-items-center" v-if="id == 0" to="/user/login">ログイン</b-nav-item>
            <b-nav-item class="d-flex align-items-center" v-if="id == 0" to="/user/new">アカウント作成</b-nav-item>
            <b-nav-item class="d-flex align-items-center" v-if="id != 0" :to="{name:'user-id',params:{id: id}}">マイページ</b-nav-item>
            <b-nav-item class="d-flex align-items-center" v-if="id != 0" @click="doLogOut">ログアウト</b-nav-item>
          </b-collapse>
        </b-navbar>
      </b-row>
    </b-container>
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