package models

import (
	"strings"

	"cloud.google.com/go/civil"
)

// ListTickersParams is the set of parameters for the ListTickers method.
type ListTickersParams struct {
	// Specify a ticker symbol. Defaults to empty string which queries all tickers.
	TickerEQ  *string `query:"ticker"`
	TickerLT  *string `query:"ticker.lt"`
	TickerLTE *string `query:"ticker.lte"`
	TickerGT  *string `query:"ticker.gt"`
	TickerGTE *string `query:"ticker.gte"`

	// Specify the type of the tickers. Find the types that we support via our Ticker Types API. Defaults to empty
	// string which queries all types.
	Type *string `query:"type"`

	// Filter by market type. By default all markets are included.
	Market *AssetClass `query:"market"`

	// Specify the asset's primary exchange Market Identifier Code (MIC) according to ISO 10383. Defaults to empty
	// string which queries all exchanges.
	Exchange *string `query:"exchange"`

	// Specify the CUSIP code of the asset you want to search for. Find more information about CUSIP codes at their
	// website. Defaults to empty string which queries all CUSIPs.
	//
	// Note: Although you can query by CUSIP, due to legal reasons we do not return the CUSIP in the response.
	CUSIP *int `query:"cusip"`

	// Specify the CIK of the asset you want to search for. Find more information about CIK codes at their website.
	// Defaults to empty string which queries all CIKs.
	CIK *int `query:"cik"`

	// Specify a point in time to retrieve tickers available on that date. Defaults to the most recent available date.
	Date *civil.Date `query:"date"`

	// Specify if the tickers returned should be actively traded on the queried date. Default is true.
	Active *bool `query:"active"`

	// Search for terms within the ticker and/or company name.
	Search *string `query:"search"`

	// The field to sort the results on. Default is ticker. If the search query parameter is present, sort is ignored
	// and results are ordered by relevance.
	Sort *Sort `query:"sort"`

	// The order to sort the results on. Default is asc (ascending).
	Order *Order `query:"order"`

	// Limit the size of the response, default is 100 and max is 1000.
	Limit *int `query:"limit"`
}

func (p ListTickersParams) WithTicker(c Comparator, q string) *ListTickersParams {
	switch c {
	case EQ:
		p.TickerEQ = &q
	case LT:
		p.TickerLT = &q
	case LTE:
		p.TickerLTE = &q
	case GT:
		p.TickerGT = &q
	case GTE:
		p.TickerGTE = &q
	}
	return &p
}

func (p ListTickersParams) WithType(q string) *ListTickersParams {
	p.Type = &q
	return &p
}

func (p ListTickersParams) WithMarket(q AssetClass) *ListTickersParams {
	p.Market = &q
	return &p
}

func (p ListTickersParams) WithExchange(q string) *ListTickersParams {
	p.Exchange = &q
	return &p
}

func (p ListTickersParams) WithCUSIP(q int) *ListTickersParams {
	p.CUSIP = &q
	return &p
}

func (p ListTickersParams) WithCIK(q int) *ListTickersParams {
	p.CIK = &q
	return &p
}

func (p ListTickersParams) WithDate(q civil.Date) *ListTickersParams {
	p.Date = &q
	return &p
}

func (p ListTickersParams) WithActive(q bool) *ListTickersParams {
	p.Active = &q
	return &p
}

func (p ListTickersParams) WithSearch(q string) *ListTickersParams {
	p.Search = &q
	return &p
}

func (p ListTickersParams) WithSort(q Sort) *ListTickersParams {
	p.Sort = &q
	return &p
}

func (p ListTickersParams) WithOrder(q Order) *ListTickersParams {
	p.Order = &q
	return &p
}

func (p ListTickersParams) WithLimit(q int) *ListTickersParams {
	p.Limit = &q
	return &p
}

// ListTickersResponse is the response returned by the ListTickers method.
type ListTickersResponse struct {
	BaseResponse

	// An array of tickers that match your query. Note: Although you can query by CUSIP, due to legal reasons we do not
	// return the CUSIP in the response.
	Results []Ticker `json:"results,omitempty"`
}

// GetTickerDetailsParams is the set of parameters for the GetTickerDetails method.
type GetTickerDetailsParams struct {
	// The ticker symbol of the asset.
	Ticker string `validate:"required" path:"ticker"`

	// Specify a point in time to get information about the ticker available on that date. When retrieving information
	// from SEC filings, we compare this date with the period of report date on the SEC filing.
	//
	// For example, consider an SEC filing submitted by AAPL on 2019-07-31, with a period of report date ending on
	// 2019-06-29. That means that the filing was submitted on 2019-07-31, but the filing was created based on
	// information from 2019-06-29. If you were to query for AAPL details on 2019-06-29, the ticker details would
	// include information from the SEC filing.
	//
	// Defaults to the most recent available date.
	Date *civil.Date `query:"date"`
}

func (p GetTickerDetailsParams) WithDate(q civil.Date) *GetTickerDetailsParams {
	p.Date = &q
	return &p
}

