package database

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"path/filepath"
)

var Staticpath *string

//解决flag redefined错误，由于在mvc框架中，flag是不能多次调用的，通过加缓存解决
func GetRuntimePath() *string {
	if Staticpath != nil {
		return Staticpath
	}
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", filepath.Join("/home", ".kube", "config"), "absolute path to the kubeconfig file")
	flag.Parse()
	Staticpath = kubeconfig
	return Staticpath
}

//创建kubernetes中clientset连接
func NewClientset() (*kubernetes.Clientset, error) {

	// 配置 k8s 集群外 kubeconfig 配置文件
	path := GetRuntimePath()

	//在 kubeconfig 中使用当前上下文环境，config 获取支持 url 和 path 方式
	config, err := clientcmd.BuildConfigFromFlags("", *path)
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// 根据指定的 config 创建一个新的 clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//defer Close()

	return clientset, err
}
