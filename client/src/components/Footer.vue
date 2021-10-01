<template>
  <div id="footer">
    <h3>FindDrinks</h3>
    <span>作成者情報&nbsp;&nbsp;&nbsp;&nbsp;</span>
    <a href="#">Twitter</a><span>&nbsp;&nbsp;&nbsp;&nbsp;</span>
    <a href="#">YouTube</a><span>&nbsp;&nbsp;&nbsp;&nbsp;</span>
    <a href="#">GitHub</a><span>&nbsp;&nbsp;&nbsp;&nbsp;</span>
  </div>
</template>

<script>
export default {
  name: 'Footer',
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
