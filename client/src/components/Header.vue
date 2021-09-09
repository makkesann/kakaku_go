<template>
  <div id="header">
    <h1>飲み物の価格</h1>
    <router-link to="/user/login" v-if="id == 0">ログイン</router-link>
    <router-link :to="{name:'user-id',params:{id: id}}" v-else>マイページ</router-link>
  </div>
</template>

<script>
export default {
  name: 'Header',
  data(){
  return {
    guest: true,

    }
  },
  computed: {
    id() {
      return this.$store.getters["login/getID"];
    },
  },
}
</script>