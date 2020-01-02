package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/CoffeeChat/server/src/internal/logic/conf"
	"github.com/CoffeeChat/server/src/internal/logic/rpcserver"
	"github.com/CoffeeChat/server/src/pkg/db"
	"github.com/CoffeeChat/server/src/pkg/helper"
	"github.com/CoffeeChat/server/src/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

var (
	defaultConf string
	pidFileName = "server.pid"
)

func waitExit(c chan os.Signal) {
	for i := range c {
		switch i {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			logger.Sugar.Info("receive exit signal ", i.String(), ",exit...")

			logger.Sugar.Infof("close db connect")
			db.DefaultManager.UnInt()

			os.Exit(0)
		}
	}
}

func init() {
	flag.String("conf", "logic-example.toml", defaultConf)
}

func main() {

	logger.InitLogger("log/log.log", "debug")
	defer logger.Sugar.Sync()

	_, err := toml.DecodeFile(defaultConf, conf.DefaultLogicConfig)
	if err != nil {
		_, err := toml.DecodeFile("logic-example.toml", conf.DefaultLogicConfig)
		if err != nil {
			logger.Sugar.Fatal("load conf file err:", err.Error())
			return
		}
	}

	// 记录进程id
	if err := helper.WritePid(pidFileName); err != nil {
		logger.Sugar.Error("write pid file error:", err.Error(), ",exit...")
		return
	}

	// init db
	err = db.DefaultManager.Init(conf.DefaultLogicConfig.Db)
	if err != nil {
		logger.Sugar.Fatal("db init failed:", err.Error())
		return
	}

	// init cache
	redis := conf.DefaultLogicConfig.Redis
	err = db.InitCache(redis.Ip, redis.Port, redis.Password, redis.KeyPrefix, redis.Pool)
	if err != nil {
		logger.Sugar.Fatal("redis cache init failed:", err.Error())
		return
	}

	// start rpc server
	go rpcserver.StartRpcServer(conf.DefaultLogicConfig.ListenIp, conf.DefaultLogicConfig.ListenPort)

	// before exit, cleanup something ...
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	waitExit(c)
}
