package middlewares

import (
    "context"
    "net/http"
    "strings"

    "github.com/golang-jwt/jwt"
    "github.com/seeduler/seeduler/models"
    "github.com/seeduler/seeduler/services"
)

type contextKey string

const userContextKey contextKey = "user"

func AuthMiddleware(userService *services.UserService) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
                http.Error(w, "Authorization header required", http.StatusUnauthorized)
                return
            }

            tokenString := strings.TrimPrefix(authHeader, "Bearer ")
            claims := &jwt.MapClaims{}
            token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
                return userService.JWTKey, nil
            })

            if err != nil || !token.Valid {
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            user := &models.User{
                Username: (*claims)["username"].(string),
                HallID:   int((*claims)["hall_id"].(float64)),
            }

            ctx := context.WithValue(r.Context(), userContextKey, user)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

func GetUserFromContext(ctx context.Context) *models.User {
    user, ok := ctx.Value(userContextKey).(*models.User)
    if !ok {
        return nil
    }
    return user
}