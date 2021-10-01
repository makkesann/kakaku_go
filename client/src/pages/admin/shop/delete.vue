<template>
  <div class="login">
    <b-alert v-if="error" variant="danger" show>
      {{error}}
    </b-alert>
    <b-container>
      <div class="bellow-error">
        <validation-observer ref="observer" v-slot="{handleSubmit}">
          <b-form @submit.stop.prevent="handleSubmit(doDeleteShop)">
            <validation-provider name="店舗名" :rules="{ required: true }" v-slot="validationContext">
              <b-form-group id="shop_id" label="店舗名" label-for="shop_id">
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
                <b-form-invalid-feedback id="shop_id-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
              </b-form-group>
            </validation-provider>
            <b-button type="submit" variant="danger">削除</b-button>
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
      shop_id: null,
      shops: [],
      error: null,
    }
  },
  created: function() {
    this.GetShops()
  },
  methods: {
    GetShops(){
      axios.get('http://localhost:8082/shops')
      .then((response) => {
        this.shops = response.data
      })
      .catch(error => {
        // handle error
        this.error = error.response.data.Detail
      })
    },
    // 商品情報を削除する
    doDeleteShop() {
      // サーバへ送信するパラメータ
      axios.post('http://localhost:8082/shop/' + this.shop_id + '/delete')
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
