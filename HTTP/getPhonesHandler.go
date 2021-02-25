package handleapis

import (
	"encoding/json"
	"net/http"
	"strconv"

	businesslogic "github.com/ahmedsheha4/GO_excersie/businessLogic"
)

//GetPhonesController .
func GetPhonesController(service businesslogic.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var res businesslogic.Response
		countryFilter := r.URL.Query().Get("country")
		filterByState := r.URL.Query().Get("state")

		//Pagination
		page := 1
		limit := 10
		if len(r.URL.Query().Get("limit")) > 0 {
			limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
		}
		if len(r.URL.Query().Get("page")) > 0 {
			page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		}
		startIndex := (page - 1) * limit
		endIndex := page * limit

		//Get phones
		phones, err := service.GetPhones(countryFilter, filterByState)
		if err != nil {
			http.Error(w, "An error happened..", http.StatusInternalServerError)
			return
		}

		if endIndex > len(phones) || endIndex < 0 {
			endIndex = len(phones)
		}
		if startIndex < 0 || startIndex > len(phones) {
			startIndex = 0
		}
		res.Phones = phones[startIndex:endIndex]
		res.TotalRows = len(phones)

		(w).Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
