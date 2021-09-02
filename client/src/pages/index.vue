<template>
  <div class="home">
    <b-container>
      <b-row>
        <router-link to="/drink">飲み物</router-link>
      </b-row>
    </b-container>
  </div>
</template>

<script>
// @ is an alias to /src
export default {
  name: 'home',
  // インスタンス作成時の処理
  created: function() {
    this.$router.push('/drink')
  },
}
</script>
