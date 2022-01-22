package server

import (
	"context"
	"example.com/web-crawler-golang/service"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"example.com/web-crawler-golang/endpoint"
)
const (
	serverPort  = "SERVER_PORT"
)

// Run starts the HTTP server
func Run() {
	log.SetFormatter(&log.JSONFormatter{})

	handler := setUpServer()

	serverPort, exists := os.LookupEnv(serverPort)
	if !exists || len(serverPort) == 0 {
		log.Fatalf("Port for web-crawer with envName : %s is not specified", serverPort)
	}

	addr := fmt.Sprintf(":%s",serverPort)
	srv := &http.Server{Addr: addr, Handler: handler}
	go func() {
		log.Info("Starting server")

		err := srv.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				log.Info("Server shut down. Waiting for connections to drain.")
			} else {
				log.WithError(err).
					WithField("server_port", srv.Addr).
					Fatal("failed to start server")
			}
		}
	}()

	// Wait for an interrupt
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)    // interrupt signal sent from terminal
	signal.Notify(sigint, syscall.SIGTERM) // sigterm signal sent from system
	<-sigint

	log.Info("Shutting down server")

	attemptGracefulShutdown(srv)
}

func setUpServer() http.Handler {
	mux := http.NewServeMux()
	crawlLogger := log.WithField("endpoint", "getUrlTask")
	crawlerService := service.NewCrawlerService()
	mux.Handle("/crawl/submit", endpoint.SubmitTask(crawlerService, crawlLogger))
	mux.Handle("/crawl/read/", endpoint.GetCrawledUrlsById(crawlerService, crawlLogger))
	return mux
}

func attemptGracefulShutdown(srv *http.Server) {
	if err := shutdownServer(srv, 25*time.Second); err != nil {
		log.WithError(err).Error("failed to shutdown server")
	}
}

func shutdownServer(srv *http.Server, maximumTime time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), maximumTime)
	defer cancel()
	return srv.Shutdown(ctx)
}
