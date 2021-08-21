package middlewares

import (
	"fmt"
	"net/http"
)

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("middleware authentication if")
		next(w, r)
	}
}
