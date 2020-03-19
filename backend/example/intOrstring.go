package main

import (
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type service struct {
	Name string             `json:"name"`
	Host string             `json:"host"`
	Port intstr.IntOrString `json:"port"`
}

func main() {
	test := "123"

	nginxService := service{
		Name: "nginx",
		Host: "192.168.1.2",
		Port: intstr.IntOrString{
			Type: 1, // 这里设置为0, 意为数据类型为整型
			//IntVal: 2222,
			StrVal: test,
		},
	}

	fmt.Println(nginxService)

	jsonBytes, _ := json.Marshal(nginxService)
	fmt.Println(string(jsonBytes))

}
