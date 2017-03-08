package controllers

// GOLANG
//***********************************************
//
//      Filename: base.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-03-07 10:48:27
// Last Modified: 2017-03-08 13:38:41
//***********************************************

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// BaseInfo 包含一些公用信息，如应用联系人、软负载地址、发布域
type BaseInfo struct {
	Name   string `json:"name,omitempty"`   //APP联系人姓名
	Number string `json:"number,omitempty"` //APP联系人电话
	Email  string `json:"email,omitempty"`  //APP联系人邮箱
	Vendor string `json:"vendor,omitempty"` //APP厂商
}
