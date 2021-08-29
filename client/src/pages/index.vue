<template>
  <div class="home">
    <b-container>
      <b-row>
        <h1>くそわろた</h1>
      </b-row>
    </b-container>
  </div>
</template>

<script>
// @ is an alias to /src
import axios from 'axios'
export default {
  name: 'home',
  components: {
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
