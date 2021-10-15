<template>
  <div class="edit">
    <b-alert v-if="error" variant="danger" show>
      {{error}}
    </b-alert>
    <b-container>
      <div class="bellow-error form">
        <validation-observer ref="observer" v-slot="{handleSubmit}">
          <b-form @submit.stop.prevent="handleSubmit(doChangeDrinkGenreName)">
            <validation-provider name="変更するジャンル" v-slot="validationContext">
              <b-form-group id="genre_id-box">
                <b-row>
                  <b-col cols="3" class="text-right">
                    <label for="genre_id">変更するジャンル：</label>
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
              name="変更後のジャンル名"
              :rules="{ required: true }"
              v-slot="validationContext"
            >
              <b-form-group id="genre_name-box">
                <b-row>
                  <b-col cols="3" class="text-right">
                    <label for="genre_name">変更後のジャンル名：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-input
                      id="genre_name"
                      name="genre_name"
                      v-model="genre_name"
                      :state="getValidationState(validationContext)"
                      aria-describedby="genre_name-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="genre_name-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
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
      genre_name: null,
      genre_id: null,
      genres: [],
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
    doChangeDrinkGenreName() {
      if (this.genre_id!=null){

        // サーバへ送信するパラメータ
        const params = new URLSearchParams()
        params.append('name', this.genre_name)
        const genre_id = this.genre_id
        axios.post('https://kakaku-real-store.tk:8082/drink/genre/name/' + genre_id + '/change', params)
        .then(() => {
          this.$router.push('/drink')
        })
        .catch(error => {
          // handle error
          this.error = error.response.data.Detail
        })
      } else{
        this.error = ("変更するジャンル名を選択してください")
      }
    },
    getValidationState({ dirty, validated, valid = null }) {
      return dirty || validated ? valid : null;
    },
  }
}
</script>
