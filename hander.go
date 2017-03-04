package main

// GOLANG
//***********************************************
//
//      Filename: hander.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-03-03 14:56:39
// Last Modified: 2017-03-03 17:18:08
//***********************************************

import (
	"encoding/json"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/client-go/pkg/api/v1"
	//extensionsV1beta1 "k8s.io/client-go/pkg/apis/extensions/v1beta1"
	"log"
	"net/http"
)

type HandeStruct struct {
	ClientSet *kubernetes.Clientset
}

func (h *HandeStruct) createDeployment(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create depolyment is here!"))
}

func (h *HandeStruct) listDeployment(w http.ResponseWriter, r *http.Request) {
	depolymentLists, err := h.ClientSet.ExtensionsV1beta1().Deployments("").List(metav1.ListOptions{})
	if err != nil {
		log.Printf(" err :%v", err)
	}
	d := depolymentLists.GetSelfLink()
	depolymentListsResp, err := json.Marshal(&d)
	if err != nil {
		log.Printf(" err :%v", err)
	}
	w.Write(depolymentListsResp)
}
