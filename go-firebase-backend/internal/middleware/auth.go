package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/mik-sea/go-firebase-backend/internal/config"
)

func FirebaseAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		ctx := context.Background()
		authClient, err := config.App.Auth(ctx)
		if err != nil {
			http.Error(w, "Auth client error", 500)
			return
		}

		token, err := authClient.VerifyIDToken(ctx, idToken)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Simpan UID ke context
		ctx = context.WithValue(r.Context(), "uid", token.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
