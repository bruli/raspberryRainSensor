package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Config struct {
	serverAddr string
}

func NewConfig(serverAddr string) *Config {
	return &Config{serverAddr: serverAddr}
}

type Server struct {
	conf *Config
	wait time.Duration
}

func NewServer(conf *Config) *Server {
	return &Server{conf: conf, wait: 15 * time.Second}
}

func (s *Server) Run() {
	srv := &http.Server{
		Handler:      newRouter().build(),
		Addr:         s.conf.serverAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// createZone a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), s.wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
