<template>
  <div class="home">
    <b-container>
      <b-row class="my-1">
        <b-col sm="3">
          <label for="drink_name">商品名:</label>
        </b-col>
        <b-col sm="9">
          <b-form-input id="drink_name" type="text" v-model="drink_name"></b-form-input>
        </b-col>
      </b-row>
      <b-row>
        <b-button pill v-on:click="doAddDrink">送信</b-button>
      </b-row>
    </b-container>
  </div>
</template>

<script>
// @ is an alias to /src
import axios from 'axios'
export default {
  name: 'home',
  components: {
  },
  data: function(){
    return {
      drinks: []
    }
  },

  methods: {
    // 商品情報を登録する
    doAddDrink() {
      // サーバへ送信するパラメータ
      const params = new URLSearchParams()
      params.append('drink_name', this.drink_name)
      axios.post('http://54.65.204.164.164:8082/drink/add', params)
      .then(response => {
          if (response.status != 200) {
              throw new Error('レスポンスエラー')
          } else {
            //一覧ページに遷移する
            this.$router.push('/drink')
          }
      })
    },
  }
}
</script>
