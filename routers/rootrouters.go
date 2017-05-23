// GOLANG
//***********************************************
//
//      Filename: rootrouters.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-03-07 10:24:38
// Last Modified: 2017-05-23 10:48:43
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
}
