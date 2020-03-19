package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/liyinda/ingress-admin/backend/api/models"
	//"github.com/gin-gonic/contrib/sessions"
	"fmt"
	//"net"
	"net/http"
	//"strings"
	"github.com/liyinda/ingress-admin/backend/pkg/e"
	//"github.com/liyinda/viewdns/backend/pkg/util"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strconv"
	//"encoding/json"
)

//查看etcd信息
//func Dnslist(c *gin.Context) {
//    //获取session中的user信息
//    /*session := sessions.Default(c)
//      user := session.Get("user")
//      code := e.INVALID_PARAMS
//      if user == nil {
//      code = e.ERROR_AUTH_SESSION
//      } else {
//      code = e.SUCCESS
//
//      }
//    */
//    //获取POST中json参数
//    //var json models.DNS_A
//    var json models.DomainName
//
//    /*if err := c.ShouldBindJSON(&json); err != nil {
//      code = e.ERROR_NOT_JSON
//      }
//    */
//    //获取url中token, page, limit
//    //token := c.Request.URL.Query().Get("token")
//    /*
//       page := c.Request.URL.Query().Get("page")
//       limit := c.Request.URL.Query().Get("limit")
//
//       pageint, _ := strconv.ParseInt(page, 10, 64)
//       limitint, _ := strconv.ParseInt(limit, 10, 64)
//    */
//
//    //获取用户信息表
//    result, _ := json.DnsList()
//    /*if err != nil{
//      code = e.ERROR
//      } else {
//      code = e.SUCCESS
//      }
//    */
//    //获取总用户数量
//    /*
//       count, err := json.Usercount()
//       if err != nil{
//       code = e.ERROR
//       } else {
//       code = e.SUCCESS
//       }
//    */
//
//    c.JSON(http.StatusOK, gin.H{
//    "msg": result,
//    })
//}

//查看kubernetes中ingress列表信息
func Ingresslist(c *gin.Context) {
	//获取session中的user信息
	//session := sessions.Default(c)
	//user := session.Get("user")

	//定义HTTP状态码
	code := e.INVALID_PARAMS

	//获取ingress信息表
	//items, err := models.ListIngress("kube-system")
	var json models.IngressMeta

	//items, err := json.ListIngress("kube-system")
	items, err := json.ListIngress("developer-center")

	fmt.Println(items)

	if err != nil {
		code = e.ERROR
	} else {
		code = e.SUCCESS
	}

	data := gin.H{
		"items": items,
		"total": 2,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})

}

func AddIngress(c *gin.Context) {
	code := e.INVALID_PARAMS
	//获取url中参数值
	ingressname := c.Query("ingressname")
	namespace := c.Query("namespace")
	host := c.Query("host")
	path := c.Query("path")
	rps := c.Query("rps")
	servicename := c.Query("servicename")
	serviceport := c.Query("serviceport")
	//string转int32
	intServiceport, err := strconv.ParseInt(serviceport, 10, 32)
	//判断port为整数
	if err != nil {
		code = e.ERROR
		fmt.Println(err)
	} else {
		code = e.SUCCESS
		//声明map
		var annotations map[string]string
		annotations = make(map[string]string, 10)
		annotations["nginx.ingress.kubernetes.io/limit-rps"] = rps

		var json models.IngressMeta
		json.IngressName = ingressname
		json.NameSpace = namespace
		json.Host = host
		json.Path = path
		json.Annotations = annotations
		json.ServiceName = servicename
		// IntOrString 这里设置为0, 意为数据类型为整型
		json.ServicePort = intstr.IntOrString{
			Type:   0,
			IntVal: int32(intServiceport),
		}
		//创建ingress
		res := json.CreateIngress("developer-center")
		if res == false {
			code = e.ERROR
		}

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": ingressname,
		})
	}

}

//
//func DelDomain(c *gin.Context) {
//    code := e.INVALID_PARAMS
//    //获取url中参数值
//    domain := c.Query("domain")
//
//    var json models.EtcdRdata
//    res, err := json.DnsDel(domain)
//    if err != nil {
//    code = e.ERROR
//    } else {
//    code = e.SUCCESS
//    }
//    c.JSON(http.StatusOK, gin.H{
//    "code":   code,
//    "data":   domain,
//    "status": res,
//    })
//
//}
