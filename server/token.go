package server

import (
	db2 "awesomeProject/db"
	"net/http"
)

type Token struct {
	Id      int
	Item_id int
}

func GetCourierByToken(token string, w http.ResponseWriter) Token {
	db := db2.Connect()

	rows, err := db.Query("SELECT id, item_id FROM tokens WHERE token = ? AND item_model = ?", token, "Couriers")

	if err != nil {
		ErrorHandler(w, err)
	}

	itemToken := Token{}

	rows.Next()
	err = rows.Scan(&itemToken.Id, &itemToken.Item_id)

	/*if err != nil {
		ErrorHandler(w, err)
	}*/

	defer rows.Close()
	defer db.Close()

	return itemToken
}
