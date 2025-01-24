package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/tiamxu/builder/logic"

	httpkit "github.com/tiamxu/kit/http"
	"github.com/tiamxu/kit/log"
)

var (
	cfg *Config
	// name = "builder"
)

func init() {
	loadConfig()
	if err := cfg.Initial(); err != nil {
		log.Fatalf("Config initialization failed: %v", err)
	}
}
func newServer() *http.Server {
	e := httpkit.NewGin(cfg.HttpSrv)
	logic.RegisterHttpRoute(e)
	return httpkit.StartServer(e, cfg.HttpSrv)
}

func main() {
	// 启动 HTTP 服务
	srv := newServer()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	httpkit.ShutdownServer(srv)

	log.Println("Server exiting")
}
