/**
 * Copyright 2024 (c) Paul DESPLANQUES
 *
 * This file is part of shifumi_backend.
 * SPECIFIC NOTE FOR HANDLER:
	 * This handler is for /internal/game/finish endpoint.
	 * It's an handler without session management
	 * It's a POST method

	 * The handler func FinishHandler is used to finish a game with a pre-existing party code. (Save the game result)
*/

package game

import (
	"net/http"
	"fmt"
	"shifumi_backend/src/utils"
)

func FinishHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.IsValidRequest(r, "POST") {
		utils.ErrorResponseHandler(w, "Method not allowed")
		return
	}

	codeParty := r.FormValue("code")
	resultParty := r.FormValue("result")

	if (codeParty == "" || resultParty == "") {
		utils.ErrorResponseHandler(w, "Missing code or result")
		return
	}

	updatePartyInfos(codeParty, resultParty)
	// TODO: update coins of user
	utils.SuccessResponseHandler(w, 0)
}

func updatePartyInfos(code string, status int32) {
	// private func to update the party status

	if utils.DatabaseConnnection == nil {
		return 
	}

	rows, err := utils.DatabaseConnnection.Query("UPDATE shifumi_parties SET status = ? WHERE status = 0 and code = ?", status, code)
	if err != nil {
		fmt.Println(err)
		return 
	}
}
