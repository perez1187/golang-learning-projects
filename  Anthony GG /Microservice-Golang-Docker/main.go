package main

import (
	"context"
	"fmt"
	"log"
)

func main() {

	// this is how we check if logging works
	// we create service logging and call method FetchPrice
	//next we create metric service and we add to svc
	svc := NewLoggingService(NewMetricService(&priceFetcher{}))
	price, err := svc.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(price)
}
