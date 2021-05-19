package backstage

import (
	"lego_demo/comm"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/lib/modules/console"
)

func NewModule() core.IModule {
	m := new(BackStage)
	return m
}

type BackStage struct {
	console.Console
	options *Options
}

func (this *BackStage) GetType() core.M_Modules {
	return comm.BackStageModule
}

func (this *BackStage) NewOptions() (options core.IModuleOptions) {
	return new(Options)
}

func (this *BackStage) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	this.options = options.(*Options)
	err = this.Console.Init(service, module, options)
	return
}

func (this *BackStage) OnInstallComp() {
	this.Console.OnInstallComp()
}
