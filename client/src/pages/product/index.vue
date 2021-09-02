<template>
  <div class="home">
    <b-container>
      <b-row>
        <b-table striped hover :items="products"></b-table>
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
      // １つの商品情報を取得する
      doFetchProduct(product) {
          axios.get('/fetchProduct', {
              params: {
                  productID: product.id
              }
          })
          .then(response => {
              if (response.status != 200) {
                  throw new Error('レスポンスエラー')
              } else {
                  var resultProduct = response.data

                  // 選択された商品情報のインデックスを取得する
                  var index = this.products.indexOf(product)

                  // spliceを使うとdataプロパティの配列の要素をリアクティブに変更できる
                  this.products.splice(index, 1, resultProduct[0])
              }
          })
      },
      // 商品情報を登録する
      doAddProduct() {
          // サーバへ送信するパラメータ
          const params = new URLSearchParams();
          params.append('productName', this.productName)
          params.append('productMemo', this.productMemo)

          axios.post('/addProduct', params)
          .then(response => {
              if (response.status != 200) {
                  throw new Error('レスポンスエラー')
              } else {
                  // 商品情報を取得する
                  this.doFetchAllProducts()

                  // 入力値を初期化する
                  this.initInputValue()
              }
          })
      },
      // 商品情報の状態を変更する
      doChangeProductstate(product) {
          // サーバへ送信するパラメータ
          const params = new URLSearchParams();
          params.append('productID', product.id)
          params.append('Productstate', product.state)

          axios.post('/changeStateProduct', params)
          .then(response => {
              if (response.status != 200) {
                  throw new Error('レスポンスエラー')
              } else {
                  // 商品情報を取得する
                  this.doFetchProduct(product)
              }
          })
      },
      // 商品情報を削除する
      doDeleteProduct(product) {
          // サーバへ送信するパラメータ
          const params = new URLSearchParams();
          params.append('productID', product.id)

          axios.post('/deleteProduct', params)
          .then(response => {
              if (response.status != 200) {
                  throw new Error('レスポンスエラー')
              } else {
                  // 商品情報を取得する
                  this.doFetchAllProducts()
              }
          })
      },
      // 入力値を初期化する
      initInputValue() {
          this.current = -1
          this.productName = ''
          this.productMemo = ''
      }
  }
}
</script>
