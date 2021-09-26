<template>
  <div class="home">
    <h2>管理者ページだよ</h2>
    <router-link to="/admin/drink/add">商品追加</router-link><br>
    <router-link to="/admin/drink/edit">商品編集</router-link><br>
    <router-link to="/admin/drink/delete">商品削除</router-link><br>
    <router-link to="/admin/drink_genre/add">ジャンル追加</router-link><br>
    <router-link to="/admin/user/delete">ユーザー削除</router-link><br>
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