// GetTickerDetailsResponse is the response returned by the GetTickerDetails method.
type GetTickerDetailsResponse struct {
	BaseResponse

	// Ticker with details.
	Results Ticker `json:"results,omitempty"`
}

// ListTickerNewsParams is the set of parameters for the ListTickerNews method.
type ListTickerNewsParams struct {
	// Return results that contain this ticker.
	TickerEQ  *string `query:"ticker"`
	TickerLT  *string `query:"ticker.lt"`
	TickerLTE *string `query:"ticker.lte"`
	TickerGT  *string `query:"ticker.gt"`
	TickerGTE *string `query:"ticker.gte"`

	// Return results published on, before, or after this date.
	PublishedUtcEQ  *Millis `query:"published_utc"`
	PublishedUtcLT  *Millis `query:"published_utc.lt"`
	PublishedUtcLTE *Millis `query:"published_utc.lte"`
	PublishedUtcGT  *Millis `query:"published_utc.gt"`
	PublishedUtcGTE *Millis `query:"published_utc.gte"`

	// Sort field used for ordering.
	Sort *Sort `query:"sort"`

	// Order results based on the sort field.
	Order *Order `query:"order"`

	// Limit the number of results returned, default is 10 and max is 1000.
	Limit *int `query:"limit"`
}

func (p ListTickerNewsParams) WithTicker(c Comparator, q string) *ListTickerNewsParams {
	switch c {
	case EQ:
		p.TickerEQ = &q
	case LT:
		p.TickerLT = &q
	case LTE:
		p.TickerLTE = &q
	case GT:
		p.TickerGT = &q
	case GTE:
		p.TickerGTE = &q
	}
	return &p
}

func (p ListTickerNewsParams) WithPublishedUTC(c Comparator, q Millis) *ListTickerNewsParams {
	switch c {
	case EQ:
		p.PublishedUtcEQ = &q
	case LT:
		p.PublishedUtcLT = &q
	case LTE:
		p.PublishedUtcLTE = &q
	case GT:
		p.PublishedUtcGT = &q
	case GTE:
		p.PublishedUtcGTE = &q
	}
	return &p
}

func (p ListTickerNewsParams) WithSort(q Sort) *ListTickerNewsParams {
	p.Sort = &q
	return &p
}

func (p ListTickerNewsParams) WithOrder(q Order) *ListTickerNewsParams {
	p.Order = &q
	return &p
}

func (p ListTickerNewsParams) WithLimit(q int) *ListTickerNewsParams {
	p.Limit = &q
	return &p
}

// GetTickerRelatedCompaniesParams is the set of parameters for the GetTickerRelatedCompanies method.
type GetTickerRelatedCompaniesParams struct {
	// The ticker symbol of the asset.
	Ticker string `validate:"required" path:"ticker"`
}

// GetTickerDetailsResponse is the response returned by the GetTickerRelatedCompanies method.
type GetTickerRelatedCompaniesResponse struct {
	BaseResponse

	// List if related tickers.
	Results []RelatedCompany `json:"results,omitempty"`
}

// ListTickerNewsResponse is the response returned by the ListTickerNews method.
type ListTickerNewsResponse struct {
	BaseResponse

	// Ticker news results.
	Results []TickerNews `json:"results,omitempty"`
}

// GetTickerTypesParams is the set of parameters for the GetTickerTypes method.
type GetTickerTypesParams struct {
	// Filter by asset class.
	AssetClass *AssetClass `query:"asset_class"`

	// Filter by locale.
	Locale *MarketLocale `query:"locale"`
}

func (p GetTickerTypesParams) WithAssetClass(q AssetClass) *GetTickerTypesParams {
	p.AssetClass = &q
	return &p
}

func (p GetTickerTypesParams) WithLocale(q MarketLocale) *GetTickerTypesParams {
	p.Locale = &q
	return &p
}

// GetTickerTypesResponse is the response returned by the GetTickerTypes method.
type GetTickerTypesResponse struct {
	BaseResponse

	// Ticker type results.
	Results []TickerType `json:"results,omitempty"`
}

