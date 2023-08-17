package middleware

import (
	"database/sql"
	"errors"
	"lib-manager/pkg/types"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type key int

const (
	admin    key    = iota
	userID   key    = iota
	userName string = ""
)

// Validating Cookies
func Middleware(writer http.ResponseWriter, request *http.Request, db *sql.DB) (int, string, int, error) {
	var sessionInfo types.ValidateCookie

	cookieId := request.Header.Get("Cookie")
	if len(cookieId) < 10 || cookieId == "SessionID=0000000000000000000000000000000000000000" { //NO Cookies ON User Side
		return 0, "", 0, errors.New("Cookie Not Set")
	} else {
		cookieId = request.Header.Get("Cookie")[10:]
	}

	rows, err := db.Query("SELECT sessionId, userId FROM cookie")
	if err != nil {
		return 0, "", 0, errors.New("Nothing in sessionID")
	}
	for rows.Next() {
		rows.Scan(&sessionInfo.SessionID, &sessionInfo.UserId)
	}

	rows, err = db.Query("SELECT  admin, userName FROM users WHERE id = ?", sessionInfo.UserId)
	if err != nil {
		return 0, "", 0, errors.New("Nothing in sessionID")
	}
	for rows.Next() {
		rows.Scan(&sessionInfo.Admin, &sessionInfo.Username)
	}
	if sessionInfo.SessionID == "" || sessionInfo.Username == "" || sessionInfo.UserId == 0 { // Session database table is empty
		return 0, "", 0, errors.New("Nothing in sessionID")
	}

	if cookieId != sessionInfo.SessionID { // Cookie Id has been tempered
		return 0, "", 0, errors.New("Cookie Was Altered On User Side")
	}
	return sessionInfo.UserId, sessionInfo.Username, sessionInfo.Admin, nil
}
