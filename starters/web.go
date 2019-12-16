package starters

import (
	"github.com/wilsonjwz/Telsaltion/base"
)

var apiInitializerRegister *base.InitializeRegister = new(base.InitializeRegister)

//注册web api初始化对象
func RegisterApi(ai base.Initializer){
	apiInitializerRegister.Register(ai)
}

//获取注册的web api初始化对象
func GetApiInitializers() []base.Initializer {
	return apiInitializerRegister.Initializers
}

type WebApiStarter struct {
	base.BaseStarter
}

func(w *WebApiStarter) Setup(ctx base.StarterContext)  {
	for _, v:= range GetApiInitializers() {
		v.Init()
	}

}