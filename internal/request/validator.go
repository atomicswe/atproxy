package request

import (
	"net/http"
	"strings"
)

func Allowed(r *http.Request) bool {
	if strings.Contains(r.URL.Hostname(), "example.com") {
		return false
	}
	return true
}
