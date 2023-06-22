package main

import (
	"SatelitePacheco/database"
	"SatelitePacheco/handlers"
	"fmt"
	"log"
)
func errorFunc(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	if !database.ConectionCheck(){
		log.Fatal("Error conectionCheck")
	}
	fmt.Println(database.PostgresDB)
	err := database.PostgresDB.Init()
	errorFunc(err)


	handlers.Handlers()
}
