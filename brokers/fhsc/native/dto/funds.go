// Code generated from Finhay OpenAPI v2; DO NOT EDIT.

package dto

type FundCertificate struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	// Enum: STOCK_FUND, BOND_FUND, BALANCE_FUND
	TypeValue string `json:"type,omitempty"`
	// Assets under management in VND.
	Aum    int64   `json:"aum,omitempty"`
	Rating float64 `json:"rating,omitempty"`
}

type FundCompany struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ShortName string `json:"short_name,omitempty"`
	ImageURL  string `json:"image_url,omitempty"`
}

type GrowthBenchmark struct {
	FundName      string  `json:"fund_name,omitempty"`
	ProfitPercent float64 `json:"profit_percent,omitempty"`
	// Projected portfolio value in VND.
	CurrentAmount int64 `json:"current_amount,omitempty"`
}

type NAVBenchmark struct {
	FundName   string `json:"fund_name,omitempty"`
	NAVRecords []struct {
		Date          string  `json:"date,omitempty"`
		Navpf         float64 `json:"navpf,omitempty"`
		ChangePercent float64 `json:"change_percent,omitempty"`
	} `json:"nav_records,omitempty"`
}

type NAVHistories struct {
	NAVHistories []struct {
		Date string  `json:"date,omitempty"`
		NAV  float64 `json:"nav,omitempty"`
		// Benchmark index value on the same date.
		BenchmarkValue float64 `json:"benchmark_value,omitempty"`
	} `json:"nav_histories,omitempty"`
	// Benchmark name (e.g. `VN-INDEX` for stock funds).
	BenchmarkName string `json:"benchmark_name,omitempty"`
}

type PortfolioFund struct {
	// Portfolio month
	Month string `json:"month,omitempty"`
	// Portfolio holdings sorted by per_nav descending
	Entries []PortfolioFundEntry `json:"entries,omitempty"`
	// Realtime stock info for the holdings
	StocksInfo []StockInfoV1 `json:"stocks_info,omitempty"`
}

type PortfolioFundEntry struct {
	FundCode string `json:"fund_code,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
	Sector   string `json:"sector,omitempty"`
	// Percentage of NAV
	PerNAV float64 `json:"per_nav,omitempty"`
	Month  string  `json:"month,omitempty"`
}

type SuggestedFund struct {
	FundName string `json:"fund_name,omitempty"`
	// How the fund was matched: - `NET_ASSET_EQUIVALENT`: similar AUM size. - `GROWTH_EQUIVALENT`: similar historical growth profile. Enum: NET_ASSET_EQUIVALENT, GROWTH_EQUIVALENT
	Criteria string `json:"criteria,omitempty"`
}
