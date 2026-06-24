// Code generated from Finhay OpenAPI v2; DO NOT EDIT.

package dto

type CommodityIndexData struct {
	// Index identifier
	Index string `json:"index,omitempty"`
	// Display name
	Name string `json:"name,omitempty"`
	// Short display name
	ShortName string `json:"short_name,omitempty"`
	// Buy price
	BuyValue float64 `json:"buy_value,omitempty"`
	// Sell price
	SellValue float64 `json:"sell_value,omitempty"`
	// USD equivalent value
	USDValue float64 `json:"usd_value,omitempty"`
	// VND value
	VNDValue float64 `json:"vnd_value,omitempty"`
	// Price date
	Date string `json:"date,omitempty"`
	// Data provider name
	Provider string `json:"provider,omitempty"`
	// Provider icon URL
	ProviderIcon *string `json:"provider_icon,omitempty"`
	// Last update timestamp
	UpdatedAt string `json:"updated_at,omitempty"`
	// Price change percentage
	ChangePercent float64 `json:"change_percent,omitempty"`
	// Buy value change percentage
	BuyValueChangePercent float64 `json:"buy_value_change_percent,omitempty"`
	// Sell value change percentage
	SellValueChangePercent float64 `json:"sell_value_change_percent,omitempty"`
}

type CryptoCurrency struct {
	// Cryptocurrency name
	Name string `json:"name,omitempty"`
	// Cryptocurrency symbol
	Symbol string `json:"symbol,omitempty"`
	// Icon URL
	IconURL string `json:"icon_url,omitempty"`
	// Current price in USD
	Price float64 `json:"price,omitempty"`
	// Price formatted in Vietnamese locale
	FormattedPrice string `json:"formatted_price,omitempty"`
	// Price change in the last 1 hour (%)
	PercentChange1h float64 `json:"percent_change_1h,omitempty"`
	// Price change in the last 24 hours (%)
	PercentChange24h float64 `json:"percent_change_24h,omitempty"`
	// Price change in the last 7 days (%)
	PercentChange7d float64 `json:"percent_change_7d,omitempty"`
	// Price change in the last 30 days (%)
	PercentChange30d float64 `json:"percent_change_30d,omitempty"`
	// Market capitalization in USD
	MarketCap float64 `json:"market_cap,omitempty"`
	// URL or data for 30-day sparkline chart
	Last30dChart string `json:"last_30d_chart,omitempty"`
}

type EconomicCalendarEvent struct {
	ID int64 `json:"id,omitempty"`
	// Event date in `YYYY-MM-DD` format
	Date string `json:"date,omitempty"`
	// Country name Enum: China, Euro Area, Japan, United States, United Kingdom, Vietnam
	Country string `json:"country,omitempty"`
	// Event name
	Event string `json:"event,omitempty"`
	// Actual value (empty if not yet released)
	Actual string `json:"actual,omitempty"`
	// Previous period value
	Previous string `json:"previous,omitempty"`
	// Market consensus forecast
	Consensus string `json:"consensus,omitempty"`
	// Analyst forecast
	Forecast string `json:"forecast,omitempty"`
	// Impact level (1 = low, 2 = medium, 3 = high)
	Impact int64 `json:"impact,omitempty"`
	// Event category
	Category string `json:"category,omitempty"`
}

// Financial metrics for one period. Contains `year` always; `quarter` is present only when period=quarterly. All other fields are dynamic metric codes (e.g. pe, roe, eps, pb) with numeric or null values.
type FinancialAnalysisEntry map[string]any

type FinancialData struct {
	// Type of financial data Enum: GOLD, SILVER, CRYPTO, BANK_INTEREST_RATE, USD_EXCHANGE_RATE
	FinancialType string `json:"financial_type,omitempty"`
	// Display name of the financial type
	Name string `json:"name,omitempty"`
	// Array of financial data values. Shape varies by `financial_type`: single-value types use `FinancialDataSingleValue`, buy/sell types use `FinancialDataBuySellValue`, USD/VND types use `FinancialDataUSDVNDValue`.
	Values []any `json:"values,omitempty"`
	// ISO timestamp of last update
	LastUpdated string `json:"last_updated,omitempty"`
}

