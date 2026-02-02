package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	client, err := pgx.Connect(context.TODO(), "postgres://postgres@localhost:5432/challenges")
	if err != nil {
		fmt.Println("unable to connect to DB")
		panic(err)
	}
	defer client.Close(context.TODO())
	Logger := slog.Default()

	mux := http.NewServeMux()
	HandleNewUser := http.HandlerFunc(NewUserFunc(client))
	mux.Handle("/register", AddLogging(HandleNewUser, Logger))
	mux.Handle("/auth", AuthUserFunc(client, AddLogging(GetUserFunc(client), Logger)))
	http.ListenAndServe("localhost:8080", mux)
}

func ExtractUser(body *io.ReadCloser) (payload User, err error) {
	data, _ := io.ReadAll(*body)
	err = json.Unmarshal(data, &payload)
	if err != nil {
		log.Default().Println("Failed to Extract User Data")
	}
	return payload, err
}

func NewUserFunc(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer db.Close(r.Context())
		user, err := ExtractUser(&r.Body)
		if err != nil {
			return
		}
		query := "INSERT into users (username, email, password) VALUES (@User, @Email,@Password)"
		args := pgx.NamedArgs{
			"User":     user.Name,
			"Email":    user.Email,
			"Password": GetMD5Hash(user.Password),
		}
		_, err = db.Exec(r.Context(), query, args)
		if err != nil {
			slog.Default().Error("Unable to add user", "Postgres Error", err)

		}
	}
}

func checkPassword(password string, expectedPassword string) bool {
	passwordHash := GetMD5Hash(password)
	if passwordHash == expectedPassword {
		return true
	}

	return false
}

func AuthUserFunc(db *pgx.Conn, handeler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		username, password, ok := r.BasicAuth()

		var expectedUsername string
		var expectedPassword string

		if ok {
			query := fmt.Sprintf("select username,password from users where username = '%s'", username)
			resultRows := db.QueryRow(r.Context(), query)
			err := resultRows.Scan(&expectedUsername, &expectedPassword)
			if err != nil {
				slog.Warn(err.Error())

			}
		}

		if checkPassword(password, expectedPassword) {
			handeler.ServeHTTP(w, r)
			return
		}

		http.Error(w, "Authentication Failed", http.StatusBadRequest)
	})
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetUserFunc(db *pgx.Conn) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.ReadAll(r.Body)
	})
}

type User struct {
	Name      string `json:"Username"`
	ID        string `json:"ID,omitempty"`
	Email     string `json:"Email"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created,omitempty"`
}

func ListUsersFunc() {

}
func AddLogging(h http.Handler, l *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		timeStart := time.Now()

		l.Debug("Request Started", "time: ", timeStart)
		defer func() { l.Info(r.URL.Path, "Elapsed: ", time.Since(timeStart).String()) }()
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
