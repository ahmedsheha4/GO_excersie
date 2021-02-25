package main

import (
	"fmt"
	"log"
	"net/http"

	handleapis "github.com/ahmedsheha4/GO_excersie/HTTP"
	businesslogic "github.com/ahmedsheha4/GO_excersie/businessLogic"
	repository "github.com/ahmedsheha4/GO_excersie/repository/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	repository, err := repository.SetupStorage()
	if err != nil {
		log.Fatalln("An error happened...", err)
	}
	service := businesslogic.NewService(repository)
	fmt.Println("Server listening on port 8000")
	router := handleapis.InitHandlers(service)
	log.Fatal(http.ListenAndServe(":8000", router))
}
