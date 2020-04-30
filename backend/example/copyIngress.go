package main

import (
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
	//"k8s.io/api/extensions/v1beta1"
	"k8s.io/api/extensions/v1beta1"
	extensions "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"path/filepath"
	"strings"
	//"reflect"
)

type MyIngress struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Annotations struct {
			KubernetesIoIngressClass string `json:"kubernetes.io/ingress.class"`
		} `json:"annotations"`
		CreationTimestamp time.Time `json:"creationTimestamp"`
		Generation        int       `json:"generation"`
		Name              string    `json:"name"`
		Namespace         string    `json:"namespace"`
		ResourceVersion   string    `json:"resourceVersion"`
		SelfLink          string    `json:"selfLink"`
		UID               string    `json:"uid"`
	} `json:"metadata"`
	Spec struct {
		Rules []struct {
			Host string `json:"host"`
			HTTP struct {
				Paths []struct {
					Backend struct {
						ServiceName string `json:"serviceName"`
						ServicePort int    `json:"servicePort"`
					} `json:"backend"`
				} `json:"paths"`
			} `json:"http"`
		} `json:"rules"`
	} `json:"spec"`
	Status struct {
		LoadBalancer struct {
		} `json:"loadBalancer"`
	} `json:"status"`
}

type IngressMeta struct {
	ingressName string
	nameSpace   string
	host        string
	serviceName string
	servicePort intstr.IntOrString
}

func (im *IngressMeta) ReturnIngress() *v1beta1.Ingress {
	ing := &extensions.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:              im.ingressName,
			Namespace:         im.nameSpace,
			CreationTimestamp: metav1.NewTime(time.Now()),
		},
		Spec: extensions.IngressSpec{
			Rules: []extensions.IngressRule{
				{
					Host: im.host,
					IngressRuleValue: extensions.IngressRuleValue{
						HTTP: &extensions.HTTPIngressRuleValue{
							Paths: []extensions.HTTPIngressPath{
								{
									Backend: extensions.IngressBackend{
										ServiceName: im.serviceName,
										ServicePort: im.servicePort,
									},
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

//获取ingress信息,更改ingress中Rules规则并返回新的ingress
func CreateIngress(clientset *kubernetes.Clientset, namespace string) bool {
	//列出namespace下的所有ingress
	ingressList, err := clientset.ExtensionsV1beta1().Ingresses(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	ingressCtrls := ingressList.Items

	if len(ingressCtrls) > 0 {
		for _, ingress := range ingressCtrls {
			//fmt.Printf("Rules %s exists in namespace %s\n", ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServiceName, ingress.Namespace)
			//判断ingress是否已经创建
			if strings.Contains(ingress.Name, "-inject") == true {
				//如果是已经创建就更新ingress的servicePort
				var newServicePort intstr.IntOrString
				for _, existIngress := range ingressCtrls {
					//获取原有ingress中的servierport
					if existIngress.Name == strings.Replace(ingress.Name, "-inject", "", -1) {
						newServicePort = existIngress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServicePort
					}
				}
				u := new(IngressMeta)
				u.ingressName = ingress.Name
				u.nameSpace = namespace
				u.host = ingress.Spec.Rules[0].Host
				u.serviceName = ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServiceName
				u.servicePort = newServicePort
				ing := u.ReturnIngress()
				clientset.ExtensionsV1beta1().Ingresses(namespace).Update(ing)
				//x, err := clientset.ExtensionsV1beta1().Ingresses(namespace).Update(ing)
				//if err != nil {
				//    panic(err.Error())
				//}

			} else {
				//如果没有创建就生成新的ingress
				c := new(IngressMeta)
				c.ingressName = ingress.Name + "-inject"
				c.nameSpace = namespace
				c.host = "2048.dev.app.yyuap.com"
				c.serviceName = ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServiceName
				c.servicePort = ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServicePort
				ing := c.ReturnIngress()
				//new, _ := clientset.ExtensionsV1beta1().Ingresses(namespace).Create(ing)
				clientset.ExtensionsV1beta1().Ingresses(namespace).Create(ing)
				//fmt.Println(new)
				if err != nil {
					//panic(err.Error())
					fmt.Println(err.Error())
				}

			}
		}
	} else {
		fmt.Println("no ingress found")
	}
	time.Sleep(1 * time.Second)
	return true
}

func main() {
	// 配置 k8s 集群外 kubeconfig 配置文件
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", filepath.Join("/home", ".kube", "config"), "absolute path to the kubeconfig file")
	flag.Parse()

	//在 kubeconfig 中使用当前上下文环境，config 获取支持 url 和 path 方式
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// 根据指定的 config 创建一个新的 clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		CreateIngress(clientset, "c87e2267-1001-4c70-bb2a-ab41f3b81aa3")
	}
}
