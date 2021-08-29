<template>
  <div>
    <Header/>
    <router-view/>
    <p>This is a default layout</p>
    <Footer/>
  </div>
</template>

<script>
import Header from '@/components/Header.vue'
import Footer from '../components/Footer.vue'
export default {
  components: {
    Header,
    Footer,
  },
}
</script>