package controllers

// GOLANG
//***********************************************
//
//      Filename: network.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-05-23 10:45:38
// Last Modified: 2017-05-23 16:52:19
//***********************************************

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/projectcalico/libcalico-go/lib/api"
	"github.com/xwisen/mtk8/client"
)

type NetWorkController struct {
	BaseController
}

func (nc *NetWorkController) CreateNetwork() {
	logs.Info("start create network ...")
	n := api.IPPool{}
	err := json.Unmarshal(nc.Ctx.Input.RequestBody, &n)
	if err != nil {
		logs.Error("parse request body err: %v", err)
	}
	resp, err := client.CalicoDefaultClient.IPPools().Create(&n)
	if err != nil {
		logs.Error("create network err: %v", err)
		nc.Data["json"] = err.Error()
	} else {
		nc.Data["json"] = &resp
	}
	nc.ServeJSON()
}
func (nc *NetWorkController) DeleteNetwork() {
	logs.Info("start delete network ...")
	n := api.IPPool{}
	err := json.Unmarshal(nc.Ctx.Input.RequestBody, &n)
	if err != nil {
		logs.Error("parse request body err: %v", err)
	}
	err = client.CalicoDefaultClient.IPPools().Delete(n.Metadata)
	if err != nil {
		logs.Error("delete network err: %v", err)
		nc.Data["json"] = err.Error()
	} else {
		nc.Data["json"] = `"status":"删除网络成功"`
	}
	nc.ServeJSON()
}
func (nc *NetWorkController) ListNetwork() {
	logs.Info("start list network ...")
	resp, err := client.CalicoDefaultClient.IPPools().List(api.IPPoolMetadata{})
	if err != nil {
		logs.Error("获取网络列表失败: %s", err)
		nc.Data["json"] = err.Error()
	} else {
		nc.Data["json"] = &resp
	}
	nc.ServeJSON()

}
func (nc *NetWorkController) PatchNetwork() {}
