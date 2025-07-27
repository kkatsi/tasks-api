package middleware

import (
	"context"
	"net/http"
	"rest-api/internal/apperrors"
	"rest-api/internal/config"
	"rest-api/internal/utils"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeaderValue := r.Header.Get("Authorization")

		if authHeaderValue == "" {
			utils.ErrorResponse(w, http.StatusUnauthorized, apperrors.ErrUnauthorized.Error())
			return
		}

		authTokenValue := strings.TrimPrefix(authHeaderValue, "Bearer ")

		if authHeaderValue == authTokenValue {
			utils.ErrorResponse(w, http.StatusUnauthorized, apperrors.ErrUnauthorized.Error())
			return
		}

		token, err := jwt.Parse(authTokenValue, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, apperrors.ErrUnauthorized
			}

			return []byte(config.Get().AccessSecret), nil
		})

		if err != nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			utils.ErrorResponse(w, http.StatusUnauthorized, apperrors.ErrUnauthorized.Error())
			return
		}

		userId := claims["userId"].(string)
		ctx := context.WithValue(r.Context(), "userId", userId)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
