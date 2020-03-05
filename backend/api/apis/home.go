package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/liyinda/ingress-admin/backend/api/models"
	//"github.com/gin-gonic/contrib/sessions"
	//"fmt"
	//"net"
	"net/http"
	//"strings"
	//"github.com/liyinda/viewdns/backend/pkg/util"
	"github.com/liyinda/ingress-admin/backend/pkg/e"
	//"strconv"
	//"encoding/json"
)

//查看etcd信息
func Dnslist(c *gin.Context) {
	//获取session中的user信息
	/*session := sessions.Default(c)
	  user := session.Get("user")
	  code := e.INVALID_PARAMS
	  if user == nil {
	      code = e.ERROR_AUTH_SESSION
	  } else {
	      code = e.SUCCESS

	  }
	*/
	//获取POST中json参数
	//var json models.DNS_A
	var json models.DomainName

	/*if err := c.ShouldBindJSON(&json); err != nil {
	      code = e.ERROR_NOT_JSON
	  }
	*/
	//获取url中token, page, limit
	//token := c.Request.URL.Query().Get("token")
	/*
	   page := c.Request.URL.Query().Get("page")
	   limit := c.Request.URL.Query().Get("limit")

	   pageint, _ := strconv.ParseInt(page, 10, 64)
	   limitint, _ := strconv.ParseInt(limit, 10, 64)
	*/

	//获取用户信息表
	result, _ := json.DnsList()
	/*if err != nil{
	      code = e.ERROR
	  } else {
	      code = e.SUCCESS
	  }
	*/
	//获取总用户数量
	/*
	   count, err := json.Usercount()
	   if err != nil{
	       code = e.ERROR
	   } else {
	       code = e.SUCCESS
	   }
	*/

	c.JSON(http.StatusOK, gin.H{
		"msg": result,
	})
}

//查看kubernetes中ingress列表信息
func Ingresslist(c *gin.Context) {
	//获取session中的user信息
	/*session := sessions.Default(c)
	  user := session.Get("user")
	  code := e.INVALID_PARAMS
	  if user == nil {
	      code = e.ERROR_AUTH_SESSION
	  } else {
	      code = e.SUCCESS

	  }
	*/
	//获取POST中json参数
	//var json models.DNS_A
	var json models.DomainName

	/*if err := c.ShouldBindJSON(&json); err != nil {
	      code = e.ERROR_NOT_JSON
	  }
	*/
	//获取url中token, page, limit
	//token := c.Request.URL.Query().Get("token")
	/*
	   page := c.Request.URL.Query().Get("page")
	   limit := c.Request.URL.Query().Get("limit")

	   pageint, _ := strconv.ParseInt(page, 10, 64)
	   limitint, _ := strconv.ParseInt(limit, 10, 64)
	*/

	//获取用户信息表
	result, _ := json.DnsList()
	/*if err != nil{
	      code = e.ERROR
	  } else {
	      code = e.SUCCESS
	  }
	*/
	//获取总用户数量
	/*
	   count, err := json.Usercount()
	   if err != nil{
	       code = e.ERROR
	   } else {
	       code = e.SUCCESS
	   }
	*/

	c.JSON(http.StatusOK, gin.H{
		"msg": result,
	})
}

func Table(c *gin.Context) {
	//获取session中的user信息
	/*session := sessions.Default(c)
	  user := session.Get("user")
	  code := e.INVALID_PARAMS
	  if user == nil {
	      code = e.ERROR_AUTH_SESSION
	  } else {
	      code = e.SUCCESS

	  }
	*/

	code := e.INVALID_PARAMS
	num := 2

	var json models.DomainName
	//获取etcd中dns列表
	items, err := json.DnsList()
	if err != nil {
		code = e.ERROR
	} else {
		code = e.SUCCESS
	}

	data := gin.H{
		"items": items,
		"total": num,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})

}

//func AddDomain(c *gin.Context) {
//	code := e.INVALID_PARAMS
//	//获取url中参数值
//	domain := c.Query("domain")
//	ipv4 := c.Query("host")
//	ttl := c.Query("ttl")
//	//判断ip合法性
//	addr := net.ParseIP(ipv4)
//	if addr == nil {
//		code = e.ERROR
//	}
//	//string转int
//	intTTL, err := strconv.Atoi(ttl)
//	//判断ttl为整数
//	if err != nil {
//		code = e.ERROR
//		fmt.Println(err)
//	} else {
//		//判断ip合法性
//		addr := net.ParseIP(ipv4)
//		if addr == nil {
//			code = e.ERROR
//		} else {
//			code = e.SUCCESS
//
//			var json models.EtcdRdata
//			json.DnsAdd(domain, ipv4, intTTL)
//
//			c.JSON(http.StatusOK, gin.H{
//				"code": code,
//				"data": domain,
//			})
//		}
//	}
//
//}
//
//func DelDomain(c *gin.Context) {
//	code := e.INVALID_PARAMS
//	//获取url中参数值
//	domain := c.Query("domain")
//
//	var json models.EtcdRdata
//	res, err := json.DnsDel(domain)
//	if err != nil {
//		code = e.ERROR
//	} else {
//		code = e.SUCCESS
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code":   code,
//		"data":   domain,
//		"status": res,
//	})
//
//}
