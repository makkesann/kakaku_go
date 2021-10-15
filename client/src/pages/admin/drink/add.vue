<template>
  <div class="login">
    <b-alert v-if="error" variant="danger" show>
      {{error}}
    </b-alert>
    <b-container>
      <div class="bellow-error form">
        <validation-observer ref="observer" v-slot="{handleSubmit}">
          <b-form @submit.stop.prevent="handleSubmit(doAddDrink)">
            <validation-provider
              name="商品名"
              :rules="{ required: true, min: 2 }"
              v-slot="validationContext"
            >
              <b-form-group id="productname-box">
                <b-row>
                  <b-col cols="3" class="text-right">
                    <label for="productname">商品名：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-input
                      id="productname"
                      name="productname"
                      v-model="productname"
                      :state="getValidationState(validationContext)"
                      aria-describedby="productname-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="productname-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
                  </b-col>
                </b-row>
              </b-form-group>
            </validation-provider>
            <validation-provider name="ジャンル" v-slot="validationContext">
              <b-form-group id="genre_id-box">
                <b-row>
                  <b-col cols="3" class="text-right">
                    <label for="genre_id">ジャンル：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-select
                      id="genre_id"
                      name="genre_id"
                      v-model="genre_id"
                      :options="genres"
                      text-field="name"
                      value-field="ID"
                      :state="getValidationState(validationContext)"
                      aria-describedby="genre_id-live-feedback"
                    ></b-form-select>
                    <b-form-invalid-feedback id="genre_id-live-feedback">
                      {{ validationContext.errors[0] }}
                    </b-form-invalid-feedback>
                  </b-col>
                </b-row>
              </b-form-group>
            </validation-provider>
            <validation-provider
              name="Janコード"
              :rules="{ numeric, length: 13 }"
              v-slot="validationContext"
            >
              <b-form-group id="Jancode-box">
                <b-row>
                  <b-col cols="3" class="text-right">
                    <label for="Jancode">Janコード：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-input
                      id="Jancode"
                      name="Jancode"
                      v-model="Jancode"
                      :state="getValidationState(validationContext)"
                      aria-describedby="Jancode-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="Jancode-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
                  </b-col>
                </b-row>
              </b-form-group>
            </validation-provider>
            <validation-provider
              name="画像ファイル名"
              v-slot="validationContext"
            >
              <b-form-group id="img_file_name-box">
                <b-row>
                  <b-col cols="3" class="text-right">
                    <label for="img_file_name">画像ファイル名：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-input
                      id="img_file_name"
                      name="img_file_name"
                      v-model="img_file_name"
                      :state="getValidationState(validationContext)"
                      aria-describedby="img_file_name-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="img_file_name-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
                  </b-col>
                </b-row>
              </b-form-group>
            </validation-provider>
            <validation-provider
              name="内容量"
              :rules="{numeric}"
              v-slot="validationContext"
            >
              <b-form-group id="quantity-box">
                <b-row>
                  <b-col cols="3" class="text-right">
                    <label for="quantity">内容量（ml）：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-input
                      id="quantity"
                      name="quantity"
                      v-model="quantity"
                      :state="getValidationState(validationContext)"
                      aria-describedby="quantity-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="quantity-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
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
      productname: null,
      genre_id: null,
      genres: [],
      Jancode: null,
      img_file_name: null,
      quantity: null,
      error: null,
    }
  },
  created: function() {
    this.GetGenres()
  },

  methods: {
    GetGenres(){
      axios.get('https://kakaku-real-store.tk:8082/drink/genres')
      .then((response) => {
        this.genres = response.data
      })
      .catch(error => {
        // handle error
        this.error = error.response.data.Detail
      })
    },
    // 商品情報を登録する
    doAddDrink() {
      // サーバへ送信するパラメータ
      const params = new URLSearchParams()
      params.append('drink_name', this.productname)
      params.append('genre_id', this.genre_id)
      params.append('jan', this.Jancode)
      params.append('image', this.img_file_name)
      params.append('quantity', this.quantity)
      axios.post('https://kakaku-real-store.tk:8082/drink/add', params)
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
