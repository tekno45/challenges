package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	client, err := pgx.Connect(context.TODO(), "postgres://postgres@localhost:5432")
	if err != nil {
		fmt.Println("unable to connect to DB")
		panic(err)
	}
	defer client.Close(context.TODO())
	l := Log{logger: log.Default()}

	mux := http.NewServeMux()
	hu := http.HandlerFunc(HandleNewUser)
	mux.Handle("/register", l.Log(hu))
	http.ListenAndServe("localhost:8080", mux)
}

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "registering")
}

type User struct {
	Name     string `json:"Name"`
	ID       string `json:"ID,omitempty"`
	Email    string `json:"Email"`
	Password string `json:"password,omitempty"`
}

type Log struct {
	logger *log.Logger
}

func (l *Log) Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		timeStart := time.Now()
		l.logger.Println("Method: ", r.Method)
		l.logger.Println("Path: ", r.URL)
		defer func() { l.logger.Println("Took: ", time.Since(timeStart)) }()
		h.ServeHTTP(w, r)
	})
}

func RequireAdmin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("auth")
		if token != "admin" {
			http.Error(w, "Admin Permissions Required", http.StatusUnauthorized)
		}
		h.ServeHTTP(w, r)

	})
}
