<template>
  <div class="home">
    <b-container>
      <b-row class="my-1">
        <b-col sm="3">
          <label for="product_name">商品名:</label>
        </b-col>
        <b-col sm="9">
          <b-form-input id="product_name" type="text" v-model="product_name"></b-form-input>
        </b-col>
      </b-row>
      <b-row>
        <b-button pill v-on:click="doAddProduct">送信</b-button>
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
      products: []
    }
  },

  methods: {
    // 商品情報を登録する
    doAddProduct() {
      // サーバへ送信するパラメータ
      const params = new URLSearchParams();
      params.append('product_name', this.product_name)
      axios.post('http://localhost:8082/product/add', params)
      .then(response => {
          if (response.status != 200) {
              throw new Error('レスポンスエラー')
          } else {
            //一覧ページに遷移する
          }
      })
    },
  }
}
</script>
