package controllers

// GOLANG
//***********************************************
//
//      Filename: network.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-05-23 10:45:38
// Last Modified: 2017-05-30 11:38:45
//***********************************************

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/projectcalico/libcalico-go/lib/api"
	"github.com/xwisen/mtk8/client"
	"github.com/xwisen/mtk8/util"
	"io/ioutil"
	"net/http"
	"time"
)

const DOCKER_SOCK = "/var/run/docker.sock"

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
		logs.Info("创建ippool 响应: %+v", resp)
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
func (nc *NetWorkController) PatchNetwork() {
	//logs.Info("请求方法为:%s", nc.Ctx.Input.Method())
	u := &util.Transport{
		DialTimeout:           100 * time.Millisecond,
		RequestTimeout:        1 * time.Second,
		ResponseHeaderTimeout: 1 * time.Second,
	}
	u.RegisterLocation("127.0.0.1:4243", DOCKER_SOCK)
	tr := &http.Transport{}
	tr.RegisterProtocol(util.Scheme, u)

	hcli := &http.Client{Transport: tr}
	dresp, err := hcli.Get("http+unix://127.0.0.1:4243/version")
	defer dresp.Body.Close()
	if err != nil {
		logs.Error("connect to sock %s err: %v", DOCKER_SOCK, err)
		nc.Data["json"] = err.Error()
	} else {
		logs.Info("resp from %s is %+v", DOCKER_SOCK, dresp)
		if dresp.StatusCode != 200 {
			logs.Error("connect to sock %s resp code: %d", DOCKER_SOCK, dresp.StatusCode)
		}
		j, err := ioutil.ReadAll(dresp.Body)
		if err != nil {
			logs.Error("从HTTP响应中读取返回值失败,错误为:%s", err)
		}
		//nc.Data["json"] = &j
		nc.Ctx.WriteString(string(j))
	}
	nc.ServeJSON()
}
