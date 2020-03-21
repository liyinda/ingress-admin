package models

//package main

import (
	"fmt"
	kubedata "github.com/liyinda/ingress-admin/backend/api/database"
	extensions "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"time"
	//"k8s.io/client-go/kubernetes"
	//"reflect"
)

//type IngressMeta struct {
//	IngressName string
//	NameSpace   string
//	Host        string
//	ServiceName string
//	ServicePort intstr.IntOrString
//	Path        string
//	Rps         string
//	Annotations map[string]string
//}

//type IngressMeta struct {
//	APIVersion string `json:"apiVersion"`
//	Kind       string `json:"kind"`
//	Metadata   struct {
//		Annotations       map[string]string
//		CreationTimestamp time.Time `json:"creationTimestamp"`
//		Generation        int       `json:"generation"`
//		Name              string    `json:"name"`
//		Namespace         string    `json:"namespace"`
//		ResourceVersion   string    `json:"resourceVersion"`
//		SelfLink          string    `json:"selfLink"`
//		UID               string    `json:"uid"`
//	} `json:"metadata"`
//	Spec struct {
//		Rules []struct {
//			Host string `json:"host"`
//			HTTP struct {
//				Paths []struct {
//					Backend struct {
//						ServiceName string `json:"serviceName"`
//						ServicePort int    `json:"servicePort"`
//					} `json:"backend"`
//				} `json:"paths"`
//			} `json:"http"`
//		} `json:"rules"`
//	} `json:"spec"`
//	Status struct {
//		LoadBalancer struct {
//		} `json:"loadBalancer"`
//	} `json:"status"`
//}

//type IngressMeta struct {
//	Name        string            `json:"name"`
//	NameSpace   string            `json:"namespace"`
//	Annotations map[string]string `json:"annotations"`
//	Host        string            `json:"host"`
//	Paths       []struct {
//		Path    string `json:"path"`
//		Backend struct {
//			ServiceName string             `json:"serviceName"`
//			ServicePort intstr.IntOrString `json:"servicePort"`
//		} `json:"backend"`
//	} `json:"paths"`
//}

type IngressMeta struct {
	Name        string            `json:"name"`
	NameSpace   string            `json:"namespace"`
	Annotations map[string]string `json:"annotations"`
	Host        string            `json:"host"`
	Paths       []IngressPaths    `json:"paths"`
}

type IngressPaths struct {
	Path        string             `json:"path"`
	ServiceName string             `json:"serviceName"`
	ServicePort intstr.IntOrString `json:"servicePort"`
}

//在namespace中获取该命名空间中的ingress信息列表
func (ingmeta *IngressMeta) ListIngress(namespace string) (ingresses []IngressMeta, err error) {
	//func (ingmeta *IngressMeta) ListIngress(namespace string) (ingresses []string, err error) {
	//初始化clientset
	clientset, _ := kubedata.NewClientset()

	//列出namespace下的所有ingress
	ingressList, err := clientset.ExtensionsV1beta1().Ingresses(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	ingressCtrls := ingressList.Items

	if len(ingressCtrls) > 0 {
		for _, ingress := range ingressCtrls {

			//结构体赋值
			ingmeta.Name = ingress.Name
			ingmeta.NameSpace = namespace
			ingmeta.Host = ingress.Spec.Rules[0].Host

			//遍历所有Paths中的后端服务和Path
			pathsCtrls := ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths

			//声明结构体数组和结构体
			ingpathsArray := []IngressPaths{}
			var ingpaths IngressPaths

			for _, paths := range pathsCtrls {
				//fmt.Printf("paths=%v\n", paths)

				ingpaths.Path = paths.Path
				ingpaths.ServiceName = paths.Backend.ServiceName
				ingpaths.ServicePort = paths.Backend.ServicePort
				//将结构体进行累加
				ingpathsArray = append(ingpathsArray, ingpaths)
				//fmt.Printf("num=%v ingpaths=%v\n", n, ingpathsArray)
			}
			//将结构体数组赋值给IngressMeta.Paths
			ingmeta.Paths = ingpathsArray
			//ingmeta.Rps = ingress.Annotations["nginx.ingress.kubernetes.io/limit-rps"]
			//结构体拼接
			ingresses = append(ingresses, *ingmeta)

		}
	} else {
		fmt.Println("no ingress found")
	}

	return
}

func (ingmeta *IngressMeta) ReturnIngress() *extensions.Ingress {
	ing := &extensions.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:              ingmeta.Name,
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
										ServiceName: ingmeta.Paths[0].ServiceName,
										ServicePort: ingmeta.Paths[0].ServicePort,
									},
									Path: ingmeta.Paths[0].Path,
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
