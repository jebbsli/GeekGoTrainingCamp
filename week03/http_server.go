package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

type httpServer struct {
	server *http.Server
}


func Hello(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("hello server!"))
}
func (p *httpServer) InitServer() {
	p.server = &http.Server{}
	p.server.Addr = ":8080"

	http.HandleFunc("/hello", Hello)
}

func (p *httpServer) Start() error {
	return p.server.ListenAndServe()
}

func main() {
	server := &httpServer{}
	server.InitServer()

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	c, errCtx := errgroup.WithContext(ctx)

	c.Go(func() error {
		return server.Start()
	})

	c.Go(func() error {
		<- errCtx.Done()
		return server.server.Shutdown(errCtx)
	})

	signalChan := make(chan os.Signal, 0)
	signal.Notify(signalChan)

	c.Go(func() error {
		for {
			select {
			case <- errCtx.Done():
				return errCtx.Err()

			case <-signalChan:
				cancel()
			}
		}

		return nil
	})

	if err := c.Wait(); err != nil {
		fmt.Println("errgroup wait error:", err)
		return
	}

	fmt.Println("server down ...")
}
