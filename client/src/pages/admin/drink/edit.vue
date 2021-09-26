<template>
  <div class="login">
    <b-alert v-if="error" variant="danger" show>
      {{error}}
    </b-alert>
    <b-container>
      <div>
        <validation-observer ref="observer" v-slot="handleSubmit">
          <b-form @submit.stop.prevent="handleSubmit(onSubmit)">
            <validation-provider name="変更する商品" v-slot="validationContext">
              <b-form-group id="drink_id">
                <b-row>
                  <b-col cols="3" class="text-right">
                    <label for="drink_id">変更する商品：</label>
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
            <validation-provider
              name="商品名"
              :rules="{ required: true, min: 3 }"
              v-slot="validationContext"
            >
              <b-form-group id="productname">
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
              <b-form-group id="genre_id">
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
              :rules="{ numeric, length: 3, length: 13 }"
              v-slot="validationContext"
            >
              <b-form-group id="Jancode">
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
              <b-form-group id="img_file_name">
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
            <b-button type="submit" variant="warning" @click="doDeleteDrink" :disabled="handleSubmit.invalid || !handleSubmit.validated">変更</b-button>
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
      drink_id: null,
      drinks: [],
      error: null,
      productname: null,
      genre_id: null,
      genres: [],
      Jancode: null,
      img_file_name: null,
    }
  },
  created: function() {
    this.GetDrinks()
    this.GetGenres()
  },

  methods: {
    GetDrinks(){
      axios.get('http://localhost:8082/drinks')
      .then(response => {
        this.drinks = response.data
      })
      .catch(error => {
        // handle error
        console.log(error)
      })
    },
    GetGenres(){
      axios.get('http://localhost:8082/drink/genres')
      .then(response => {
        this.genres = response.data
      })
      .catch(error => {
        // handle error
        console.log(error)
      })
    },
    // 商品情報を削除する
    doAddDrink() {
      // サーバへ送信するパラメータ
      const params = new URLSearchParams()
      params.append('drink_name', this.productname)
      params.append('genre_id', this.genre_id)
      params.append('jan', this.Jancode)
      params.append('image', this.img_file_name)
      axios.post('http://localhost:8082/drink/' + this.drink_id + 'change')
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
