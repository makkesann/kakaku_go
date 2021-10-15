<template>
  <div class="login">
    <b-alert v-if="error" variant="danger" show>
      {{error}}
    </b-alert>
    <b-container>
      <div class="bellow-error form">
        <validation-observer ref="observer" v-slot="{handleSubmit}">
          <b-form @submit.stop.prevent="handleSubmit(doAddDrinkGenre)">
            <validation-provider
              name="ジャンル名"
              :rules="{ required: true }"
              v-slot="validationContext"
            >
              <b-form-group id="genrename-box">
                <b-row>
                  <b-col cols="3" class="text-right">
                    <label for="genrename">ジャンル名：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-input
                      id="genrename"
                      name="genrename"
                      v-model="genrename"
                      :state="getValidationState(validationContext)"
                      aria-describedby="genrename-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="genrename-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
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
      genrename: null,
      error: null,
    }
  },
  methods: {
    // 商品情報を登録する
    doAddDrinkGenre() {
      // サーバへ送信するパラメータ
      const params = new URLSearchParams()
      params.append('genre_name', this.genrename)
      axios.post('https://kakaku-real-store.tk:8082/drink/genre/add', params)
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
