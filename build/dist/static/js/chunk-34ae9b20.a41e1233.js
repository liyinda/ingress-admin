(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-34ae9b20"],{9406:function(t,a,e){"use strict";e.r(a);var r=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",{attrs:{id:"app"}},[e("div",{staticClass:"allWapper"},[e("el-tabs",{on:{"tab-click":t.handleClick},model:{value:t.activeName,callback:function(a){t.activeName=a},expression:"activeName"}},[e("el-tab-pane",{staticClass:"temp",attrs:{label:"Monitor",name:"first"}},[e("input",{directives:[{name:"model",rawName:"v-model",value:t.monitor_url,expression:"monitor_url"}],attrs:{placeholder:"edit me"},domProps:{value:t.monitor_url},on:{input:function(a){a.target.composing||(t.monitor_url=a.target.value)}}}),t._v(" "),e("p",[t._v("Message is: "+t._s(t.monitor_url))]),t._v(" "),t.ifArr.first?e("iframe",{staticClass:"ifa",attrs:{scrolling:"auto",src:t.monitor_url,frameborder:"0"}}):t._e()]),t._v(" "),e("el-tab-pane",{staticClass:"temp",attrs:{label:"Install",name:"second"}},[t.ifArr.second?e("iframe",{staticClass:"ifa",attrs:{scrolling:"auto",src:"https://github.com/liyinda/ingress-admin",frameborder:"0"}}):t._e()]),t._v(" "),e("el-tab-pane",{staticClass:"temp",attrs:{label:"Help",name:"third"}},[t.ifArr.third?e("iframe",{staticClass:"ifa",attrs:{scrolling:"auto",src:"https://www.google.com/",frameborder:"0"}}):t._e()])],1)],1)])},i=[],s=(e("7f7f"),{name:"App",data:function(){return{monitor_url:"",message:"www.t.com",value:"te",activeName:"first",ifArr:{first:!0,second:!1,third:!1}}},methods:{handleClick:function(t,a){var e=t.name;this.ifArr[e]=!0}}}),n=s,l=(e("f8a7"),e("2877")),o=Object(l["a"])(n,r,i,!1,null,null,null);a["default"]=o.exports},acf4:function(t,a,e){},f8a7:function(t,a,e){"use strict";var r=e("acf4"),i=e.n(r);i.a}}]);