package models

//package main

import (
	"encoding/json"
	"fmt"
	kubedata "github.com/liyinda/ingress-admin/backend/api/database"
	extensions "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"time"
	//"k8s.io/client-go/kubernetes"
	//"reflect"
)

type IngressMeta struct {
	IngressName string
	NameSpace   string
	Host        string
	ServiceName string
	ServicePort intstr.IntOrString
	Path        string
	Rps         string
	Annotations map[string]string
}

//在namespace中获取该命名空间中的ingress信息列表
func (ingmeta *IngressMeta) ListIngress(namespace string) (ingresses []IngressMeta, err error) {
	//初始化clientset
	clientset, _ := kubedata.NewClientset()

	//创建map和slince
	var m map[string]interface{}
	var s []map[string]interface{}

	var value string

	//列出namespace下的所有ingress
	ingressList, err := clientset.ExtensionsV1beta1().Ingresses(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	ingressCtrls := ingressList.Items

	if len(ingressCtrls) > 0 {
		for _, ingress := range ingressCtrls {

			//将获取的ingress中的信息JSON序列化
			m = make(map[string]interface{})

			m["id"] = "650000199402171139"
			m["ingressname"] = ingress.Name
			m["namespace"] = namespace
			m["host"] = ingress.Spec.Rules[0].Host
			m["servicename"] = ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServiceName
			m["serviceport"] = ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServicePort
			m["path"] = ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Path
			m["rps"] = ingress.Annotations["nginx.ingress.kubernetes.io/limit-rps"]
			s = append(s, m)
			data, err := json.Marshal(s)
			if err != nil {
				fmt.Printf("json error %v\n", err)
			}
			value = string(data)
			fmt.Println(value)
		}
	} else {
		fmt.Println("no ingress found")
	}

	err = json.Unmarshal([]byte(value), &ingresses)

	return
}

func (ingmeta *IngressMeta) ReturnIngress() *extensions.Ingress {
	ing := &extensions.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:              ingmeta.IngressName,
			Namespace:         ingmeta.NameSpace,
			CreationTimestamp: metav1.NewTime(time.Now()),
			Annotations:       ingmeta.Annotations,
		},
		Spec: extensions.IngressSpec{
			Rules: []extensions.IngressRule{
				{
					Host: ingmeta.Host,
					IngressRuleValue: extensions.IngressRuleValue{
						HTTP: &extensions.HTTPIngressRuleValue{
							Paths: []extensions.HTTPIngressPath{
								{
									Backend: extensions.IngressBackend{
										ServiceName: ingmeta.ServiceName,
										ServicePort: ingmeta.ServicePort,
									},
									Path: ingmeta.Path,
								},
							},
						},
					},
				},
			},
		},
	}
	return ing
}

//在namespace中创建该命名空间中的ingress
func (ingmeta *IngressMeta) CreateIngress(namespace string) bool {
	//初始化clientset
	clientset, _ := kubedata.NewClientset()

	//使用结构体,创建ingress
	ing := ingmeta.ReturnIngress()
	_, err := clientset.ExtensionsV1beta1().Ingresses(namespace).Create(ing)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

//func main() {
//    //func (ingress *Ingress) GetList() v1beta1.Ingress {
//    clientset, _ := kubeconfig.NewClientset()
//    list, _ := ListIngress(clientset, "kube-system")
//    fmt.Println("type:", reflect.TypeOf(list))
//    fmt.Println(list)
//    //return list
//}
