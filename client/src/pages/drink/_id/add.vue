<template>
  <div class="login">
    <b-alert v-if="error" variant="danger" show>
      {{error}}
    </b-alert>
    <b-container>
      <div class="bellow-error form">
        <validation-observer ref="observer" v-slot="{handleSubmit}">
          <b-form @submit.stop.prevent="handleSubmit(doAddPrice)">
            <validation-provider name="店名" v-slot="validationContext" :rules="{ required: true}">
              <b-form-group id="shop_id-box">
                <b-row>
                  <b-col cols="3" class="text-right">
                    <label for="shop_id">店名：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-select
                      id="shop_id"
                      name="shop_id"
                      v-model="shop_id"
                      :options="shops"
                      text-field="Name"
                      value-field="ID"
                      :state="getValidationState(validationContext)"
                      aria-describedby="shop_id-live-feedback"
                    ></b-form-select>
                    <b-form-invalid-feedback id="shop_id-live-feedback">
                      {{ validationContext.errors[0] }}
                    </b-form-invalid-feedback>
                  </b-col>
                </b-row>
              </b-form-group>
            </validation-provider>
            <validation-provider
              name="価格"
              :rules="{ required, numeric }"
              v-slot="validationContext"
            >
              <b-form-group id="price-box">
                <b-row>
                  <b-col cols="3" class="text-right">
                    <label for="price">価格：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-input
                      id="price"
                      name="price"
                      v-model="price"
                      :state="getValidationState(validationContext)"
                      aria-describedby="price-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="price-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
                  </b-col>
                </b-row>
              </b-form-group>
            </validation-provider>
            <b-button type="submit" variant="primary">追加</b-button>
          </b-form>
        </validation-observer>
      </div>
    </b-container>
  </div>
</template>
<script>
// @ is an alias to /src
import axios from 'axios'
export default {
  name: 'home',
  data() {
    return {
      error: null,
      shops:[],
      shop_id:null,
      price:null,
    }
  },
  created: function() {
    this.GetShops()
  },

  methods: {
    GetShops(){
      axios.get('https://kakaku-real-store.tk:8082/shops')
      .then(response => {
        this.shops = response.data
      })
      .catch(error => {
        // handle error
        this.error = error.response.data.Detail
      })
    },
    doAddPrice() {
      // サーバへ送信するパラメータ
      let drink_id = this.$route.params.id
      const params = new URLSearchParams()
      params.append('drink_id', drink_id)
      params.append('price', this.price)
      params.append('shop_id', this.shop_id)
      axios.post('https://kakaku-real-store.tk:8082/price/add', params)
      .then(response => {
          if (response.status != 200) {
              throw new Error('レスポンスエラー')
          } else {
            //一覧ページに遷移する
            this.$router.push('/drink')
          }
      })
    },
    getValidationState({ dirty, validated, valid = null }) {
      return dirty || validated ? valid : null;
    },
  }
}
</script>
