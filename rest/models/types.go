package models

import (
	"encoding/json"
	"strconv"
	"time"
)

// MarketType is the type of market.
type MarketType string

const (
	Stocks MarketType = "stocks"
	Forex  MarketType = "forex"
	Crypto MarketType = "crypto"
)

// ContractType is the type of contract.
type ContractType string

const (
	ContractCall  ContractType = "call"
	ContractPut   ContractType = "put"
	ContractOther ContractType = "other"
)

// Locale is the market location.
type MarketLocale string

const (
	US     MarketLocale = "us"
	Global MarketLocale = "global"
)

// Timespan is the size of the time window.
type Timespan string

const (
	Second  Timespan = "second"
	Minute  Timespan = "minute"
	Hour    Timespan = "hour"
	Day     Timespan = "day"
	Week    Timespan = "week"
	Month   Timespan = "month"
	Quarter Timespan = "quarter"
	Year    Timespan = "year"
)

// Sort is a query param type that specifies how the results should be sorted.
type Sort string

const (
	TickerSymbol       Sort = "ticker"
	Name               Sort = "name"
	Market             Sort = "market"
	Locale             Sort = "locale"
	PrimaryExchange    Sort = "primary_exchange"
	Type               Sort = "type"
	CurrencySymbol     Sort = "currency_symbol"
	CurrencyName       Sort = "currency_name"
	BaseCurrencySymbol Sort = "base_currency_symbol"
	BaseCurrencyName   Sort = "base_currency_name"
	CIK                Sort = "cik"
	CompositeFIGI      Sort = "composite_figi"
	ShareClassFIGI     Sort = "share_class_figi"
	PublishedUTC       Sort = "published_utc"
	LastUpdatedUTC     Sort = "last_updated_utc"
	DelistedUTC        Sort = "delisted_utc"
	Timestamp          Sort = "timestamp"
	StrikePrice        Sort = "strike_price"
	ExpirationDate     Sort = "expiration_date"
	FilingDate         Sort = "filing_date"
	PeriodOfReportDate Sort = "period_of_report_date"
)

// Order the results. asc will return results in ascending order (oldest at the top), desc will return results in
// descending order (newest at the top).
type Order string

const (
	Asc  Order = "asc"
	Desc Order = "desc"
)

type SeriesType string

const (
	High  SeriesType = "high"
	Open  SeriesType = "open"
	Low   SeriesType = "low"
	Close SeriesType = "close"
)

// Direction is the direction of the snapshot results to return.
type Direction string

const (
	Gainers Direction = "gainers"
	Losers  Direction = "losers"
)

// AssetClass is an identifier for a group of similar financial instruments.
type AssetClass string

const (
	AssetStocks  AssetClass = "stocks"
	AssetOptions AssetClass = "options"
	AssetCrypto  AssetClass = "crypto"
	AssetFx      AssetClass = "fx"
	AssetOTC     AssetClass = "otc"
	AssetIndices AssetClass = "indices"
)

// DataType is the type of data.
type DataType string

const (
	DataTrade DataType = "trade"
	DataBBO   DataType = "bbo"
	DataNBBO  DataType = "nbbo"
)

// SIP is the type of Securies Information Processor.
type SIP string

const (
	CTA  SIP = "CTA"
	UTP  SIP = "UTP"
	OPRA SIP = "OPRA"
)

// Frequency is the number of times a dividend is paid out over the course of one year.
type Frequency int64

const (
	OneTime    Frequency = 0
	Annually   Frequency = 1
	BiAnnually Frequency = 2
	Quarterly  Frequency = 4
	Monthly    Frequency = 12
)

// DividendType is the type of dividend.
type DividendType string

const (
	DividendCD DividendType = "CD"
	DividendLT DividendType = "LT"
	DividendSC DividendType = "SC"
	DividendST DividendType = "ST"
)

// Comparator is the type of comparison to make for a specific query parameter.
type Comparator string

const (
	EQ  Comparator = "eq"
	LT  Comparator = "lt"
	LTE Comparator = "lte"
	GT  Comparator = "gt"
	GTE Comparator = "gte"
)

// NameComparator is the type of comparison to make for the company_name query parameter in Stock Financials.
type NameComparator string

const (
	Full   NameComparator = "full"
	Search NameComparator = "search"
)

// TimeFrame is the type of time frame query parameter for stock financials.
type Timeframe string

const (
	TFAnnual    Timeframe = "annual"
	TFQuarterly Timeframe = "quarterly"
)

// Time represents a long date string of the following format: "2006-01-02T15:04:05.000Z".
type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	unquoteData, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	// attempt to parse time
	if parsedTime, err := time.Parse("2006-01-02T15:04:05.000-0700", unquoteData); err == nil {
		*t = Time(parsedTime)
		return nil
	}

	// attempt to parse time again
	if parsedTime, err := time.Parse("2006-01-02T15:04:05-07:00", unquoteData); err == nil {
		*t = Time(parsedTime)
		return nil
	}

	// attempt with a different format
	if parsedTime, err := time.Parse("2006-01-02T15:04:05.000Z", unquoteData); err == nil {
		*t = Time(parsedTime)
		return nil
	}

	// attempt with yet another format
	if parsedTime, err := time.Parse("2006-01-02T15:04:05Z", unquoteData); err != nil {
		return err
	} else {
		*t = Time(parsedTime)
	}

	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*t).Format("2006-01-02T15:04:05.000Z"))
}


// Millis represents a Unix time in milliseconds since January 1, 1970 UTC.
type Millis time.Time

func (m *Millis) UnmarshalJSON(data []byte) error {
	d, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*m = Millis(time.UnixMilli(d))
	return nil
}

func (m Millis) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(m).UnixMilli())
}

// Nanos represents a Unix time in nanoseconds since January 1, 1970 UTC.
type Nanos time.Time

func (n *Nanos) UnmarshalJSON(data []byte) error {
	d, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	// Go Time package does not include a method to convert UnixNano to a time.
	timeNano := time.Unix(d/1_000_000_000, d%1_000_000_000)
	*n = Nanos(timeNano)
	return nil
}

func (n Nanos) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(n).UnixNano())
}
