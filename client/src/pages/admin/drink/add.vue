<template>
  <div class="login">
    <b-alert v-if="error" variant="danger" show>
      {{error}}
    </b-alert>
    <b-container>
      <div>
        <validation-observer ref="observer" v-slot="handleSubmit">
          <b-form @submit.stop.prevent="handleSubmit(onSubmit)">
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
            <b-button type="submit" variant="primary" @click="doAddDrink" :disabled="handleSubmit.invalid || !handleSubmit.validated">追加</b-button>
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
      error: null
    }
  },
  created: function() {
    this.GetGenres()
  },

  methods: {
    GetGenres(){
      axios.get('http://54.65.204.164:8082/drink/genres')
      .then(response => {
        this.genres = response.data
      })
      .catch(error => {
        // handle error
        console.log(error)
      })
    },
    // 商品情報を登録する
    doAddDrink() {
      // サーバへ送信するパラメータ
      const params = new URLSearchParams()
      params.append('drink_name', this.productname)
      params.append('genre_id', this.genre_id)
      axios.post('http://54.65.204.164:8082/drink/add', params)
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
