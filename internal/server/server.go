package server

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Mux      *http.ServeMux
	Server   *http.Server
	Handlers *Handlers
}

func NewServer() Server {
	mux := http.NewServeMux()
	handlers := NewHandlers()
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	return Server{
		Mux: mux,
		Server: &http.Server{
			Addr:    ":" + port,
			Handler: mux,
		},
		Handlers: handlers,
	}
}

func (s Server) RegisterHandlers() {
	s.Mux.HandleFunc("GET /ping", LoggerMiddleware(s.Handlers.PingHandler))
	s.Mux.HandleFunc("GET /bars", LoggerMiddleware(s.Handlers.BarsHandler))
}

func (s Server) MustRun() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		slog.Info("starting http server", "port", s.Server.Addr)
		if err := s.Server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	<-stop
	slog.Info("server shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Server.Shutdown(ctx); err != nil {
		slog.Error("couldn't gracefully shut down the server")
	}

	slog.Info("server has been gracefully shut down")
}
