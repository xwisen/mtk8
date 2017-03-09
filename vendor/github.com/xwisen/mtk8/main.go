// GOLANG
//***********************************************
//
//      Filename: main.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-03-07 08:48:22
// Last Modified: 2017-03-08 13:34:13
//***********************************************
package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/xwisen/mtk8/routers"
)

func main() {
	routers.RootRouters()
	logs.Info("Start mtk8 .........")
	beego.Run()
}
