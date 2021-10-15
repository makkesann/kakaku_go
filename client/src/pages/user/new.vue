<template>
  <div class="login">
    <b-alert v-if="error" variant="danger" show>
      {{error}}
    </b-alert>
    <b-container>
      <div class="bellow-error form">
        <validation-observer ref="observer" v-slot="{handleSubmit}">
          <b-form @submit.stop.prevent="handleSubmit(doLogon)">
            <validation-provider
              name="ユーザー名"
              :rules="{ required: true, min: 3, my_alpha_dash }"
              v-slot="validationContext"
            >
              <b-form-group id="username-box">
                <b-row>
                  <b-col cols="3" class="text-right text-left-sp">
                    <label for="username">ユーザー名：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-input
                      id="username"
                      name="username"
                      v-model="username"
                      :state="getValidationState(validationContext)"
                      aria-describedby="username-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="username-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
                  </b-col>
                </b-row>
              </b-form-group>
            </validation-provider>
            <validation-provider
              name="パスワード"
              :rules="{ required: true, min: 3, my_alpha_dash }"
              v-slot="validationContext"
            >
              <b-form-group id="user_pass-box">
                <b-row>
                  <b-col cols="3" class="text-right text-left-sp">
                    <label for="user_pass">パスワード：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-input
                      id="user_pass"
                      name="user_pass"
                      type="password"
                      v-model="user_pass"
                      :state="getValidationState(validationContext)"
                      aria-describedby="user_pass-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="user_pass-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
                  </b-col>
                </b-row>
              </b-form-group>
            </validation-provider>
            <validation-provider
              name="パスワード確認入力"
              :rules="{ required, min: 3, my_alpha_dash, confirmed: 'パスワード' }"
              v-slot="validationContext"
            >
              <b-form-group id="user_pass_confirmed-box">
                <b-row>
                  <b-col cols="3" class="text-right text-left-sp">
                    <label for="user_pass_confirmed">パスワード確認入力：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-input
                      id="user_pass_confirmed"
                      name="user_pass_confirmed"
                      type="password"
                      v-model="user_pass_confirmed"
                      :state="getValidationState(validationContext)"
                      aria-describedby="user_pass_confirmed-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="user_pass_confirmed-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
                  </b-col>
                </b-row>
              </b-form-group>
            </validation-provider>
            <b-button type="submit" variant="primary">アカウント作成</b-button>
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
  data(){
    return {
      username: null,
      user_pass: null,
      user_pass_confirmed: null,
      error: null,
    }
  },
  beforeCreate() {
    if(this.$store.state.login.id != 0){
      this.$router.push('/drink')
    }
  },
  methods: {
    doLogon() {
      // サーバへ送信するパラメータ
      const params = new URLSearchParams()
      params.append('username', this.username)
      params.append('pass', this.user_pass)
      axios.post('https://kakaku-real-store.tk:8082/user/new', params)
      .then(() => {
        this.$router.push('/drink')
      })
      .catch(() => {
        // handle error
        // this.error = error.response.data.Message
        this.error = "このユーザーネームは既に使われています"
      })
    },
    getValidationState({ dirty, validated, valid = null }) {
      return dirty || validated ? valid : null;
    },
    resetForm() {
      this.username=null
      this.user_pass=null

      this.$nextTick(() => {
        this.$refs.observer.reset();
      });
    },
  }
}
</script>
