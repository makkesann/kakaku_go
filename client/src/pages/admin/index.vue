<template>
  <div class="admin-home">
    <b-container>
      <h3 class="text-left">商品</h3>
      <b-row class="text-left ad-home-row">
        <router-link to="/admin/drink/add">商品追加</router-link>
        <router-link to="/admin/drink/edit">商品編集</router-link>
        <router-link to="/admin/drink/delete">商品削除</router-link>
      </b-row>
      <h3 class="text-left">ジャンル</h3>
      <b-row class="text-left ad-home-row">
        <router-link to="/admin/drink_genre/add">ジャンル追加</router-link>
        <router-link to="/admin/drink_genre/edit">ジャンル編集</router-link>
        <router-link to="/admin/drink_genre/delete">ジャンル削除</router-link>
      </b-row>
      <h3 class="text-left">店舗</h3>
      <b-row class="text-left ad-home-row">
        <router-link to="/admin/shop/add">店舗追加</router-link>
        <router-link to="/admin/shop/edit">店舗編集</router-link>
        <router-link to="/admin/shop/delete">店舗削除</router-link>
      </b-row>
      <h3 class="text-left">価格</h3>
      <b-row class="text-left ad-home-row">
        <router-link to="/admin/price/add">価格追加</router-link>
        <router-link to="/admin/price/edit">価格編集</router-link>
        <router-link to="/admin/price/delete">価格削除</router-link>
      </b-row>
      <h3 class="text-left">ユーザー</h3>
      <b-row class="text-left ad-home-row">

        <router-link to="/admin/user/delete">ユーザー削除</router-link>
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