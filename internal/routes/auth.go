package routes

import (
	"encoding/json"
	"net/http"

	"github.com/dragonbytelabs/webgamedev/internal/dbx"

	"golang.org/x/crypto/bcrypt"
)

func RegisterAuth(mux *http.ServeMux, db *dbx.DB) {
	mux.HandleFunc("POST /api/register", func(w http.ResponseWriter, r *http.Request) {
		var in struct {
			Email         string `json:"email"`
			Password      string `json:"password"`
			CheckPassword string `json:"checkPassword"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}
		if in.Password != in.CheckPassword {
			http.Error(w, "passwords do not match", http.StatusBadRequest)
			return
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		u, err := db.CreateUser(r.Context(), in.Email, string(hash), &in.Email)
		if err != nil {
			http.Error(w, "could not create user", 500)
			return
		}
		_ = json.NewEncoder(w).Encode(u)
	})

	mux.HandleFunc("POST /api/login", func(w http.ResponseWriter, r *http.Request) {
		var in struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}
		u, err := db.GetUserByEmail(r.Context(), in.Email)
		if err != nil || u == nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(in.Password)) != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}
