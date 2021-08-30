<template>
  <div class="home">
    <b-container>
      <b-row class="my-1" v-for="product in products" :key="product">
        <b-col sm="3">
          <label :for="`type-${type}`">商品名 <code>{{ type }}</code>:</label>
        </b-col>
        <b-col sm="9">
          <b-form-input :id="`type-${type}`" :type="type"></b-form-input>
        </b-col>
      </b-row>
      <b-row>
        <b-button pill disabled="!isEntered" click="doAddProduct">送信</b-button>
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

  // インスタンス作成時の処理
  created: function() {
      this.doFetchAllProducts()
  },

  methods: {
      // 全ての商品情報を取得する
      doFetchAllProducts() {
          axios.get('http://localhost:8082/products')
          .then(response => {
              if (response.status != 200) {
                  throw new Error('レスポンスエラー')
              } else {
                  var resultProducts = response.data

                  // サーバから取得した商品情報をdataに設定する
                  this.products = resultProducts
              }
          })
      },

      // 商品情報を登録する
      doAddProduct() {
          // サーバへ送信するパラメータ
          const params = new URLSearchParams();
          params.append('productName', this.productName)

          axios.post('http://localhost:8082/add', params)
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
