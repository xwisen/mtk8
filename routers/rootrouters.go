// GOLANG
//***********************************************
//
//      Filename: rootrouters.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-03-07 10:24:38
// Last Modified: 2017-05-30 11:21:32
//***********************************************

package routers

import (
	"github.com/astaxie/beego"
	"github.com/xwisen/mtk8/controllers"
)

func init() {
	beego.Router("/api/deployment", &controllers.DeploymentController{}, "get:ListDeployment")
	beego.Router("/api/deployment", &controllers.DeploymentController{}, "post:CreateDeployment")
	beego.Router("/api/deployment", &controllers.DeploymentController{}, "delete:DeleteDeployment")
	beego.Router("/api/deployment", &controllers.DeploymentController{}, "patch:PatchDeployment")
	beego.Router("/api/network", &controllers.NetWorkController{}, "get:ListNetwork")
	beego.Router("/api/network", &controllers.NetWorkController{}, "post:CreateNetwork")
	beego.Router("/api/network", &controllers.NetWorkController{}, "delete:DeleteNetwork")
	beego.Router("/api/network", &controllers.NetWorkController{}, "patch:PatchNetwork;put:PatchNetwork")
}
