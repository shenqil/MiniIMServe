package main

import (
	config "MiniIMServe/configs"
	"MiniIMServe/listeners"
	"MiniIMServe/server"
	"encoding/json"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	// 1.创建全局日志对象
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// 2.加载配置
	config.MustLoad("./configs/config.toml")
	jsonData, err := json.Marshal(config.C)
	if err != nil {
		log.Error("日志序列化失败",
			slog.String("error", err.Error()))
		return
	}
	log.Info("全局配置", "config", string(jsonData))

	// 3.创建服务
	server := server.New(log)

	// 4.创建ws监听器
	wsCfg := config.C.WebSocket
	ws := listeners.NewWebsocket(&listeners.WebsocketConfig{
		ID:      wsCfg.ID,
		Address: wsCfg.Address,
	})
	err = server.AddListener(ws)
	if err != nil {
		log.Error("ws监听器添加失败",
			slog.String("error", err.Error()))
		return
	}

	// 5.启动服务
	go func() {
		err := server.Serve()
		if err != nil {
			log.Error("ws监听器添加失败",
				slog.String("error", err.Error()))
			return
		}
	}()

	<-done
	server.Log.Warn("caught signal, stopping...")
	_ = server.Close()
	server.Log.Info("main.go finished")
}
