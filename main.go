// GOLANG
//***********************************************
//
//      Filename: main.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-03-03 11:22:42
// Last Modified: 2017-03-03 14:01:16
//***********************************************

package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = version
	app.Usage = "k8s API demo"
	app.Commands = commands
	app.Flags = globalFlags
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
		os.Exit(1)
	}
}
