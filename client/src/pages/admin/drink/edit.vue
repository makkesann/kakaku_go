<template>
  <div class="login">
    <b-alert v-if="error" variant="danger" show>
      {{error}}
    </b-alert>
    <b-container>
      <div class="bellow-error form">
        <b-row class="mb-5">
          <b-col cols="3" class="text-right text-left-sp">
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
            ></b-form-select>
          </b-col>
        </b-row>
        <validation-observer ref="observer" v-slot="{handleSubmit}">
          <b-form @submit.stop.prevent="handleSubmit(doChangeDrinkName)">
            <validation-provider
              name="商品名"
              :rules="{ required: true, min: 3 }"
              v-slot="validationContext"
            >
              <b-form-group id="productname-box">
                <b-row>
                  <b-col cols="3" class="text-right text-left-sp">
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
            <b-button type="submit" variant="warning">変更</b-button>
          </b-form>
        </validation-observer>
        <validation-observer ref="observer" v-slot="{handleSubmit}">
          <b-form @submit.stop.prevent="handleSubmit(doChangeDrinkGenre)">
            <validation-provider name="ジャンル" v-slot="validationContext" :rules="{ required: true}">
              <b-form-group id="genre_id-box">
                <b-row>
                  <b-col cols="3" class="text-right text-left-sp">
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
            <b-button type="submit" variant="warning">変更</b-button>
          </b-form>
        </validation-observer>
        <validation-observer ref="observer" v-slot="{handleSubmit}">
          <b-form @submit.stop.prevent="handleSubmit(doChangeDrinkJan)">
            <validation-provider
              name="Janコード"
              :rules="{ required, numeric, length: 13 }"
              v-slot="validationContext"
            >
              <b-form-group id="Jancode-box">
                <b-row>
                  <b-col cols="3" class="text-right text-left-sp">
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
            <b-button type="submit" variant="warning">変更</b-button>
          </b-form>
        </validation-observer>
        <validation-observer ref="observer" v-slot="{handleSubmit}">
          <b-form @submit.stop.prevent="handleSubmit(doChangeDrinkImage)">
            <validation-provider
              name="画像ファイル名"
              v-slot="validationContext"
              :rules="{ required: true}"
            >
              <b-form-group id="img_file_name-box">
                <b-row>
                  <b-col cols="3" class="text-right text-left-sp">
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
            <b-button type="submit" variant="warning">変更</b-button>
          </b-form>
        </validation-observer>
        <validation-observer ref="observer" v-slot="{handleSubmit}">
          <b-form @submit.stop.prevent="handleSubmit(doChangeDrinkQuantity)">
            <validation-provider
              name="内容量"
              v-slot="validationContext"
              :rules="{ required: true,  numeric,}"
            >
              <b-form-group id="quantity-box">
                <b-row>
                  <b-col cols="3" class="text-right text-left-sp">
                    <label for="quantity">内容量：</label>
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
      drink_id: null,
      drinks: [],
      error: null,
      productname: null,
      genre_id: null,
      genres: [],
      Jancode: null,
      img_file_name: null,
      quantity: null,
    }
  },
  created: function() {
    this.GetDrinks()
    this.GetGenres()
  },

  methods: {
    GetDrinks(){
      axios.get('https://kakaku-real-store.tk:8082/drinks')
      .then((response) => {
        this.drinks = response.data
      })
      .catch(error => {
        // handle error
        this.error = error.response.data.Detail
      })
    },
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
    doChangeDrinkName() {
      if (this.drink_id!=null){

        // サーバへ送信するパラメータ
        const drink_id = this.drink_id
        const params = new URLSearchParams()
        params.append('drink_name', this.productname)
        axios.post('https://kakaku-real-store.tk:8082/drink/name/' + drink_id + '/change', params)
        .then(() => {
          this.$router.push('/drink')
        })
        .catch(error => {
          // handle error
          this.error = error.response.data.Detail
        })
        } else{
          this.error = ("変更する商品を選択してください")
        }
    },
    doChangeDrinkGenre() {
      if (this.drink_id!=null){

        // サーバへ送信するパラメータ
        const params = new URLSearchParams()
        params.append('genre_id', this.genre_id)
        axios.post('https://kakaku-real-store.tk:8082/drink/genre/' + this.drink_id + '/change', params)
      .then(() => {
        this.$router.push('/drink')
      })
      .catch(error => {
        // handle error
        this.error = error.response.data.Detail
      })
      } else{
        this.error = ("変更する商品を選択してください")
      }
    },
    doChangeDrinkJan() {
      if (this.drink_id!=null){

        // サーバへ送信するパラメータ
        const params = new URLSearchParams()
        params.append('jan', this.Jancode)
        axios.post('https://kakaku-real-store.tk:8082/drink/jan/' + this.drink_id + '/change', params)
      .then(() => {
        this.$router.push('/drink')
      })
      .catch(error => {
        // handle error
        this.error = error.response.data.Detail
      })
      } else{
        this.error = ("変更する商品を選択してください")
      }
    },
    doChangeDrinkImage() {
      if (this.drink_id!=null){

        // サーバへ送信するパラメータ
        const params = new URLSearchParams()
        params.append('image', this.img_file_name)
        const drink_id = this.drink_id 
       axios.post('https://kakaku-real-store.tk:8082/drink/iamge/' + drink_id + '/change', params)
      .then(() => {
        this.$router.push('/drink')
      })
      .catch(error => {
        // handle error
        this.error = error.response.data.Detail
      })
      } else{
        this.error = ("変更する商品を選択してください")
      }
    },
    doChangeDrinkQuantity() {
      if (this.drink_id!=null){

        // サーバへ送信するパラメータ
        const params = new URLSearchParams()
        params.append('quantity', this.quantity)
        axios.post('https://kakaku-real-store.tk:8082/drink/quantity/' + this.drink_id + '/change', params)
        .then(() => {
          this.$router.push('/drink')
        })
        .catch(error => {
          // handle error
          this.error = error.response.data.Detail
        })
      } else{
        this.error = ("変更する商品を選択してください")
      }
    },
    getValidationState({ dirty, validated, valid = null }) {
      return dirty || validated ? valid : null;
    },
  }
}
</script>
