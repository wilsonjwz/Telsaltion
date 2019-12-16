package starters

import (
	log "github.com/sirupsen/logrus"
	"github.com/wilsonj/Telsaltion/base"
	"net/rpc"
	"reflect"

	"net"
)

var rpcServer *rpc.Server

func RpcServer() *rpc.Server {
	base.Check(rpcServer)
	return rpcServer
}

func RpcRegister(ri interface{}) {
	typ := reflect.TypeOf(ri)
	log.Infof("goRPC Register: %s", typ.String())
	err := RpcServer().Register(ri)
	if err != nil {
		log.Panic(err)
	}
}

type GoRpcStarter struct {
	base.BaseStarter
	server *rpc.Server
}

func (g *GoRpcStarter) Init(ctx base.StarterContext) {
	g.server = rpc.NewServer()
	rpcServer = g.server
}

func (g *GoRpcStarter) Start(ctx base.StarterContext) {
	port := ctx.Props().GetDefault("app.rpc.port", "18888")
	// 监听网络端口
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Panic(err)
	}
	log.Info("tcp port listened for rpc:", port)
	go g.server.Accept(listener)

}
