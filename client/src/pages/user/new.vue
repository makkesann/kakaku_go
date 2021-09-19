<template>
  <div class="home">
    <b-alert v-if="error" variant="danger" show>
      {{error}}
    </b-alert>
    <b-container>
      <b-row class="my-1">
        <b-col sm="3">
          <label for="username">username:</label>
        </b-col>
        <b-col sm="9">
          <b-form-input id="username" type="text" v-model="username"></b-form-input>
        </b-col>
      </b-row>
      <b-row class="my-1">
        <b-col sm="3">
          <label for="user_pass">パスワード:</label>
        </b-col>
        <b-col sm="9">
          <b-form-input id="user_pass" type="text" v-model="user_pass"></b-form-input>
        </b-col>
      </b-row>
      <b-row>
        <b-button pill v-on:click="doLogin">送信</b-button>
      </b-row>
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
      error: ""
    }
  },
  methods: {
    doLogin() {
      // サーバへ送信するパラメータ
      const params = new URLSearchParams()
      params.append('username', this.username)
      params.append('pass', this.user_pass)
      axios.post('http://54.65.204.164:8082/user/new', params)
      .then(response => {
        console.log(response)
        // this.$store.dispatch('login/doSetID',response.data.Value.ID)
          //一覧ページに遷移する
        this.$router.push('/drink')
      })
      .catch(error => {
        // handle error
        this.error = error.response.data.Message
      })
    },
  }
}
</script>
