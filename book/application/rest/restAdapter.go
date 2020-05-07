package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/BackAged/library-management/book/configuration"
	"github.com/BackAged/library-management/book/domain/book"
	"github.com/BackAged/library-management/book/infrastructure/database"
	"github.com/BackAged/library-management/book/infrastructure/repository"
	"github.com/go-chi/chi"
)

// Serve serves rest api
func Serve(cfgPath string) error {
	cfg, err := configuration.Load(cfgPath)
	if err != nil {
		return err
	}

	fmt.Println(cfg)

	rds, err := database.NewClient(cfg.Mongo.URI, cfg.Mongo.Database)
	if err != nil {
		fmt.Println(err)
		return err
	}

	bkRepo := repository.NewBookRepository(rds, "books")
	bkSvc := book.NewService(bkRepo)
	bkHndlr := NewBookHandler(bkSvc)

	r := chi.NewRouter()
	r.Mount("/api/v1/book", NewBookRouter(bkHndlr))

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		Handler:      r,
	}

	go func() {
		log.Println("Staring server with address ", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Failed to listen and serve", err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.GracefulTimeout)*time.Second)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
	return nil
}
