package apis

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liyinda/ingress-admin/backend/api/models"
	//"net"
	"github.com/liyinda/ingress-admin/backend/pkg/e"
	"net/http"
	"strings"
	//"github.com/liyinda/viewdns/backend/pkg/util"
	"encoding/json"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strconv"
)

//查看kubernetes中ingress列表信息
func Ingresslist(c *gin.Context) {
	//获取session中的user信息
	session := sessions.Default(c)
	user := session.Get("user")

	//定义HTTP状态码
	code := e.INVALID_PARAMS
	if user == nil {
		code = e.ERROR_AUTH_SESSION
	} else {
		code = e.SUCCESS
	}

	//获取ingress信息表
	//items, err := models.ListIngress("kube-system")
	var json models.IngressMeta

	//items, err := json.ListIngress("kube-system")
	items, err := json.ListIngress("developer-center")

	//fmt.Println(items)

	if err != nil {
		code = e.ERROR
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

//查看ingress中annotations信息
func Annotationslist(c *gin.Context) {
	//获取session中的user信息
	session := sessions.Default(c)
	user := session.Get("user")

	//定义HTTP状态码
	code := e.INVALID_PARAMS
	if user == nil {
		code = e.ERROR_AUTH_SESSION
	} else {
		code = e.SUCCESS
	}

	//获取ingress信息表
	//items, err := models.ListIngress("kube-system")
	var json models.IngressMeta

	//items, err := json.ListIngress("kube-system")
	items, err := json.ListAnnotations("developer-center")

	fmt.Println(items)

	if err != nil {
		code = e.ERROR
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

//在kubernetes中添加ingress
func AddIngress(c *gin.Context) {
	//获取session中的user信息
	session := sessions.Default(c)
	user := session.Get("user")

	//定义HTTP状态码
	code := e.INVALID_PARAMS
	if user == nil {
		code = e.ERROR_AUTH_SESSION
	} else {
		code = e.SUCCESS
	}

	//获取url中参数值
	object := c.Query("object")
	childrenArray := c.QueryArray("children[]")

	//object反序列化json格式
	var Json models.IngressMeta

	objectByte := []byte(object)
	err := json.Unmarshal(objectByte, &Json)
	if err != nil {
		fmt.Println("unmarshal error: ", err)
	}

	//children数组反序列化json格式
	var jsonChildren models.IngressPaths
	var Children models.Children

	//申明结构体数组
	jsonChildrenArray := []models.IngressPaths{}

	for _, children := range childrenArray {
		childrenByte := []byte(children)
		err := json.Unmarshal(childrenByte, &Children)
		if err != nil {
			fmt.Println("unmarshal error: ", err)
		}

		jsonChildren.Path = Children.Path
		jsonChildren.ServiceName = Children.ServiceName
		//string转int32
		intServiceport, _ := strconv.ParseInt(Children.ServicePort, 10, 32)
		jsonChildren.ServicePort = intstr.IntOrString{
			Type:   0,
			IntVal: int32(intServiceport),
		}
		jsonChildrenArray = append(jsonChildrenArray, jsonChildren)
	}
	Json.Paths = jsonChildrenArray

	//获取url中参数值
	hashmap := c.Query("hashmap")

	//hashmap反序列化Map格式
	//var mapResult map[string]interface{}
	var mapResult map[string]string
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(hashmap), &mapResult); err != nil {
		fmt.Println("unmarshal error: ", err)
	}
	fmt.Println(mapResult)

	//vue中hashmap无法传递. 需要在后端重新构造
	var newMapResult map[string]string
	newMapResult = make(map[string]string)
	for key, value := range mapResult {
		if strings.Contains(key, "nginxingresskubernetesio") {
			newkey := strings.Replace(key, "nginxingresskubernetesio", "nginx.ingress.kubernetes.io", -1)
			newMapResult[newkey] = value

		} else {
			newkey := strings.Replace(key, " kubernetesio/ingressclass", " kubernetes.io/ingress.class", -1)
			newMapResult[newkey] = value
		}

	}

	//将转化的map结果传递给annotations
	Json.Annotations = newMapResult
	//声明map
	//var annotations map[string]string
	//annotations = make(map[string]string, 10)
	//annotations["nginx.ingress.kubernetes.io/limit-rps"] = rps

	//创建ingress
	res := Json.CreateIngress("developer-center", len(childrenArray))
	if res == false {
		code = e.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": Json.Name,
	})
	//}

}

//在kubernetes中更新ingress
func EditIngress(c *gin.Context) {
	//获取session中的user信息
	session := sessions.Default(c)
	user := session.Get("user")

	//定义HTTP状态码
	code := e.INVALID_PARAMS
	if user == nil {
		code = e.ERROR_AUTH_SESSION
	} else {
		code = e.SUCCESS
	}

	//获取url中参数值
	object := c.Query("object")
	childrenArray := c.QueryArray("children[]")

	//object反序列化json格式
	var Json models.IngressMeta

	objectByte := []byte(object)
	err := json.Unmarshal(objectByte, &Json)
	if err != nil {
		fmt.Println("unmarshal error: ", err)
	}

	//children数组反序列化json格式
	var jsonChildren models.IngressPaths
	var Children models.Children

	//申明结构体数组
	jsonChildrenArray := []models.IngressPaths{}

	for _, children := range childrenArray {
		childrenByte := []byte(children)
		err := json.Unmarshal(childrenByte, &Children)
		if err != nil {
			fmt.Println("unmarshal error: ", err)
		}

		jsonChildren.Path = Children.Path
		jsonChildren.ServiceName = Children.ServiceName
		//string转int32
		intServiceport, _ := strconv.ParseInt(Children.ServicePort, 10, 32)
		jsonChildren.ServicePort = intstr.IntOrString{
			Type:   0,
			IntVal: int32(intServiceport),
		}
		jsonChildrenArray = append(jsonChildrenArray, jsonChildren)
	}
	Json.Paths = jsonChildrenArray

	//获取url中参数值
	hashmap := c.Query("hashmap")

	//hashmap反序列化Map格式
	var mapResult map[string]string
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(hashmap), &mapResult); err != nil {
		fmt.Println("unmarshal error: ", err)
	}

	//vue中hashmap无法传递. 需要在后端重新构造
	var newMapResult map[string]string
	newMapResult = make(map[string]string)
	for key, value := range mapResult {
		if strings.Contains(key, "nginxingresskubernetesio") {
			newkey := strings.Replace(key, "nginxingresskubernetesio", "nginx.ingress.kubernetes.io", -1)
			newMapResult[newkey] = value

		} else {
			newkey := strings.Replace(key, " kubernetesio/ingressclass", " kubernetes.io/ingress.class", -1)
			newMapResult[newkey] = value
		}

	}

	//将转化的map结果传递给annotations
	Json.Annotations = newMapResult

	//创建ingress
	res := Json.UpdateIngress("developer-center", len(childrenArray))
	if res == false {
		code = e.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": Json.Name,
	})
	//}

}

//在kubernetes中删除指定ingress
func DelIngress(c *gin.Context) {
	//获取session中的user信息
	session := sessions.Default(c)
	user := session.Get("user")

	//定义HTTP状态码
	code := e.INVALID_PARAMS
	if user == nil {
		code = e.ERROR_AUTH_SESSION
	} else {
		code = e.SUCCESS
	}

	//获取url中参数值
	name := c.Query("name")

	var Json models.IngressMeta

	//删除ingress
	Json.Name = name
	res := Json.DeleteIngress("developer-center")
	if res == false {
		code = e.ERROR
	} else {
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"data":   Json.Name,
		"status": res,
	})

}
