(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-34ae9b20"],{9406:function(a,t,e){"use strict";e.r(t);var r=function(){var a=this,t=a.$createElement,e=a._self._c||t;return e("div",{attrs:{id:"app"}},[e("div",{staticClass:"allWapper"},[e("el-tabs",{on:{"tab-click":a.handleClick},model:{value:a.activeName,callback:function(t){a.activeName=t},expression:"activeName"}},[e("el-tab-pane",{staticClass:"temp",attrs:{label:"Monitor",name:"first"}},[e("input",{directives:[{name:"model",rawName:"v-model",value:a.monitor_url,expression:"monitor_url"}],attrs:{placeholder:"edit me"},domProps:{value:a.monitor_url},on:{input:function(t){t.target.composing||(a.monitor_url=t.target.value)}}}),a._v(" "),e("p",[a._v("Message is: "+a._s(a.monitor_url))]),a._v(" "),a.ifArr.first?e("iframe",{staticClass:"ifa",attrs:{scrolling:"auto",src:a.monitor_url,frameborder:"0"}}):a._e()]),a._v(" "),e("el-tab-pane",{staticClass:"temp",attrs:{label:"Install",name:"second"}},[a.ifArr.second?e("iframe",{staticClass:"ifa",attrs:{scrolling:"auto",src:"https://github.com/liyinda/ingress-admin",frameborder:"0"}}):a._e()]),a._v(" "),e("el-tab-pane",{staticClass:"temp",attrs:{label:"Help",name:"third"}},[a.ifArr.third?e("iframe",{staticClass:"ifa",attrs:{scrolling:"auto",src:"https://www.baidu.com/",frameborder:"0"}}):a._e()])],1)],1)])},i=[],s=(e("7f7f"),{name:"App",data:function(){return{monitor_url:"",message:"www.t.com",value:"te",activeName:"first",ifArr:{first:!0,second:!1,third:!1}}},methods:{handleClick:function(a,t){var e=a.name;this.ifArr[e]=!0}}}),n=s,l=(e("f8a7"),e("2877")),o=Object(l["a"])(n,r,i,!1,null,null,null);t["default"]=o.exports},acf4:function(a,t,e){},f8a7:function(a,t,e){"use strict";var r=e("acf4"),i=e.n(r);i.a}}]);