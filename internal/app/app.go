package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/shafaalafghany/segokuning-social-app/config"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/utils/validation"
	handler "github.com/shafaalafghany/segokuning-social-app/internal/handler/user"
	"github.com/shafaalafghany/segokuning-social-app/internal/repository"
	"github.com/shafaalafghany/segokuning-social-app/pkg/db"
)

func Run(cfg *config.Configuration) {
	var validate *validator.Validate

	pgx := db.NewPsqlDB(cfg)

	ur := repository.NewUserRepo(pgx)

	validate = validator.New()
	if err := validation.RegisterCustomValidation(validate); err != nil {
		log.Fatalf("error register custom validation")
	}

	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		handler.NewUserHandler(r, ur, validate, *cfg)
	})

	s := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: r,
	}
	go func() {
		fmt.Println("Listen and Serve at port 8080")
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("error in ListenAndServe: %s", err)
		}
	}()
	log.Print("Server Started")

	stopped := make(chan os.Signal, 1)
	signal.Notify(stopped, os.Interrupt)
	<-stopped

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("shutting down gracefully...")
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("error in Server Shutdown: %s", err)
	}
	fmt.Println("server stopped")
}
