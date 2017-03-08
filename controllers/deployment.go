// GOLANG
//***********************************************
//
//      Filename: deployment.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-03-07 10:36:33
// Last Modified: 2017-03-08 13:55:22
//***********************************************

package controllers

import (
	//"encoding/json"
	"github.com/astaxie/beego/logs"
	metav1 "k8s.io/client-go/pkg/api/v1"
	//"k8s.io/client-go/tools/clientcmd"
	"github.com/xwisen/mtk8/client"
)

type DeploymentController struct {
	BaseController
	BaseInfo
}

func (dc *DeploymentController) CreateDeployment() {
	dc.Ctx.WriteString("CreateDeployment handler\n")
}

func (dc *DeploymentController) DeleteDeployment() {
	dc.Ctx.WriteString("DeleteDeployment handler\n")
}

func (dc *DeploymentController) PatchDeployment() {
	dc.Ctx.WriteString("PatchDeployment handler\n")
}

func (dc *DeploymentController) ListDeployment() {
	depolymentLists, err := client.K8sDefaultClient.ExtensionsV1beta1().Deployments("").List(metav1.ListOptions{})
	if err != nil {
		logs.Error("list deployment error : %v", err)
	}
	dc.Data["json"] = &depolymentLists
	dc.ServeJSON()
}
