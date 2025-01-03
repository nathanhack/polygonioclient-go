package models_test

import (
	"testing"

	"cloud.google.com/go/civil"
	"github.com/polygon-io/client-go/rest/models"
)

func TestGetOptionsContractParams(t *testing.T) {
	ticker := "A"
	date := civil.Date{Year: 2023, Month: 3, Day: 23}
	expect := models.GetOptionsContractParams{
		Ticker: ticker,
		AsOf:   &date,
	}
	actual := models.GetOptionsContractParams{
		Ticker: ticker,
	}.WithAsOf(date)
	checkParams(t, expect, *actual)
}

func TestListOptionsContractsParams(t *testing.T) {
	date := civil.Date{Year: 2023, Month: 3, Day: 23}
	contractType := "call"
	expired := true
	strike := 100.0
	sort := models.TickerSymbol
	order := models.Asc
	limit := 100
	ticker := "A"
	expect := models.ListOptionsContractsParams{
		ContractType:        &contractType,
		UnderlyingTickerEQ:  &ticker,
		UnderlyingTickerLT:  &ticker,
		UnderlyingTickerLTE: &ticker,
		UnderlyingTickerGT:  &ticker,
		UnderlyingTickerGTE: &ticker,
		ExpirationDateEQ:    &date,
		ExpirationDateLT:    &date,
		ExpirationDateLTE:   &date,
		ExpirationDateGT:    &date,
		ExpirationDateGTE:   &date,
		StrikePriceEQ:       &strike,
		StrikePriceLT:       &strike,
		StrikePriceLTE:      &strike,
		StrikePriceGT:       &strike,
		StrikePriceGTE:      &strike,
		AsOf:                &date,
		Expired:             &expired,
		Sort:                &sort,
		Order:               &order,
		Limit:               &limit,
	}
	actual := models.ListOptionsContractsParams{}.
		WithContractType(contractType).
		WithUnderlyingTicker(models.EQ, ticker).
		WithUnderlyingTicker(models.LT, ticker).
		WithUnderlyingTicker(models.LTE, ticker).
		WithUnderlyingTicker(models.GT, ticker).
		WithUnderlyingTicker(models.GTE, ticker).
		WithExpirationDate(models.EQ, date).
		WithExpirationDate(models.LT, date).
		WithExpirationDate(models.LTE, date).
		WithExpirationDate(models.GT, date).
		WithExpirationDate(models.GTE, date).
		WithStrikePrice(models.EQ, strike).
		WithStrikePrice(models.LT, strike).
		WithStrikePrice(models.LTE, strike).
		WithStrikePrice(models.GT, strike).
		WithStrikePrice(models.GTE, strike).
		WithAsOf(date).
		WithExpired(expired).
		WithSort(sort).
		WithOrder(order).
		WithLimit(limit)

	checkParams(t, expect, *actual)
}
