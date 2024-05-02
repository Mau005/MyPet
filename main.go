package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mau005/MyPet/configuration"
	"github.com/Mau005/MyPet/db"
	"github.com/Mau005/MyPet/router"
)

func main() {
	if err := configuration.LoadConfiguration("config.yml"); err != nil {
		log.Fatal(err)
	}
	if err := db.ConnectionDataBase(); err != nil {
		log.Panic(err)
	}

	log.Println("Listening Server Run")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", configuration.Config.Server.Ip, configuration.Config.Server.Port), router.NewRouter()))
}
