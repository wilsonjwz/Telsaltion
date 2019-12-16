package starters

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	irisrecover "github.com/kataras/iris/middleware/recover"
	"github.com/sirupsen/logrus"
	"github.com/wilsonj/Telsaltion/base"
	"time"
)

var isrsApp *iris.Application

func Iris() *iris.Application {
	return isrsApp
}

type IrisServerStarter struct {
	base.BaseStarter
}

func (i *IrisServerStarter) Init(ctx base.StarterContext) {
	//创建实例
	isrsApp = initIris()
	//日志组件配置和扩展
	irisLogger := isrsApp.Logger()
	irisLogger.Install(logrus.StandardLogger())
}

func (i *IrisServerStarter) Start(ctx base.StarterContext) {
	//把路由信息打印到控制台
	routes := Iris().GetRoutes()
	for _, r := range routes {
		logrus.Info(r.Trace())
	}
	//启动iris
	port := ctx.Props().GetDefault("app.server.port", "9081")
	Iris().Run(iris.Addr(":" + port))
}

func (i *IrisServerStarter) StartBlocking() bool {
	return true
}

func initIris() *iris.Application {
	app := iris.New()
	app.Use(irisrecover.New())
	cfg := logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		Query:              true,
		Columns:            false,
		MessageContextKeys: nil,
		MessageHeaderKeys:  nil,
		LogFunc: func(now time.Time, latency time.Duration,
			status, ip, method, path string,
			message interface{},
			headerMessage interface{}) {
			app.Logger().Infof("| %s | %s | %s | %s | %s | %s | %s | %s",
				now.Format("2006-01-02.15:04:05.000000"),
				latency.String(), status, ip, method, path, headerMessage, message,
			)
		},
		Skippers: nil,
	}
	app.Use(logger.New(cfg))
	return app
}
