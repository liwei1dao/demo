package services

import (
	"fmt"
	"lego_demo/sys/db"

	"github.com/liwei1dao/lego/base/cluster"
	"github.com/liwei1dao/lego/sys/log"
)

type ServiceBase struct {
	cluster.ClusterService
}

func (this *ServiceBase) InitSys() {
	this.ClusterService.InitSys()
	if err := db.OnInit(this.Service.GetSettings().Sys["db"]); err != nil {
		panic(fmt.Sprintf("statr db err:%v", err))
	} else {
		log.Infof("Sys db Init success !")
	}
}
