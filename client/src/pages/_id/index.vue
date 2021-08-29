<template>
  <div class="home">
    {{ users.massage }}
    <img src="@/assets/logo.png">
  </div>
</template>

<script>
// @ is an alias to /src
import axios from 'axios'
export default {
  name: 'home',
  data: function(){
    return {
      users: []
    }
  },
  mounted :function(){
    const id = this.$route.params.id
    axios.get('http://localhost:8081/' + id)
          .then(response => this.users = response.data)
          .catch(error => console.log(error))
  }
}
</script>
