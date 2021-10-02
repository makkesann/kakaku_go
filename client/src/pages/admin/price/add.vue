<template>
  <div class="login">
    <b-alert v-if="error" variant="danger" show>
      {{error}}
    </b-alert>
    <b-container>
      <div class="bellow-error">
        <validation-observer ref="observer" v-slot="{handleSubmit}">
          <b-form @submit.stop.prevent="handleSubmit(doAddPrice)">
            <validation-provider name="商品名" v-slot="validationContext" :rules="{ required: true}">
              <b-form-group id="drink_id">
                <b-row>
                  <b-col cols="3" class="text-right">
                    <label for="drink_id">商品名：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-select
                      id="drink_id"
                      name="drink_id"
                      v-model="drink_id"
                      :options="drinks"
                      text-field="name"
                      value-field="ID"
                      :state="getValidationState(validationContext)"
                      aria-describedby="drink_id-live-feedback"
                    ></b-form-select>
                    <b-form-invalid-feedback id="drink_id-live-feedback">
                      {{ validationContext.errors[0] }}
                    </b-form-invalid-feedback>
                  </b-col>
                </b-row>
              </b-form-group>
            </validation-provider>
            <validation-provider name="店名" v-slot="validationContext" :rules="{ required: true}">
              <b-form-group id="shop_id">
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
              <b-form-group id="price">
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
      drinks:[],
      drink_id:null,
      shop_id:null,
      price:null,
    }
  },
  created: function() {
    this.GetDrinks()
    this.GetShops()
  },

  methods: {
    GetDrinks(){
      axios.get('http://kakaku-real-store.tk:8082/drinks')
      .then((response) => {
        this.drinks = response.data
      })
      .catch(error => {
        // handle error
        this.error = error.response.data.Detail
      })
    },
    GetShops(){
      axios.get('http://kakaku-real-store.tk:8082/shops')
      .then((response) => {
        this.shops = response.data
      })
      .catch(error => {
        // handle error
        this.error = error.response.data.Detail
      })
    },
    // 価格情報を登録する
    doAddPrice() {
      // サーバへ送信するパラメータ
      const params = new URLSearchParams()
      params.append('price', this.price)
      params.append('drink_id', this.drink_id)
      params.append('shop_id', this.shop_id)

      axios.post('http://kakaku-real-store.tk:8082/price/add', params)
      .then(() => {
        this.$router.push('/drink')
      })
      .catch(error => {
        // handle error
        this.error = error.response.data.Detail
      })
    },
    getValidationState({ dirty, validated, valid = null }) {
      return dirty || validated ? valid : null;
    },
  }
}
</script>
