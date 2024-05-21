/**
 * Copyright 2024 (c) Paul DESPLANQUES
 *
 * This file is part of shifumi_backend.
 * It's a group of utils functions for the global project.
*/

package utils

import "net/http"

func IsValidRequest(r *http.Request, requestMethod string) bool {
	return r.Method == requestMethod
}