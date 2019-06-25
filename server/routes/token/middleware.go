package token

import (
	"net/http"

	"github.com/cjburchell/reefstatus/server/settings"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		auth := request.Header.Get("Authorization")
		if auth != "APIKEY "+settings.DataServiceToken {
			response.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(response, request)
	})
}
