(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-34ae9b20"],{9406:function(a,t,e){"use strict";e.r(t);var r=function(){var a=this,t=a.$createElement,e=a._self._c||t;return e("div",{attrs:{id:"app"}},[e("div",{staticClass:"allWapper"},[e("el-tabs",{on:{"tab-click":a.handleClick},model:{value:a.activeName,callback:function(t){a.activeName=t},expression:"activeName"}},[e("el-tab-pane",{staticClass:"temp",attrs:{label:"Monitor",name:"first"}},[a.ifArr.first?e("iframe",{staticClass:"ifa",attrs:{scrolling:"auto",src:"http://172.20.58.151/grafana/d/Vyy4xH2mz/node-exporter-stats?refresh=5m&orgId=1&from=1585205789196&to=1585207589197&var-node=10.16.112.10?",frameborder:"0"}}):a._e()]),a._v(" "),e("el-tab-pane",{staticClass:"temp",attrs:{label:"Install",name:"second"}},[a.ifArr.second?e("iframe",{staticClass:"ifa",attrs:{scrolling:"auto",src:"https://segmentfault.com/",frameborder:"0"}}):a._e()]),a._v(" "),e("el-tab-pane",{staticClass:"temp",attrs:{label:"Help",name:"third"}},[a.ifArr.third?e("iframe",{staticClass:"ifa",attrs:{scrolling:"auto",src:"https://www.baidu.com/",frameborder:"0"}}):a._e()])],1)],1)])},s=[],i=(e("7f7f"),{name:"App",data:function(){return{activeName:"first",ifArr:{first:!0,second:!1,third:!1}}},methods:{handleClick:function(a,t){var e=a.name;this.ifArr[e]=!0}}}),n=i,l=(e("f8a7"),e("2877")),c=Object(l["a"])(n,r,s,!1,null,null,null);t["default"]=c.exports},acf4:function(a,t,e){},f8a7:function(a,t,e){"use strict";var r=e("acf4"),s=e.n(r);s.a}}]);