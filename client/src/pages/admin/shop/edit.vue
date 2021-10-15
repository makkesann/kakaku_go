<template>
  <div class="edit">
    <b-alert v-if="error" variant="danger" show>
      {{error}}
    </b-alert>
    <b-container>
      <div class="bellow-error form">
        <validation-observer ref="observer" v-slot="{handleSubmit}">
          <b-form @submit.stop.prevent="handleSubmit(doChangeShopName)">
            <validation-provider name="変更する店舗名" v-slot="validationContext" :rules="{ required: true }">
              <b-form-group id="shop_id-box">
                <b-row>
                  <b-col cols="3" class="text-right text-left-sp">
                    <label for="shop_id">変更する店舗名：</label>
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
              name="変更後の店舗名"
              :rules="{ required: true }"
              v-slot="validationContext"
            >
              <b-form-group id="shop_name-box">
                <b-row>
                  <b-col cols="3" class="text-right text-left-sp">
                    <label for="shop_name">変更後の店舗名：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-input
                      id="shop_name"
                      name="shop_name"
                      v-model="shop_name"
                      :state="getValidationState(validationContext)"
                      aria-describedby="shop_name-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="shop_name-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
                  </b-col>
                </b-row>
              </b-form-group>
            </validation-provider>
            <b-button type="submit" variant="warning">変更</b-button>
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
      shop_name: null,
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
      axios.get('http://kakaku-real-store.tk:8082/shops')
      .then((response) => {
        this.shops = response.data
      })
      .catch(error => {
        // handle error
        this.error = error.response.data.Detail
      })
    },
    doChangeShopName() {
      if (this.shop_id!=null){
        // サーバへ送信するパラメータ
        const params = new URLSearchParams()
        params.append('name', this.shop_name)
        axios.post('http://kakaku-real-store.tk:8082/shop/name/' + this.shop_id + '/change', params)
        .then(() => {
          this.$router.push('/drink')
        })
        .catch(error => {
          // handle error
          this.error = error.response.data.Detail
        })
      } else{
        this.error = ("変更する店舗名を選択してください")
      }
    },
    getValidationState({ dirty, validated, valid = null }) {
      return dirty || validated ? valid : null;
    },
  }
}
</script>
