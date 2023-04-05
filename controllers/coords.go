package controllers

import (
	"awesomeProject/db"
	"awesomeProject/encoder"
	"awesomeProject/server"
	"encoding/json"
	"errors"
	"net/http"
)

type requestData struct {
	Id  string `json:"id"`
	Geo geo    `json:"geo"`
}

type geo struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	Hdn float64 `json:"hdn"`
	Acu float64 `json:"acu"`
	Spd float64 `json:"spd"`
	Stp int64   `json:"stp"`
}

type CC struct {
	id int
}

func SaveCoords(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	token := r.Header.Get("Authorization")

	courier := server.GetCourierByToken(token, w)

	if courier.Item_id <= 0 {
		server.ErrorHandler(w, errors.New("Courier not found"))
		return
	}

	data := json.NewDecoder(r.Body)

	d := &requestData{}

	err := data.Decode(d)

	if err != nil {
		server.ErrorHandler(w, err)
		return
	}

	request := d

	id := encoder.Decode(string(request.Id))

	db := db.Connect()

	rows, err := db.Query("SELECT id FROM courier_receive WHERE id = ?", id)

	if err != nil {
		server.ErrorHandler(w, err)
		return
	}

	defer rows.Close()

	courierReceives := CC{}

	rows.Next()
	err = rows.Scan(&courierReceives.id)

	if err != nil {
		server.ErrorHandler(w, err)
		return
	}

	result, err := db.Exec("INSERT INTO geolog(courier_id, event, cr_id, lat, lng, hdn, acu, spd, stp) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		courier.Item_id,
		"track",
		courierReceives.id,
		request.Geo.Lat,
		request.Geo.Lng,
		request.Geo.Hdn,
		request.Geo.Acu,
		request.Geo.Spd,
		request.Geo.Stp)

	defer db.Close()

	if err != nil {
		server.ErrorHandler(w, err)
		return
	} else {
		lastId, err := result.LastInsertId()
		if lastId > 0 {
			var status = &server.Success{
				Success: true,
			}

			j, _ := json.Marshal(status)
			w.Write(j)
			return
		} else {
			server.ErrorHandler(w, err)
			return
		}
	}
}
