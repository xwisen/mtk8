package main

// GOLANG
//***********************************************
//
//      Filename: flags.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-03-03 13:20:52
// Last Modified: 2017-03-03 15:37:52
//***********************************************

import (
	//"fmt"
	"github.com/urfave/cli"
)

var globalFlags = []cli.Flag{}
var serverFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "listen",
		Usage:  "listen address for server",
		EnvVar: "LISTEN",
		Value:  ":8888",
	},
	cli.StringFlag{
		Name:   "k8s",
		Usage:  "k8s master API address",
		EnvVar: "K8S_SERVER",
		Value:  "http://127.0.0.1:8080",
	},
}
