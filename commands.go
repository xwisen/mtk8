package main

// GOLANG
//***********************************************
//
//      Filename: commands.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-03-03 13:31:52
// Last Modified: 2017-03-03 15:59:05
//***********************************************

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"k8s.io/client-go/kubernetes"
	//extensionsV1beta1 "k8s.io/client-go/pkg/apis/extensions/v1beta1"
	"k8s.io/client-go/rest"
	"log"
	"net/http"
	"os"
	"time"
)

var version = "v1.0"
var commands = []cli.Command{
	{
		Name:    "version",
		Aliases: []string{"c"},
		Usage:   "print app version",
		Action: func(ctx *cli.Context) {
			fmt.Printf("App version is: %s\n", version)
		},
	},
	{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "server for request",
		Flags:   serverFlags,
		Action:  server,
	},
}

func server(ctx *cli.Context) {
	log.SetPrefix(fmt.Sprint("%s >>>>>:", time.Now()))
	log.Printf("Start server ......")
	listenAddress := ctx.String("listen")
	k8Address := ctx.String("k8s")
	config := rest.Config{
		Host:     k8Address,
		Insecure: true,
	}
	clientset, err := kubernetes.NewForConfig(&config)
	if err != nil {
		log.Printf("create k8s clientset err :%v", err)
		os.Exit(1)
	}

	h := HandeStruct{
		ClientSet: clientset,
	}

	router := mux.NewRouter()
	router.HandleFunc("/create", h.createDeployment)
	router.HandleFunc("/list", h.listDeployment)

	srv := &http.Server{
		Handler:      router,
		Addr:         listenAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
