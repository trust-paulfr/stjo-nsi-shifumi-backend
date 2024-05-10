package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func ErrorResponseHandler(w http.ResponseWriter, detailsInfos string) {
	errorReponse := ErrorResponse{
		Status:  false,
		Message: "An error occurred",
		Details: detailsInfos,
	}
	jsonResponse, err := json.Marshal(errorReponse)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonResponse)
}

func SuccessResponseHandler(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
