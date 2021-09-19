<template>
  <div class="home">
    <b-container>
      <b-row class="my-1">
        <b-col sm="3">
          <label for="username">ID:</label>
        </b-col>
        <b-col sm="9">
          <b-form-input :state="username_state" id="username" type="text" v-model="username" placeholder="Enter your username"></b-form-input>
        </b-col>
      </b-row>
      <b-row class="my-1">
        <b-col sm="3">
          <label for="user_pass">パスワード:</label>
        </b-col>
        <b-col sm="9">
          <b-form-input :state="user_pass_state" id="user_pass" type="text" v-model="user_pass" placeholder="Enter your password"></b-form-input>
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
  data() {
    return {
      username: "",
      user_pass: "",
    }
  },
  beforeCreate() {
    if(this.$store.state.login.id != 0){
      this.$router.push('/drink')
    }
  },
  computed:{
    username_state(){
      return this.username.length == 0 ? false : true
    },
    user_pass_state(){
      return this.user_pass.length == 0 ? false : true
    },
  },
  methods: {
    doLogin() {
      // サーバへ送信するパラメータ
      const params = new URLSearchParams()
      params.append('username', this.username)
      params.append('pass', this.user_pass)
      axios.post('http://54.65.204:8082/login', params)
      .then(response => {
        this.$store.dispatch('login/doSetID',response.data.Value.ID)
          //一覧ページに遷移する
        this.$router.push('/drink')
      })
      .catch(error => {
        // handle error
        console.log(error)
      })
    },
  }
}
</script>
