package starters

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/wilsonj/Telsaltion/base"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func XormEngine() *xorm.Engine {
	base.Check(engine)
	return engine
}

type XormEngineStarter struct {
	base.BaseStarter
}

func (x *XormEngineStarter) Init(ctx base.StarterContext) {
	logrus.Info("xorm init")
}

func (x *XormEngineStarter) Setup(ctx base.StarterContext) {
	logrus.Info("xorm setup")
	conf := ctx.Props()
	driverName := conf.GetDefault("mysql.driverName", "mysql")
	user := conf.GetDefault("mysql.root", "root")
	pwd := conf.GetDefault("mysql.password", "")
	database := conf.GetDefault("mysql.database", "test")
	e, err := xorm.NewEngine(driverName, user+":"+pwd+"@/"+database+"?charset=utf8")
	if err != nil {
		panic(err)
	}
	logrus.Info(e.Ping())
	engine = e
}