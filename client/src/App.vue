<template>
  <div id="app">
    {{ users.message }}
    <p>sdfssddfsadf</p>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'App',
  data: function(){
    return {
      users: []
    }
  },
  mounted :function(){
    axios.get('http://localhost:8081')
          .then(response => this.users = response.data)
          .catch(error => console.log(error))
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
