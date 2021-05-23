package console

import (
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/lib"
	"github.com/liwei1dao/lego/lib/modules/console"
)

func NewModule() core.IModule {
	m := new(Console)
	return m
}

type Console struct {
	console.Console
	options *Options
}

func (this *Console) GetType() core.M_Modules {
	return lib.SM_ConsoleModule
}

func (this *Console) NewOptions() (options core.IModuleOptions) {
	return new(Options)
}

func (this *Console) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	this.options = options.(*Options)
	err = this.Console.Init(service, module, options)
	return
}

func (this *Console) OnInstallComp() {
	this.Console.OnInstallComp()
}
