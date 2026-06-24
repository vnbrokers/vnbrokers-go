package dto

import (
	"bytes"
	"encoding/json"
)

type JSONObject map[string]any

type GetCompanyFinancialAnalysisRequest struct {
	Symbol string
	Period string
}

type GetCompanyFinancialOverviewRequest struct{ Symbol string }

type GetEconomicCalendarEventsRequest struct {
	Weeks   int64
	Country string
}

type GetFinancialStatementV2Request struct {
	Symbol string
	Type   string
	Period string
}

type GetGrowthBenchmarkRequest struct {
	FundNames string
	Amount    int64
	Period    string
}

type GetNAVBenchmarkRequest struct {
	FundNames string
	Period    string
	FromMonth string
	ToMonth   string
}

type GetFundCertificatesRequest struct {
	FundType      string
	FundCompanyID *int64
}

type GetFundSuggestionsRequest struct{ FundName string }

type GetFundMonthsRequest struct{ Fund string }

type GetFundNAVHistoriesRequest struct {
	FundName string
	Period   string
}

type GetFundPortfolioRequest struct {
	Fund  string
	Month string
}

type GetGlobalNewsListRequest struct {
	Category string
	Page     int64
	PageSize int64
}

type GetGlobalNewsDetailRequest struct{ ID int64 }

type GetChartDataRequest struct{ Days int64 }

type GetIndexRealtimeRequest struct{ Index string }

type GetMacroDataRequest struct {
	Type    string
	Country string
	Period  string
}

type GetMarketDataRequest struct {
	Type  string
	Limit int64
}

type GetStockEventsRequest struct {
	Stock    string
	Stocks   string
	FromDate string
	ToDate   string
}

type GetPriceHistoriesChartRequest struct {
	Symbol     string
	Resolution string
	From       int64
	To         int64
}

type GetRecommendationReportsRequest struct{ Symbol string }

type GetStockRealtimeRequest struct {
	Symbol   string
	Symbols  string
	Exchange string
}

type GetTradingEconomicsDataRequest struct {
	Country  string
	Category string
	Year     int64
}

type GetFundMonthsResponse = []string

type ChartDataResponse = []JSONObject

type StockRealtimeResponse struct {
	Single *StockRealtime
	Many   []StockRealtime
}

func (r *StockRealtimeResponse) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")) {
		return nil
	}
	if trimmed[0] == '[' {
		return json.Unmarshal(trimmed, &r.Many)
	}
	var single StockRealtime
	if err := json.Unmarshal(trimmed, &single); err != nil {
		return err
	}
	r.Single = &single
	return nil
}

type RecommendationReportsResponse struct {
	Empty bool
	Data  *RecommendationReportResponse
}

func (r *RecommendationReportsResponse) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")) {
		return nil
	}
	if trimmed[0] == '[' {
		var empty []any
		if err := json.Unmarshal(trimmed, &empty); err != nil {
			return err
		}
		r.Empty = len(empty) == 0
		return nil
	}
	var payload RecommendationReportResponse
	if err := json.Unmarshal(trimmed, &payload); err != nil {
		return err
	}
	r.Data = &payload
	return nil
}
