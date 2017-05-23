// GOLANG
//***********************************************
//
//      Filename: main.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-03-07 08:48:22
// Last Modified: 2017-05-23 10:42:08
//***********************************************
package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/xwisen/mtk8/routers"
)

func main() {
	logs.Info("Start mtk8 .........")
	beego.Run()
}
