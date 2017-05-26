// GOLANG
//***********************************************
//
//      Filename: deployment.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-03-07 10:36:33
// Last Modified: 2017-05-25 23:19:12
//***********************************************

package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/xwisen/mtk8/client"
	//metav1 "k8s.io/client-go/pkg/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//labels "k8s.io/apimachinery/pkg/labels"
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
	d := extensionsV1beta1.Deployment{}
	err := json.Unmarshal(dc.Ctx.Input.RequestBody, &d)
	if err != nil {
		logs.Error("parse request body err: %v", err)
	}
	err = client.K8sDefaultClient.ExtensionsV1beta1().Deployments(d.Namespace).Delete(d.Name, &metav1.DeleteOptions{})
	if err != nil {
		logs.Error("delete deployment response err: %v", err)
		dc.Data["json"] = err.Error()
	} else {
		dc.Data["json"] = `"status":"成功删除应用"`
	}
	dc.ServeJSON()
}

func (dc *DeploymentController) PatchDeployment() {
}

func (dc *DeploymentController) ListDeployment() {
	deploymentLists, err := client.K8sDefaultClient.ExtensionsV1beta1().Deployments("").List(metav1.ListOptions{})
	if err != nil {
		logs.Error("list deployment error : %v", err)
	}
	for did, dp := range deploymentLists.Items {
		logs.Info("deployment %s in %s,id is %d", dp.GetName(), dp.GetNamespace(), did)
		rsLists, err := client.K8sDefaultClient.ExtensionsV1beta1().ReplicaSets("").List(metav1.ListOptions{LabelSelector: fmt.Sprintf("run=%s", dp.Spec.Selector.MatchLabels["run"])})
		if err != nil {
			logs.Error("list rs error : %v", err)
		}
		for rid, rs := range rsLists.Items {
			logs.Info("rs %s in %s,id is %d,belong to dp %s", rs.GetName(), rs.GetNamespace(), rid, dp.GetName())
			podLists, err := client.K8sDefaultClient.CoreV1().Pods("").List(metav1.ListOptions{LabelSelector: fmt.Sprintf("pod-template-hash=%s", rs.Spec.Selector.MatchLabels["pod-template-hash"])})
			if err != nil {
				logs.Error("list pod error : %v", err)
			}
			for pid, pod := range podLists.Items {
				logs.Info("pod %s in %s ,id is %d,belong to rs %s", pod.GetName(), pod.GetNamespace(), pid, rs.GetName())
			}
		}
	}
	dc.Data["json"] = &deploymentLists
	dc.ServeJSON()
}
