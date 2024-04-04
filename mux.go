package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ko44d/go-api-sample/clock"
	"github.com/ko44d/go-api-sample/config"
	"github.com/ko44d/go-api-sample/handler"
	"github.com/ko44d/go-api-sample/service"
	"github.com/ko44d/go-api-sample/store"
	"net/http"
)

func NewMux(ctx context.Context, c *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	v := validator.New()
	db, cleanup, err := store.New(ctx, c)
	if err != nil {
		return nil, cleanup, err
	}
	r := store.Repository{Clocker: clock.RealClocker{}}
	at := &handler.AddTask{
		Service:   &service.AddTask{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/tasks", at.ServeHTTP)
	lt := &handler.ListTask{
		Service: &service.ListTask{DB: db, Repo: &r},
	}
	mux.Get("/tasks", lt.ServeHTTP)
	ru := &handler.RegisterUser{
		Service:   &service.RegisterUser{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/register", ru.ServeHTTP)
	return mux, cleanup, nil
}
