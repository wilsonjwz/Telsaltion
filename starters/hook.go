package starters

import (
	"github.com/sirupsen/logrus"
	"github.com/wilsonjwz/Telsaltion/base"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

var callbacks []func()

func Register(fn func()) {
	callbacks = append(callbacks, fn)
}

type HookStarter struct {
	base.BaseStarter
}

func (s *HookStarter) Init(ctx base.StarterContext) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		for {
			c := <-sigs
			logrus.Info("notify: ", c)
			for _, fn := range callbacks {
				fn()
			}
			break
			os.Exit(0)
		}
	}()
}

func (s *HookStarter) Start(ctx base.StarterContext) {
	starters := base.GetStarters()

	for _, s := range starters {
		typ := reflect.TypeOf(s)
		logrus.Infof("【Register Notify Stop】:%s.Stop()", typ.String())
		//注册所有stop方法
		Register(func() {
			s.Stop(ctx)
		})
	}
}
