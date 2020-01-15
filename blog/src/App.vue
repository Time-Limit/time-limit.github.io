<template>
  <div id="app">
    <body>
      <Menu class="nav" mode="horizontal" :theme="navTheme" @on-select="changeRoute">
        <div class="navlayout">
          <Menu-item><Icon type="md-images"/>手账</Menu-item>
          <Menu-item><Icon type="ios-paper-plane"/>树莓派</Menu-item>
          <Menu-item><Icon type="ios-planet"/>读书笔记</Menu-item>
        </div>
      </Menu>
      <router-view/>
    </body>
  </div>
</template>

<script>
export default {
  name: 'App',

  data () {
    return {
      navTheme: 'light'
    }
  },

  methods: {
    changeRoute () {
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.nav {
  position: fixed !important;
  left: 0;
  top: 0;
  width: 100vw;
}

.navlayout {
  margin-left: 50px;
}

</style>