// Ticker contains detailed information on a specified ticker symbol.
type Ticker struct {
	Active                      bool           `json:"active"`
	Address                     CompanyAddress `json:"address,omitempty"`
	Branding                    Branding       `json:"branding,omitempty"`
	CIK                         string         `json:"cik,omitempty"`
	CompositeFIGI               string         `json:"composite_figi,omitempty"`
	CurrencyName                string         `json:"currency_name,omitempty"`
	CurrencySymbol              string         `json:"currency_symbol,omitempty"`
	BaseCurrencyName            string         `json:"base_currency_name,omitempty"`
	BaseCurrencySymbol          string         `json:"base_currency_symbol,omitempty"`
	DelistedUTC                 Time           `json:"delisted_utc,omitempty"`
	Description                 string         `json:"description,omitempty"`
	HomepageURL                 string         `json:"homepage_url,omitempty"`
	LastUpdatedUTC              Time           `json:"last_updated_utc,omitempty"`
	ListDate                    civil.Date     `json:"list_date,omitempty"`
	Locale                      string         `json:"locale,omitempty"`
	Market                      string         `json:"market,omitempty"`
	MarketCap                   float64        `json:"market_cap,omitempty"`
	Name                        string         `json:"name,omitempty"`
	PhoneNumber                 string         `json:"phone_number,omitempty"`
	PrimaryExchange             string         `json:"primary_exchange,omitempty"`
	ShareClassFIGI              string         `json:"share_class_figi,omitempty"`
	ShareClassSharesOutstanding int64          `json:"share_class_shares_outstanding,omitempty"`
	SICCode                     string         `json:"sic_code,omitempty"`
	SICDescription              string         `json:"sic_description,omitempty"`
	Ticker                      string         `json:"ticker,omitempty"`
	TickerRoot                  string         `json:"ticker_root,omitempty"`
	TickerSuffix                string         `json:"ticker_suffix,omitempty"`
	TotalEmployees              int32          `json:"total_employees,omitempty"`
	Type                        string         `json:"type,omitempty"`
	WeightedSharesOutstanding   int64          `json:"weighted_shares_outstanding,omitempty"`
	SourceFeed                  string         `json:"source_feed,omitempty"`
}

// CompanyAddress contains information on the physical address of a company.
type CompanyAddress struct {
	Address1   string `json:"address1,omitempty"`
	Address2   string `json:"address2,omitempty"` // todo: add this to the spec
	City       string `json:"city,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	State      string `json:"state,omitempty"`
}

// Branding contains information related to a company's brand.
type Branding struct {
	LogoURL string `json:"logo_url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

// TickerNews contains information on a ticker news article.
type TickerNews struct {
	AMPURL       string     `json:"amp_url,omitempty"`
	ArticleURL   string     `json:"article_url,omitempty"`
	Author       string     `json:"author,omitempty"`
	Description  string     `json:"description,omitempty"`
	ID           string     `json:"id,omitempty"`
	ImageURL     string     `json:"image_url,omitempty"`
	Insights     []Insights `json:"insights"`
	Keywords     []string   `json:"keywords,omitempty"`
	PublishedUTC Time       `json:"published_utc,omitempty"`
	Publisher    Publisher  `json:"publisher,omitempty"`
	Tickers      []string   `json:"tickers,omitempty"`
	Title        string     `json:"title,omitempty"`
}

// Insights contains sentiment, reasoning, and the ticker symbol associated with the insight.
type Insights struct {
	Ticker             string `json:"ticker"`
	Sentiment          string `json:"sentiment"`
	SentimentReasoning string `json:"sentiment_reasoning"`
}

// Publisher contains information on a new article publisher.
type Publisher struct {
	FaviconURL  string `json:"favicon_url,omitempty"`
	HomepageURL string `json:"homepage_url,omitempty"`
	LogoURL     string `json:"logo_url,omitempty"`
	Name        string `json:"name,omitempty"`
}

// RelatedCompany represents a related ticker based on news or sec filings.
type RelatedCompany struct {
	Ticker string `json:"ticker,omitempty"`
}

// TickerType represents a type of ticker with a code that the API understands.
type TickerType struct {
	AssetClass  string `json:"asset_class,omitempty"`
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Locale      string `json:"locale,omitempty"`
}

// GetTickerEventsParams is the set of parameters for the GetTickerEvents method.
type GetTickerEventsParams struct {
	// ID Identifier of an asset. This can currently be a Ticker, CUSIP, or Composite FIGI.
	// When given a ticker, we return events for the entity currently represented by that ticker.
	// To find events for entities previously associated with a ticker, find the relevant identifier
	// using the Ticker Details Endpoint (https://polygon.io/docs/stocks/get_v3_reference_tickers__ticker).
	ID string `validate:"required" path:"id"`

	// A comma-separated list of the types of event to include. Currently, ticker_change is the only supported event_type. Leave blank to return all supported event_types.
	Types *string `query:"types"`
}

func (p GetTickerEventsParams) WithTypes(types ...string) *GetTickerEventsParams {
	q := strings.Join(types, ",")
	p.Types = &q
	return &p
}

// GetTickerEventsResponse is the response returned by the GetTickerEvents method.
type GetTickerEventsResponse struct {
	BaseResponse
	Results []TickerEventResult `json:"results,omitempty"`
}

// TickerEventResult is the data for a ticker event.
type TickerEventResult struct {
	// Name is the company name.
	Name   string        `json:"name"`
	Events []TickerEvent `json:"events"`
}

// TickerEvent contains the data for the different type of ticker events that could occur.
type TickerEvent struct {
	Date         civil.Date         `json:"date"`
	Type         string             `json:"type"`
	TickerChange *TickerChangeEvent `json:"ticker_change,omitempty"`
}

// TickerChangeEvent represents the data relevant to a ticker_change typed event.
type TickerChangeEvent struct {
	Ticker string `json:"ticker"`
}
