package main

import (
	"crypto/sha256"
	"net/http"
)

type AuthHandler struct {
	Username string
	Password string
}

func (auth *AuthHandler) auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if ok {
			uhash := sha256.Sum256([]byte(username))
			phash := sha256.Sum256([]byte(password))
			expectuhash := sha256.Sum256([]byte(auth.Username))
			expectphash := sha256.Sum256([]byte(auth.Password))
			if uhash == expectuhash && phash == expectphash {
				next.ServeHTTP(w, r)
				return
			}
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
