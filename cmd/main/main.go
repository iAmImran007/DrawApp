package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iAmImran007/draw-app-js-go/pkg/config"
	"github.com/iAmImran007/draw-app-js-go/pkg/models"
	"github.com/iAmImran007/draw-app-js-go/pkg/routes"
)


func main() {

	//load the env fie 
	config.LoadEnv()

	//conect to db
	config.ConectDb()

	//run migration
	config.DB.AutoMigrate(&models.User{}, &models.Drawing{})

	//create a router
	router := mux.NewRouter()

	//register routes
	routes.RegisterApproutes(router)
  

	//start a server 
	fmt.Println("server runing on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

}