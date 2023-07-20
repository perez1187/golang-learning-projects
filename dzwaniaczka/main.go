package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

func init() {
	log.Println("INFO: dzwaniaczka init start")

	SetEnv()
}

func main() {

	// userch := make(chan string)

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

	// go AlertaLoop()

	// ticker := time.NewTicker(2 * time.Second)
	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		fmt.Println("tick")
	// 	}
	// }

}
