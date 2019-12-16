package starters

import (
	"github.com/wilsonj/Telsaltion/base"
	"github.com/sirupsen/logrus"
	"github.com/tietang/props/kvs"
)

//获取配置starter
var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

type PropsStarter struct {
	base.BaseStarter
}

func (p *PropsStarter)Init(ctx base.StarterContext) {
	//获取配置文件 读取配置文件
	props = ctx.Props()
	//此时 props 就是配置文件 能随时读取
}

func (p *PropsStarter) SetUp(ctx base.StarterContext) {
	logrus.Info(" props setup")
}
func (p *PropsStarter) Start(ctx base.StarterContext) {

}