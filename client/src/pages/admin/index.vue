<template>
  <div class="admin-home">
    <b-container>
      <h3 class="text-left">商品</h3>
      <b-row class="text-left ad-home-row">
        <router-link to="/admin/drink/add">商品追加</router-link>
        <router-link to="/admin/drink/edit">商品編集</router-link>
        <router-link to="/admin/drink/delete">商品削除</router-link><br>
      </b-row>
      <h3 class="text-left">ジャンル</h3>
      <b-row class="text-left ad-home-row">
        <router-link to="/admin/drink_genre/add">ジャンル追加</router-link><br>
        <router-link to="/admin/drink_genre/delete">ジャンル削除</router-link><br>
      </b-row>
      <h3 class="text-left">ユーザー</h3>
      <b-row class="text-left ad-home-row">

        <router-link to="/admin/user/delete">ユーザー削除</router-link><br>
      </b-row>
    </b-container>
  </div>
</template>
<script>
export default {
  name: 'home',
  beforeCreate() {
    if(this.$store.state.login.admin == false){
      this.$router.push('/drink')
    }
  },
}
</script>