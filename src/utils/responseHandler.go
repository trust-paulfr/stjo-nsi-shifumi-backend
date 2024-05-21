/**
 * Copyright 2024 (c) Paul DESPLANQUES
 *
 * This file is part of shifumi_backend.
 * This file contains the response handler for the API.
 * Use for handle an uniform response for the API. (Bool, Message, Details)
*/

package utils

import (
	"encoding/json"
	"net/http"
)

type CustomResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func ErrorResponseHandler(w http.ResponseWriter, detailsInfos string) {
	errorReponse := CustomResponse{
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

func SuccessResponseHandler(w http.ResponseWriter, detailsInfos string) {
	CustomResponse := CustomResponse{
		Status:  true,
		Message: "Success",
		Details: detailsInfos,
	}

	jsonResponse, err := json.Marshal(CustomResponse)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
