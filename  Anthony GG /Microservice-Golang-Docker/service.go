package main

import (
	"context"
	"fmt"
)

// PriceFetcher in an interface that can fetch Price
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implements PriceFetcher interface
type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	// business logic here
	// in this example we will use mock data instead of coinmarket
	// business logic have to be clean, do not use json representation
	return MockPriceFetcher(ctx, ticker)

}

var priceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
	"GG":  100_000.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	// mimic this behavior of api call
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) not supported", ticker)
	}

	return price, nil

}
