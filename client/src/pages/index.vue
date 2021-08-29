<template>
  <div class="home">
    {{ users.massage }}
    <img src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import axios from 'axios'
export default {
  name: 'home',
  components: {
    HelloWorld
  },
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
