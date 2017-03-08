// GOLANG
//***********************************************
//
//      Filename: deployment.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-03-07 10:36:33
// Last Modified: 2017-03-08 17:48:43
//***********************************************

package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/xwisen/mtk8/client"
	metav1 "k8s.io/client-go/pkg/api/v1"
	extensionsV1beta1 "k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

type DeploymentController struct {
	BaseController
	BaseInfo
}

func (dc *DeploymentController) CreateDeployment() {
	d := extensionsV1beta1.Deployment{}
	err := json.Unmarshal(dc.Ctx.Input.RequestBody, &d)
	if err != nil {
		logs.Error("parse request body err: %v", err)
	}
	//logs.Info("%v", &d)
	resp, err := client.K8sDefaultClient.ExtensionsV1beta1().Deployments("default").Create(&d)
	if err != nil {
		logs.Error("create deployment response err: %v", err)
		dc.Data["json"] = err.Error()
	} else {
		dc.Data["json"] = &resp
	}
	dc.ServeJSON()
}

func (dc *DeploymentController) DeleteDeployment() {
}

func (dc *DeploymentController) PatchDeployment() {
}

func (dc *DeploymentController) ListDeployment() {
	depolymentLists, err := client.K8sDefaultClient.ExtensionsV1beta1().Deployments("").List(metav1.ListOptions{})
	if err != nil {
		logs.Error("list deployment error : %v", err)
	}
	dc.Data["json"] = &depolymentLists
	dc.ServeJSON()
}
