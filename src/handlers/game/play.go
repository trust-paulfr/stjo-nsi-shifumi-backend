/**
 * Copyright 2024 (c) Paul DESPLANQUES
 *
 * This file is part of shifumi_backend.
 * SPECIFIC NOTE FOR HANDLER:
	 * This handler is for /internal/game/play endpoint.
	 * It's an handler without session management
	 * It's a GET method

	 * The handler func PlayHandler is used to play a game with a pre-existing party code.
*/

package game

import (
	"net/http"
	"fmt"
	"shifumi_backend/src/utils"
)

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	// The handler func PlayHandler is used to play a game with a pre-existing party code.
	if !utils.IsValidRequest(r, "GET") {
		utils.ErrorResponseHandler(w, "Method not allowed")
		return
	}

	codeParty := r.URL.Query().Get("code")
	if codeParty == "" {
		utils.ErrorResponseHandler(w, "Missing code")
		return
	}

	isValidParty, dataParty := getPartieByCode(codeParty)

	if !isValidParty {
		utils.ErrorResponseHandler(w, "Party not found")
		return
	}

	utils.SuccessResponseHandler(w, dataParty.PrenomUser)
}


type shifumiPartie struct {
	ID int32
	code string
	status int32
	PrenomUser string
}

func getPartieByCode(code string) (bool, shifumiPartie) {
	// private func, return bool (Is Valid Party Code?) and shifumiPartie (Data of the party)

	dataPartie := shifumiPartie{}

	if utils.DatabaseConnnection == nil {
		return false, dataPartie
	}

	rows, err := utils.DatabaseConnnection.Query("SELECT shifumi_parties.ID, code, status, shifumi_users.prenom AS PrenomUser  FROM shifumi_parties INNER JOIN shifumi_users ON shifumi_parties.userID = shifumi_users.ID WHERE code = ?", code)
	if err != nil {
		fmt.Println(err)
		return false, dataPartie
	}

	defer rows.Close()

	var isNotEmptyRow bool

	for rows.Next() {
		var ID int32
		var code string
		var status int32
		var PrenomUser string

		err = rows.Scan(&ID, &code, &status, &PrenomUser)
		if err != nil {
			return false, dataPartie
		}

		isNotEmptyRow = true
		dataPartie = shifumiPartie{
			ID: ID,
			code: code,
			status: status,
			PrenomUser: PrenomUser,
		}
	}

	return isNotEmptyRow, dataPartie
}