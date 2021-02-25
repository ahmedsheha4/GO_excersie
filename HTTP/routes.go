package handleapis

import (
	businesslogic "github.com/ahmedsheha4/GO_excersie/businessLogic"
	"github.com/gorilla/mux"
)

//InitHandlers for routes
func InitHandlers(service businesslogic.Service) *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/phones", GetPhonesController(service)).Methods("GET")

	return router
}
