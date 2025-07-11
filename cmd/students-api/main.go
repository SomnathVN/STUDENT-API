package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SomnathVN/students-api/internal/config"
	"github.com/SomnathVN/students-api/internal/http/handlers/student"
	//"github.com/SomnathVN/students-api/internal/storage/sqlite"
	"github.com/SomnathVN/students-api/internal/storage/firestore"
	"github.com/SomnathVN/students-api/internal/http/middleware"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//database setup

	// storage, err := sqlite.New(cfg)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	storage, err := firestore.New(cfg)
	if err != nil {
		log.Fatal(err)
	}


	slog.Info("storge initialized", slog.String("env", cfg.Env), slog.String("version","1.0.0"))


	//setup router
	router := http.NewServeMux()


	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))
	router.HandleFunc("PUT /api/students/{id}", student.Update(storage))
	router.HandleFunc("DELETE /api/students/{id}", student.Delete(storage))

	//api key
	apiKey := cfg.APIKey

    handler := middleware.Logging(
        middleware.RateLimit(
            middleware.APIKeyAuth(apiKey)(router),
        ),
    )

	//setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: handler,
	}

	slog.Info("server started", slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start the server")
		}
	}()

	<-done

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown the server", slog.String("error", err.Error()))
	}
	// err = server.Shutdown(ctx)
	// if err != nil {
	// 	slog.Error("failed to shutdown the server", slog.String("error", err.Error()))
	// }

	slog.Info("server shutdown succesfuly")

}
