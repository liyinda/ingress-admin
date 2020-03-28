<template>
  <div id="app">
<div class="allWapper">
  <!-- logo -->
    <!-- 顶部tabs -->
   <el-tabs v-model="activeName" @tab-click="handleClick">
    <el-tab-pane class="temp" label="Monitor" name="first">
      <iframe v-if="ifArr.first" class="ifa" scrolling=auto src="http://172.20.58.151/grafana/d/Vyy4xH2mz/node-exporter-stats?refresh=5m&orgId=1&from=1585205789196&to=1585207589197&var-node=10.16.112.10?" frameborder="0"></iframe>
    </el-tab-pane>
    <el-tab-pane class="temp"  label="Install" name="second">
       <iframe v-if="ifArr.second"   class="ifa" scrolling=auto src="https://segmentfault.com/" frameborder="0"></iframe>
    </el-tab-pane>
    <el-tab-pane class="temp"  label="Help" name="third">
       <iframe v-if="ifArr.third"  class="ifa" scrolling=auto src="https://www.baidu.com/" frameborder="0"></iframe>
    </el-tab-pane>
  </el-tabs>

</div>
<!-- </div> -->
    <!-- <router-view/> -->
  </div>
</template>

<script>
export default {
  name: 'App',
  data(){
    return{
        activeName: 'first',
        ifArr:{
          first:true,
          second:false,
          third:false
        }
    }
  },
  methods:{
      handleClick(tab, event) {
       let flagName=tab.name
       this.ifArr[flagName]=true
      }
  }
}
</script>

import { mapGetters } from 'vuex'

<style>
body{
  margin:0;
  padding:5px;
}
.ifa{
    width:100%;
    height:100%;
  }
  .el-tabs__content{
    border: 1px solid red;
    border-top:none;
    position: absolute;
    top: 62px;
    left: 0;
    bottom: 0;
    right: 0;
  }
  .allWapper{
    display:flex;
    border-bottom:1px solid #e4e7ed;
  }
  .el-tabs__header{
    margin-bottom:0px;
  }
  .el-tabs__header .el-tabs__item{
    margin:8px 0;
    font-size:16px;
    padding-left:29px;
  }
  .temp{
      width:100%;
    height:100%;
  }
.myLogo{
      width: 200px;
    height: 53px;
    margin-right:35px;
}
.myLogo img{
  width:100%;
}
</style>
