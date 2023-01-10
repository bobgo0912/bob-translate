package main

import (
	"context"
	"github.com/bobgo0912/b0b-common/pkg/config"
	"github.com/bobgo0912/b0b-common/pkg/etcd"
	"github.com/bobgo0912/b0b-common/pkg/log"
	"github.com/bobgo0912/b0b-common/pkg/redis"
	"github.com/bobgo0912/b0b-common/pkg/server"
	"github.com/bobgo0912/bob-translate/internal/http"
	"github.com/bobgo0912/bob-translate/internal/rpc"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, ca := context.WithCancel(context.Background())
	log.InitLog()
	newConfig := config.NewConfig(config.Json)
	newConfig.Category = "../config"
	newConfig.InitConfig()
	mainServer := server.NewMainServer()
	etcdClient := etcd.NewClientFromCnf()

	r := server.NewRouter()
	r.Use(http.CORS())
	http.NewTranslateMain(r)

	client, err := redis.NewClient()
	if err != nil {
		log.Panic(err)
	}

	httpServer := server.NewMuxServer(config.Cfg.Host, config.Cfg.Port, r)
	//irisServer := NewIrisServer(config.Cfg.Host, 9999)
	//irisServer.Iris.Get("test", func(c context2.Context) {
	//	log.Info("xxdd")
	//})
	grpcServer := server.NewGrpcServer(config.Cfg.Host, config.Cfg.RpcPort)
	rpc.RegService(grpcServer, client)

	mainServer.AddServer(httpServer)
	mainServer.AddServer(grpcServer)
	//mainServer.AddServer(irisServer)
	err = mainServer.Start(ctx)
	if err != nil {
		log.Panic(err)
	}
	mainServer.Discover(ctx, etcdClient)
	//
	time.Sleep(time.Second * 15)
	address := server.GetRpcNodeAddress("testServers")
	if address == "" {
		log.Info(" bad address")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	ca()
}
