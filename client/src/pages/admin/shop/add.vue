<template>
  <div class="login">
    <b-alert v-if="error" variant="danger" show>
      {{error}}
    </b-alert>
    <b-container>
      <div class="bellow-error">
        <validation-observer ref="observer" v-slot="{handleSubmit}">
          <b-form @submit.stop.prevent="handleSubmit(doAddDrinkShop)">
            <validation-provider
              name="店名"
              :rules="{ required: true }"
              v-slot="validationContext"
            >
              <b-form-group id="shopname">
                <b-row>
                  <b-col cols="3" class="text-right">
                    <label for="shopname">店名：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-input
                      id="shopname"
                      name="shopname"
                      v-model="shopname"
                      :state="getValidationState(validationContext)"
                      aria-describedby="shopname-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="shopname-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
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
      shopname: null,
      error: null,
    }
  },
  methods: {
    // 商品情報を登録する
    doAddDrinkGenre() {
      // サーバへ送信するパラメータ
      const params = new URLSearchParams()
      params.append('shop_name', this.shopname)
      axios.post('http://54.65.204.164:8082/drink/shop/add', params)
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
