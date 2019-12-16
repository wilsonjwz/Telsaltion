package Telsaltion

import (
	log "github.com/sirupsen/logrus"
	"github.com/tietang/props/kvs"
	"github.com/wilsonjwz/Telsaltion/base"

	"reflect"
)

type BootApplication struct {
	IsTest     bool
	conf       kvs.ConfigSource
	starterCtx base.StarterContext
}

//程序注册入口
func New(conf kvs.ConfigSource) *BootApplication {
	boot := &BootApplication{conf: conf, starterCtx: base.StarterContext{}}
	//设定配置
	boot.starterCtx.SetProps(conf)
	return boot
}

func (boot *BootApplication) Start() {
	//1. 初始化starter
	boot.init()
	//2. 安装starter
	boot.setup()
	//3. 启动starter
	boot.start()
}

//程序初始化
func (boot *BootApplication) init() {
	log.Info("Initializing starters...")
	for _, v := range base.GetStarters() {
		typ := reflect.TypeOf(v)
		log.Debugf("Initializing: PriorityGroup=%d,Priority=%d,type=%s", v.PriorityGroup(), v.Priority(), typ.String())
		v.Init(boot.starterCtx)
	}
}

//程序安装
func (boot *BootApplication) setup() {

	log.Info("Setup starters...")
	for _, v := range base.GetStarters() {
		typ := reflect.TypeOf(v)
		//展示 对象名字
		log.Debug("Setup: ", typ.String())
		v.Setup(boot.starterCtx)
	}

}

//程序开始运行，开始接受调用
func (boot *BootApplication) start() {

	log.Info("Starting starters...")
	for i, v := range base.GetStarters() {

		typ := reflect.TypeOf(v)
		log.Debug("Starting: ", typ.String())
		if v.StartBlocking() {
			if i+1 == len(base.GetStarters()) {
				v.Start(boot.starterCtx)
			} else {
				go v.Start(boot.starterCtx)
			}
		} else {
			v.Start(boot.starterCtx)
		}

	}
}

//程序开始运行，开始接受调用
func (boot *BootApplication) Stop() {

	log.Info("Stoping starters...")
	for _, v := range base.GetStarters() {
		typ := reflect.TypeOf(v)
		log.Debug("Stoping: ", typ.String())
		v.Stop(boot.starterCtx)
	}
}
