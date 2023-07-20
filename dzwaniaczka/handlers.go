package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

func CreateFolderLogs() {
	// println("test test hello")
	path := "logs"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
func SetEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("ERROR: loading .env file: ", err)
	}
	log.Println("INFO: env successfully loaded")

}

func AlertaLoop() {
	ticker := time.NewTicker(2 * time.Second)
	ALERTATOKEN, ok := viper.Get("AlertaToken").(string)
	if !ok {
		log.Fatalf("Invalid type token")
	}

	var c = NewServer(ALERTATOKEN) //client to work with ALERTA api

	for {
		select {
		case <-ticker.C:
			fmt.Println("toock")

			result, err := c.SearchAllerts() // alertaAlerts.go

			if err != nil {
				log.Println("ERROR: c.SearchAllerts(): ", err)
			}

			// fmt.Println("SearchAllerts: ", result)
			fmt.Println("SearchAllerts Status: ", result.Status)
			fmt.Println("SearchAllerts len Alerts: ", len(result.Alerts))
			fmt.Println("SearchAllerts []Alerts[0].Event: ", result.Alerts[0].Event)
			fmt.Println("SearchAllerts []Alerts[0].Service[0]: ", result.Alerts[0].Service[0])
			fmt.Println("SearchAllerts []Alerts[0].Enviroment: ", result.Alerts[0].Environment)
			fmt.Println("SearchAllerts []Alerts[0].Resource: ", result.Alerts[0].Resource)
		}
	}
}
