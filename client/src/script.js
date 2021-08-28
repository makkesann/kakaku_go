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