<template>
  <div class="login">
    <b-alert v-if="error" variant="danger" show>
      {{error}}
    </b-alert>
    <b-container>
      <div class="bellow-error form">
        <validation-observer ref="observer" v-slot="{handleSubmit}">
          <b-form @submit.stop.prevent="handleSubmit(doDeleteDrinkGenre)">
            <validation-provider name="ジャンル名" :rules="{ required: true }" v-slot="validationContext">
              <b-form-group id="genre_id-box" label="ジャンル名" label-for="genre_id">
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
                <b-form-invalid-feedback id="genre_id-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
              </b-form-group>
            </validation-provider>
            <b-button type="submit" variant="primary">削除</b-button>
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
      genre_id: null,
      genres: [],
      error: null,
    }
  },
  created: function() {
    this.GetDrinkGenres()
  },
  methods: {
    GetDrinkGenres(){
      axios.get('https://kakaku-real-store.tk:8082/drink/genres')
      .then((response) => {
        this.genres = response.data
      })
      .catch(error => {
        // handle error
        this.error = error.response.data.Detail
      })
    },
    // 商品情報を削除する
    doDeleteDrinkGenre() {
      // サーバへ送信するパラメータ
      const genre_id = this.genre_id
      axios.post('https://kakaku-real-store.tk:8082/drink/genre/' + genre_id + '/delete')
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
