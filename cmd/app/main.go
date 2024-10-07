package main

import (
	"context"
	"errors"
	"github.com/Pasca11/justNotes/internal/config"
	"github.com/Pasca11/justNotes/internal/repository/postgres"
	"github.com/Pasca11/justNotes/internal/service"
	"github.com/Pasca11/justNotes/internal/transport/controller"
	"github.com/Pasca11/justNotes/internal/transport/router"
	"github.com/Pasca11/justNotes/internal/transport/server"
	"github.com/Pasca11/justNotes/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/Pasca11/justNotes/docs"
)

// @title WeatherApi
// @version 1.0
// @description Sample weather app

// @host :8080
// @basePath /
func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	lg := logger.New(cfg.Logger)
	lg.Debug("Logger initialized")

	db, err := postgres.NewDatabase(cfg.DB)
	if err != nil {
		lg.Error("Failed to connect to database:", err.Error())
		return
	}
	lg.Debug("Database initialized")

	notesService := service.NewNotesService(db)
	userService := service.NewGRPCUserService("auth:8081")

	c, err := controller.New(notesService, userService, lg)
	if err != nil {
		lg.Error("Failed to create controller: " + err.Error())
		return
	}
	lg.Debug("Controller initialized")

	r := router.New(c)
	srv := server.New(cfg.Server, r)
	lg.Debug("Server initialized")

	go func() {
		lg.Info("Server started on port " + cfg.Server.Port)
		err := srv.Run()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			lg.Error(err.Error())
		}
		lg.Info("server stopped")
	}()

	waitForShutdown()
	err = srv.Shutdown(context.Background())
	if err != nil {
		lg.Error("error shutting down server: ", err.Error())
	}
}

func waitForShutdown() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
