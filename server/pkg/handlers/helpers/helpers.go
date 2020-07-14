package helpers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// GetintKeyFromRequestParam returns the param value from request
func GetintKeyFromRequestParam(r *http.Request, param string) int {
	keyStr := chi.URLParam(r, param)
	keyInt, _ := strconv.ParseInt(keyStr, 10, 64)

	return int(keyInt)
}

// GetKeyFromRequestParam returns the param value from request
func GetKeyFromRequestParam(r *http.Request, param string) string {
	keyStr := chi.URLParam(r, param)

	return keyStr
}
