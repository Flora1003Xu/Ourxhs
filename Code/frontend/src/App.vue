<template>
  <div id="app">
    <div>
      <keep-alive>
  <router-view></router-view>
    </keep-alive>
    
 <!-- <router-view v-if="!$route.meta.keepAlive" class="router-view" ></router-view> -->
  <!-- 可以让它跳转到index.js里 -->
   </div>
  </div>

</template>

<script>

export default {
  name: 'app',
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #000000;
}
.el-radio__input.is-checked .el-radio__inner {
  background: #f12d41 !important;
  border-color: #f71f38 !important;
}
.el-menu {
  background-color: #fbfafacc;
  color:#030303cc !important;
}
</style>
