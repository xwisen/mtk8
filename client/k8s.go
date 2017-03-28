// GOLANG
//***********************************************
//
//      Filename: k8s.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-03-08 13:23:03
// Last Modified: 2017-03-28 10:51:33
//***********************************************

package client

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
)

func NewK8sClient() *kubernetes.Clientset {
	config := rest.Config{
		Host:     "http://" + beego.AppConfig.String("k8s_ip") + ":" + beego.AppConfig.String("k8s_port"),
		Insecure: false,
	}
	logs.Info("k8s info is: %s, Insecure is: %v", config.Host, config.Insecure)
	clientset, err := kubernetes.NewForConfig(&config)
	if err != nil {
		logs.Error("init k8s clientset failed, err is :%v", err)
		os.Exit(1)
	}
	return clientset
}

var K8sDefaultClient = NewK8sClient()
