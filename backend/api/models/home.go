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

type IngressMeta struct {
	Name        string            `json:"name"`
	NameSpace   string            `json:"namespace"`
	Annotations map[string]string `json:"annotations"`
	Host        string            `json:"host"`
	Paths       []IngressPaths    `json:"children"`
	Rps         string            `json:"rps"`
}

type IngressPaths struct {
	Path        string             `json:"path"`
	ServiceName string             `json:"servicename"`
	ServicePort intstr.IntOrString `json:"serviceport"`
}

type Children struct {
	Path        string `json:"path"`
	ServiceName string `json:"servicename"`
	ServicePort string `json:"serviceport"`
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

//列出ingress中annotations信息
func (ingmeta *IngressMeta) ListAnnotations(namespace string) (ingresses []IngressMeta, err error) {
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
			ingmeta.Rps = ingress.Annotations["nginx.ingress.kubernetes.io/limit-rps"]
			//结构体拼接
			ingresses = append(ingresses, *ingmeta)
		}
	} else {
		fmt.Println("no ingress found")
	}

	return
}

func (ingmeta *IngressMeta) ReturnIngress(n int) *extensions.Ingress {
	paths := []extensions.HTTPIngressPath{}

	//将extensions.HTTPIngressPath结构体数组进行累加,n为数组长度
	for i := 0; i < n; i++ {
		ingressBackend := extensions.IngressBackend{
			ServiceName: ingmeta.Paths[i].ServiceName,
			ServicePort: ingmeta.Paths[i].ServicePort,
		}
		path := extensions.HTTPIngressPath{
			Backend: ingressBackend,
			Path:    ingmeta.Paths[i].Path,
		}
		paths = append(paths, path)
	}

	//fmt.Println(paths)
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
							Paths: paths,
						},
					},
				},
			},
		},
	}
	return ing
}

//在namespace中创建该命名空间中的ingress
func (ingmeta *IngressMeta) CreateIngress(namespace string, n int) bool {
	//初始化clientset
	clientset, _ := kubedata.NewClientset()

	//使用结构体,创建ingress
	ing := ingmeta.ReturnIngress(n)
	_, err := clientset.ExtensionsV1beta1().Ingresses(namespace).Create(ing)
	//_, err := clientset.ExtensionsV1beta1().Ingresses(namespace).Update(ing)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

//在namespace中更新该命名空间中的ingress
func (ingmeta *IngressMeta) UpdateIngress(namespace string, n int) bool {
	//初始化clientset
	clientset, _ := kubedata.NewClientset()

	//使用结构体,创建ingress
	ing := ingmeta.ReturnIngress(n)
	//_, err := clientset.ExtensionsV1beta1().Ingresses(namespace).Create(ing)
	_, err := clientset.ExtensionsV1beta1().Ingresses(namespace).Update(ing)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

//在namespace中删除该命名空间中的ingress
func (ingmeta *IngressMeta) DeleteIngress(namespace string) bool {
	//初始化clientset
	clientset, _ := kubedata.NewClientset()

	//删除ingress
	err := clientset.ExtensionsV1beta1().Ingresses(namespace).Delete(ingmeta.Name, &metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
