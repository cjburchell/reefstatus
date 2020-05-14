package token

import (
	"net/http"
)

func Middleware(next http.Handler, dataServiceToken string) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		auth := request.Header.Get("Authorization")
		if auth != "APIKEY "+dataServiceToken {
			response.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(response, request)
	})
}
