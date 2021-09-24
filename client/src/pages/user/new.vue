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
              name="ユーザー名"
              :rules="{ required: true, min: 3, my_alpha_dash }"
              v-slot="validationContext"
            >
              <b-form-group id="username">
                <b-row>
                  <b-col cols="3" class="text-right">
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
              :rules="{ required: true, min: 3 }"
              v-slot="validationContext"
            >
              <b-form-group id="user_pass">
                <b-row>
                  <b-col cols="3" class="text-right">
                    <label for="user_pass">パスワード：</label>
                  </b-col>
                  <b-col cols="9">
                    <b-form-input
                      id="user_pass"
                      name="user_pass"
                      v-model="user_pass"
                      :state="getValidationState(validationContext)"
                      aria-describedby="user_pass-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="user_pass-live-feedback">{{ validationContext.errors[0] }}</b-form-invalid-feedback>
                  </b-col>
                </b-row>
              </b-form-group>
            </validation-provider>
            <b-button type="submit" variant="primary" @click="doLogon" :disabled="handleSubmit.invalid || !handleSubmit.validated">アカウント作成</b-button>
            <b-button class="ml-2" @click="resetForm()">フォームリセット</b-button>
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
      error: null
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
      axios.post('http://localhost:8082/user/new', params)
      .then(response => {
        console.log(response)
        // this.$store.dispatch('login/doSetID',response.data.Value.ID)
          //一覧ページに遷移する
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