type FinancialDataBuySellValue struct {
	Name          string  `json:"name,omitempty"`
	BuyValue      float64 `json:"buy_value,omitempty"`
	SellValue     float64 `json:"sell_value,omitempty"`
	ChangePercent float64 `json:"change_percent,omitempty"`
}

type FinancialDataSingleValue struct {
	Name          string  `json:"name,omitempty"`
	Value         float64 `json:"value,omitempty"`
	ChangePercent float64 `json:"change_percent,omitempty"`
}

type FinancialDataUSDVNDValue struct {
	Name     string  `json:"name,omitempty"`
	USDValue float64 `json:"usd_value,omitempty"`
	VNDValue float64 `json:"vnd_value,omitempty"`
}

// One metric value for one period.
type FinancialMetricValue struct {
	Symbol string `json:"symbol,omitempty"`
	// Metric identifier (e.g. "tongDoanhThu", "lnst")
	MetricCode  string  `json:"metricCode,omitempty"`
	MetricValue float64 `json:"metricValue,omitempty"`
	// Period type identifier
	TimeType string `json:"timeType,omitempty"`
	Year     int64  `json:"year,omitempty"`
	// 0 for annual records
	Quarter int64 `json:"quarter,omitempty"`
}

type FinancialOverview struct {
	// Price-to-Earnings ratio
	PE *float64 `json:"pe,omitempty"`
	// Price-to-Book ratio
	PB *float64 `json:"pb,omitempty"`
	// EV/EBITDA ratio
	EvEbitda *float64 `json:"ev_ebitda,omitempty"`
	// Industry average ratios
	Industry struct {
		PE       *float64 `json:"pe,omitempty"`
		PB       *float64 `json:"pb,omitempty"`
		EvEbitda *float64 `json:"ev_ebitda,omitempty"`
	} `json:"industry,omitempty"`
	// Gross profit margin
	GrossMargin *float64 `json:"gross_margin,omitempty"`
	// Return on Equity
	ROE *float64 `json:"roe,omitempty"`
	// Earnings per Share
	EPS           *float64 `json:"eps,omitempty"`
	DividendYield *float64 `json:"dividend_yield,omitempty"`
	// Net Interest Margin (banks only)
	Nim *float64 `json:"nim,omitempty"`
	// Margin loan to equity ratio (securities firms only)
	MarginLoanToEquityRatio *float64 `json:"margin_loan_to_equity_ratio,omitempty"`
	// Return on Assets
	Roa *float64 `json:"roa,omitempty"`
}

type IndexRealtime struct {
	// Index code (echoes the requested `index`).
	Index string `json:"index,omitempty"`
	// Current index points value.
	IndexValue float64 `json:"indexValue,omitempty"`
	// Point change from reference.
	Change float64 `json:"change,omitempty"`
	// Percent change from reference (already in percent units, e.g. 0.12 = 0.12%).
	ChangePercent float64 `json:"changePercent,omitempty"`
	// Reference value — previous session close.
	Reference float64 `json:"reference,omitempty"`
	// Total matched volume across the market (shares).
	AllQuantity int64 `json:"allQuantity,omitempty"`
	// Total matched trading value across the market (VND).
	AllValue float64 `json:"allValue,omitempty"`
	// Number of advancing symbols.
	Advances int64 `json:"advances,omitempty"`
	// Number of declining symbols.
	Declines int64 `json:"declines,omitempty"`
	// Number of unchanged symbols.
	Nochanges int64 `json:"nochanges,omitempty"`
	// Number of symbols at ceiling price.
	Ceiling int64 `json:"ceiling,omitempty"`
	// Number of symbols at floor price.
	Floor int64 `json:"floor,omitempty"`
	// Intraday index point series (index-aligned with times/volumes).
	Values []float64 `json:"values,omitempty"`
	// Intraday volume series (index-aligned with times/values).
	Volumes []float64 `json:"volumes,omitempty"`
	// Intraday timestamp series — Unix ms.
	Times []int64 `json:"times,omitempty"`
	// Intraday advancing-symbol series. KRX indices only.
	AdvancesArr []int64 `json:"advancesArr,omitempty"`
	// Intraday declining-symbol series. KRX indices only.
	DeclinesArr []int64 `json:"declinesArr,omitempty"`
	// Intraday unchanged-symbol series. KRX indices only.
	NochangesArr []int64 `json:"nochangesArr,omitempty"`
	// Intraday ceiling-count series. KRX indices only.
	Ceilings []int64 `json:"ceilings,omitempty"`
	// Intraday floor-count series. KRX indices only.
	Floors []int64 `json:"floors,omitempty"`
	// Current trading session state of the exchange.
	SessionInExchange string `json:"sessionInExchange,omitempty"`
	// Index name (currently same as `index`).
	Name string `json:"name,omitempty"`
}

