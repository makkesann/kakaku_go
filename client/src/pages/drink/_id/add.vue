<template>
  <div class="home">
    <b-container>
      <h2>お店</h2>{{ shop_input }}
      <b-form-select v-model="shop_input" :options="shops" value-field="ID" text-field="Name"></b-form-select>
      <h2>価格</h2>
      <b-form-input v-model="price_input" placeholder="価格を入力してください" type="number"></b-form-input>
      <b-row>
        <b-button pill v-on:click="doAddPrice">送信</b-button>
      </b-row>
    </b-container>
  </div>
</template>

<script>
// @ is an alias to /src
import axios from 'axios'
export default {
  name: 'home',
  data: function(){
    return {
      shops: [],
      price_input: 0,
      shop_input: 0,
    }
  },

  // インスタンス作成時の処理
  created: function() {
    this.GetShops()
  },

  methods: {
    GetShops(){
      axios.get('http://localhost:8082/shops')
      .then(response => {
        this.shops = response.data
      })
      .catch(error => {
        // handle error
        console.log(error)
      })
    },
    doAddPrice() {
      // サーバへ送信するパラメータ
      let drink_id = this.$route.params.id
      const params = new URLSearchParams()
      console.log(drink_id)
      params.append('drink_id', drink_id)
      params.append('price', this.price_input)
      params.append('shop_id', this.shop_input)
      // console.log(params)
      axios.post('http://localhost:8082/price/add', params)
      .then(response => {
          if (response.status != 200) {
              throw new Error('レスポンスエラー')
          } else {
            //一覧ページに遷移する
            this.$router.push('/drink')
          }
      })
    },
  },
}
</script>
