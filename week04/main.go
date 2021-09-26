package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"os"
	"os/signal"
	"syscall"
	"week04/api"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	return &Server{engine: api.ServerEngine}
}

func (p *Server) Start() {
	p.engine.Run()
}

func ExitFunc(s os.Signal) {
	os.Exit(0)
}

func SignalFunc() {
	// 处理信号
	signalChan := make(chan os.Signal)

	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range signalChan {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				{
					ExitFunc(s)
				}
			}
		}
	}()
}

func main() {
	wire.Build(api.ApiRegistry, SignalFunc)
	server := NewServer()
	server.Start()
}
