package middleware

import (
    "net/http"
    "strings"
    "gopanel/internal/auth"
)

func RequireAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if !strings.HasPrefix(authHeader, "Bearer ") {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        token := strings.TrimPrefix(authHeader, "Bearer ")
        user, err := auth.ValidateToken(token)
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Optionally: Store user info in context
        ctx := r.Context()
        ctx = auth.WithUser(ctx, user)
        r = r.WithContext(ctx)

        next.ServeHTTP(w, r)
    })
}
