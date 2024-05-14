package handlers

import (
	"net/http"
	"strings"

	"diplom-backend/internal/common/auth"
	"diplom-backend/internal/common/errors"
	"diplom-backend/internal/domain"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/auth/sign-in" ||
			r.URL.Path == "/api/v1/auth/sign-up" ||
			r.URL.Path == "/api/v1/courses" && r.Method == "POST" ||
			strings.HasPrefix(r.URL.Path, "/docs") && r.Method == "GET" ||
			strings.HasPrefix(r.URL.Path, "/static") && r.Method == "GET" {
			next.ServeHTTP(w, r)
			return
		}

		headerValue := r.Header.Get("Authorization")
		if len(headerValue) < 7 || strings.ToLower(headerValue[0:6]) != "bearer" {
			ErrorResponse(w, r, errors.NewAuthError("token is not presented", "token"))
			return
		}

		bearerToken := headerValue[7:]
		if bearerToken == "" {
			ErrorResponse(w, r, errors.NewAuthError("token is empty", "token"))
			return
		}

		claims, err := auth.ValidateJWT(bearerToken)
		if err != nil {
			ErrorResponse(w, r, errors.NewAuthError("invalid token", "token"))
			return
		}

		ac := &domain.AuthContext{
			ID:   claims.ID,
			Type: claims.Type,
			Name: claims.Name,
		}

		ctx := auth.SetAuthContext(r.Context(), ac)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
