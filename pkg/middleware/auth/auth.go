package auth

import (
	"context"
	"net/http"
)

type contextKey string // int?
var tokenContextKey = contextKey("token")

func Auth() func(
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			token := request.Header.Get("Authorization")
			if token == "" {
				next(writer, request)
				return
			}

			ctx := context.WithValue(
				request.Context(),
				tokenContextKey,
				token,
			)
			next(writer, request.WithContext(ctx))
		}
	}
}

func FromContext(ctx context.Context) (token string, ok bool) {
	token, ok = ctx.Value(tokenContextKey).(string)
	return
}
