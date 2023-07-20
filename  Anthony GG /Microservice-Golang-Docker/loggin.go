package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next PriceFetcher
}

func NewLoggingService(next PriceFetcher) PriceFetcher {
	// we are not returning looging, we returning PriceFetcher
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		// logrus is for logginh
		logrus.WithFields(logrus.Fields{
			"took":  time.Since(begin),
			"err":   err,
			"price": price,
		}).Info("fetchPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}
