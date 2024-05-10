package utils

import "net/http"

func isValidRequest(r *http.Request, requestMethod string) bool {
	return r.Method == requestMethod
}

func AreValidArgsRequest(r *http.Request, requestArgs []string) bool {
	for _, arg := range requestArgs {
		if r.FormValue(arg) == "" {
			return false
		}
	}
	return true
}
