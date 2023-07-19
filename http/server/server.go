package server

import (
	"context"
	"fmt"
	"net/http"
	"snake-tail/config"
	"snake-tail/http/router"
	"time"
)

type Server struct {
	httpSrv   *http.Server
	wait      time.Duration
}

func NewServer(appConfig config.AppConfig) *Server {
	// initialize the router
	r := router.Init()

	address := "0.0.0.0:" + appConfig.Server.Port
	srv := new(Server)
	// Initialize HTTP server
	srv.httpSrv = &http.Server{
		Addr: address,
		// good practice to set timeouts to avoid Slowloris attacks
		WriteTimeout: time.Minute * 2,
		ReadTimeout:  time.Minute * 2,
		IdleTimeout:  time.Minute * 2,

		// pass our instance of gorilla/mux in
		Handler: r,
	}

	return srv
}

func (server *Server) Start() error {
	// run HTTP server in a goroutine so that it doesn't block
	go func() {
		err := server.httpSrv.ListenAndServe()
		if err != nil {
			fmt.Println("http.server.Init", err)
			panic("HTTP server shutting down unexpectedly...")
		}
	}()

	fmt.Println("http.server.Init", fmt.Sprintf("HTTP server listening on %s", server.httpSrv.Addr))

	return nil
}

func (server *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), server.wait)
	defer cancel()

	err := server.httpSrv.Shutdown(ctx)
	if err != nil {
		fmt.Println("http.server.ShutDown", "Unable to stop HTTP server")
	}
}
