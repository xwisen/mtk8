package client

// GOLANG
//***********************************************
//
//      Filename: etcd.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-05-23 11:06:08
// Last Modified: 2017-05-23 16:44:26
//***********************************************
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/projectcalico/libcalico-go/lib/api"
	"github.com/projectcalico/libcalico-go/lib/client"
)

func NewCalicoClient() *client.Client {
	config := api.NewCalicoAPIConfig()
	config.Spec.DatastoreType = api.EtcdV2
	config.Spec.EtcdConfig.EtcdEndpoints = beego.AppConfig.String("calico_etcd_endpoints")
	c, err := client.New(*config)
	if err != nil {
		logs.Error("init calico client error: %s", err)
	}
	return c
}

var CalicoDefaultClient = NewCalicoClient()
