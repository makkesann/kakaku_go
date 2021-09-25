<template>
  <div class="login">
    <b-alert v-if="error" variant="danger" show>
      {{error}}
    </b-alert>
    <b-container>
      <div>
        <validation-observer ref="observer" v-slot="handleSubmit">
          <b-form @submit.stop.prevent="handleSubmit(onSubmit)">
            <validation-provider name="商品名" v-slot="validationContext">
              <b-form-group id="drink_id" label="商品名" label-for="drink_id">
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

                <b-form-invalid-feedback id="drink_id-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
              </b-form-group>
            </validation-provider>
            <b-button type="submit" variant="warning" @click="doDeleteDrink" :disabled="handleSubmit.invalid || !handleSubmit.validated">削除</b-button>
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
      error: null
    }
  },
  created: function() {
    this.GetDrinks()
  },

  methods: {
    GetDrinks(){
      axios.get('http://54.65.204.164:8082/drinks')
      .then(response => {
        this.drinks = response.data
      })
      .catch(error => {
        // handle error
        console.log(error)
      })
    },
    // 商品情報を削除する
    doAddDrink() {
      // サーバへ送信するパラメータ
      axios.post('http://54.65.204.164:8082/drink/' + this.drink_id + 'delete')
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