type MacroData struct {
	// Macro data type Enum: IIP, CPI, PMI, PCE, CORE_PCE, NFP, GOODS_RETAIL, SERVICE_RETAIL, TOTAL_EXPORT, FDI_EXPORT, DOMESTIC_EXPORT, FED_FUNDS_RATE, INTERBANK_RATE, GOVERNMENT_10Y_BOND_YIELD, UNEMPLOYMENT_RATE
	TypeValue string `json:"type,omitempty"`
	// Enum: VN, US
	Country string `json:"country,omitempty"`
	// Year-month in `YYYY-MM` format
	Month string `json:"month,omitempty"`
	// Macro indicator value
	Value float64 `json:"value,omitempty"`
	// Exact date in `YYYY-MM-DD` format (if available)
	Date *string `json:"date,omitempty"`
}

type MarketData struct {
	TypeValue string `json:"type,omitempty"`
	// Enum: US, KR, HK, CN, JP, UK
	Country string `json:"country,omitempty"`
	// Date in YYYY-MM-DD format
	Date  string  `json:"date,omitempty"`
	Year  int64   `json:"year,omitempty"`
	Month int64   `json:"month,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// Columnar arrays — all arrays are the same length. `time[i]`, `open[i]`, `close[i]`, `high[i]`, `low[i]`, `volume[i]` represent one data point.
type PriceHistoriesChart struct {
	Symbol string `json:"symbol,omitempty"`
	// Enum: 1D, 5, 15, 30, 1H, 4H
	Resolution string `json:"resolution,omitempty"`
	// Unix timestamps (seconds)
	Time []float64 `json:"time,omitempty"`
	// Open prices
	Open []float64 `json:"open,omitempty"`
	// Close prices
	Close []float64 `json:"close,omitempty"`
	// High prices
	High []float64 `json:"high,omitempty"`
	// Low prices
	Low []float64 `json:"low,omitempty"`
	// Trading volumes
	Volume []float64 `json:"volume,omitempty"`
}

type StockInfoV1 struct {
	Symbol string `json:"symbol,omitempty"`
	Name   string `json:"name,omitempty"`
	// Enum: HOSE, HNX, UPCOM
	Exchange string `json:"exchange,omitempty"`
	// Floor price
	Floor float64 `json:"floor,omitempty"`
	// Ceiling price
	Ceiling float64 `json:"ceiling,omitempty"`
	// Reference price
	Reference float64 `json:"reference,omitempty"`
	// Enum: STOCK, ETF, BOND, CW, FUTURES
	StockType          string   `json:"stock_type,omitempty"`
	Price              float64  `json:"price,omitempty"`
	PriceChange        float64  `json:"price_change,omitempty"`
	PriceChangePercent float64  `json:"price_change_percent,omitempty"`
	Volume             *float64 `json:"volume,omitempty"`
	TotalVolume        *float64 `json:"total_volume,omitempty"`
}

type StockRealtime struct {
	Symbol string   `json:"symbol,omitempty"`
	Price  *float64 `json:"price,omitempty"`
	Volume *float64 `json:"volume,omitempty"`
	// Price change from reference
	Change *float64 `json:"change,omitempty"`
	// Price change percentage from reference
	ChangePercent *float64 `json:"changePercent,omitempty"`
	// Ceiling price
	Ceiling *float64 `json:"ceiling,omitempty"`
	// Floor price
	Floor *float64 `json:"floor,omitempty"`
	// Reference price
	Reference *float64 `json:"reference,omitempty"`
	// Average (medium) price
	Average *float64 `json:"average,omitempty"`
	High    *float64 `json:"high,omitempty"`
	Low     *float64 `json:"low,omitempty"`
	Open    *float64 `json:"open,omitempty"`
	// Same as `price`
	Close         *float64 `json:"close,omitempty"`
	BuyPrice1     *float64 `json:"buyPrice1,omitempty"`
	BuyPrice2     *float64 `json:"buyPrice2,omitempty"`
	BuyPrice3     *float64 `json:"buyPrice3,omitempty"`
	BuyVol1       *float64 `json:"buyVol1,omitempty"`
	BuyVol2       *float64 `json:"buyVol2,omitempty"`
	BuyVol3       *float64 `json:"buyVol3,omitempty"`
	SellPrice1    *float64 `json:"sellPrice1,omitempty"`
	SellPrice2    *float64 `json:"sellPrice2,omitempty"`
	SellPrice3    *float64 `json:"sellPrice3,omitempty"`
	SellVol1      *float64 `json:"sellVol1,omitempty"`
	SellVol2      *float64 `json:"sellVol2,omitempty"`
	SellVol3      *float64 `json:"sellVol3,omitempty"`
	TotalVolume   *float64 `json:"totalVolume,omitempty"`
	TotalValue    *float64 `json:"totalValue,omitempty"`
	ForeignBought *float64 `json:"foreignBought,omitempty"`
	ForeignSold   *float64 `json:"foreignSold,omitempty"`
	ForeignRemain *float64 `json:"foreignRemain,omitempty"`
	RemainBid     *float64 `json:"remainBid,omitempty"`
	RemainAsk     *float64 `json:"remainAsk,omitempty"`
	// Normalized stock type Enum: STOCK, ETF, BOND, CW, FUTURES
	StockType string `json:"stockType,omitempty"`
	// Normalized exchange code Enum: HOSE, HNX, UPCOM
	Exchange string `json:"exchange,omitempty"`
	Name     string `json:"name,omitempty"`
	// Unix timestamp
	CreatedAt float64  `json:"createdAt,omitempty"`
	PE        *float64 `json:"pe,omitempty"`
	PB        *float64 `json:"pb,omitempty"`
	ROE       *float64 `json:"roe,omitempty"`
	MarketCap *float64 `json:"marketCap,omitempty"`
	// Enum: Micro Cap, Small Cap, Mid Cap, Large Cap
	MarketCapCategory *string        `json:"marketCapCategory,omitempty"`
	HasNewestNews     bool           `json:"hasNewestNews,omitempty"`
	StockSummary      *string        `json:"stockSummary,omitempty"`
	AdditionalInfo    map[string]any `json:"additionalInfo,omitempty"`
	// KRX stocks only
	SymbolStatus *string `json:"symbolStatus,omitempty"`
	// KRX stocks only
	SymbolStatusCode *string `json:"symbolStatusCode,omitempty"`
	// KRX stocks only
	FloorCode         *string  `json:"floorCode,omitempty"`
	InfluenceScore    *float64 `json:"influenceScore,omitempty"`
	OutstandingShares *float64 `json:"outstandingShares,omitempty"`
}

type TradingEconomicsData struct {
	// Indicator name
	Indicator string `json:"indicator,omitempty"`
	// Country name Enum: China, Euro Area, Japan, United States, United Kingdom, Vietnam
	Country string `json:"country,omitempty"`
	// Indicator category Enum: GDP, Labour, Prices, Money, Trade, Government, Business, Consumer, Housing
	Category string `json:"category,omitempty"`
	// Most recent value
	LastValue *float64 `json:"lastValue,omitempty"`
	// Prior period value
	PreviousValue *float64 `json:"previousValue,omitempty"`
	// Year of the data point
	Year int64 `json:"year,omitempty"`
	// Month of the data point (1–12)
	Month int64 `json:"month,omitempty"`
	// Unit of measurement
	Unit string `json:"unit,omitempty"`
}
