package marketdata

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/internal/httpx"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const (
	CapabilityGetBankInterestRates        core.Capability = "native.marketdata.get_bank_interest_rates"
	CapabilityGetCompanyFinancialAnalysis core.Capability = "native.marketdata.get_company_financial_analysis"
	CapabilityGetCompanyFinancialOverview core.Capability = "native.marketdata.get_company_financial_overview"
	CapabilityGetTopTrendingCryptos       core.Capability = "native.marketdata.get_top_trending_cryptos"
	CapabilityGetEconomicCalendarEvents   core.Capability = "native.marketdata.get_economic_calendar_events"
	CapabilityGetFinancialData            core.Capability = "native.marketdata.get_financial_data"
	CapabilityGetFinancialStatementV2     core.Capability = "native.marketdata.get_financial_statement_v2"
	CapabilityGetGrowthBenchmark          core.Capability = "native.marketdata.get_growth_benchmark"
	CapabilityGetNAVBenchmark             core.Capability = "native.marketdata.get_nav_benchmark"
	CapabilityGetFundCompanies            core.Capability = "native.marketdata.get_fund_companies"
	CapabilityGetFundSuggestions          core.Capability = "native.marketdata.get_fund_suggestions"
	CapabilityGetFundMonths               core.Capability = "native.marketdata.get_fund_months"
	CapabilityGetFundNAVHistories         core.Capability = "native.marketdata.get_fund_nav_histories"
	CapabilityGetFundPortfolio            core.Capability = "native.marketdata.get_fund_portfolio"
	CapabilityGetFundCertificates         core.Capability = "native.marketdata.get_fund_certificates"
	CapabilityGetGlobalNewsList           core.Capability = "native.marketdata.get_global_news_list"
	CapabilityGetGlobalNewsDetail         core.Capability = "native.marketdata.get_global_news_detail"
	CapabilityGetGoldChartData            core.Capability = "native.marketdata.get_gold_chart_data"
	CapabilityGetGoldProviderData         core.Capability = "native.marketdata.get_gold_provider_data"
	CapabilityGetGoldData                 core.Capability = "native.marketdata.get_gold_data"
	CapabilityGetIndexRealtime            core.Capability = "native.marketdata.get_index_realtime"
	CapabilityGetMacroData                core.Capability = "native.marketdata.get_macro_data"
	CapabilityGetMarketData               core.Capability = "native.marketdata.get_market_data"
	CapabilityGetMetalProviderData        core.Capability = "native.marketdata.get_metal_provider_data"
	CapabilityGetStockEvents              core.Capability = "native.marketdata.get_stock_events"
	CapabilityGetPriceHistoriesChart      core.Capability = "native.marketdata.get_price_histories_chart"
	CapabilityGetRecommendationReports    core.Capability = "native.marketdata.get_recommendation_reports"
	CapabilityGetSilverChartData          core.Capability = "native.marketdata.get_silver_chart_data"
	CapabilityGetSilverData               core.Capability = "native.marketdata.get_silver_data"
	CapabilityGetStockRealtime            core.Capability = "native.marketdata.get_stock_realtime"
	CapabilityGetTradingEconomicsData     core.Capability = "native.marketdata.get_trading_economics_data"
)

type RealtimeService interface{}

type Service interface {
	Realtime() RealtimeService
	GetBankInterestRates(context.Context) (*dto.JSONObject, error)
	GetCompanyFinancialAnalysis(context.Context, dto.GetCompanyFinancialAnalysisRequest) (*[]dto.FinancialAnalysisEntry, error)
	GetCompanyFinancialOverview(context.Context, dto.GetCompanyFinancialOverviewRequest) (*dto.FinancialOverview, error)
	GetTopTrendingCryptos(context.Context) (*[]dto.CryptoCurrency, error)
	GetEconomicCalendarEvents(context.Context, dto.GetEconomicCalendarEventsRequest) (*[]dto.EconomicCalendarEvent, error)
	GetFinancialData(context.Context) (*[]dto.FinancialData, error)
	GetFinancialStatementV2(context.Context, dto.GetFinancialStatementV2Request) (*[]dto.FinancialMetricValue, error)
	GetGrowthBenchmark(context.Context, dto.GetGrowthBenchmarkRequest) (*[]dto.GrowthBenchmark, error)
	GetNAVBenchmark(context.Context, dto.GetNAVBenchmarkRequest) (*[]dto.NAVBenchmark, error)
	GetFundCompanies(context.Context) (*[]dto.FundCompany, error)
	GetFundSuggestions(context.Context, dto.GetFundSuggestionsRequest) (*[]dto.SuggestedFund, error)
	GetFundMonths(context.Context, dto.GetFundMonthsRequest) (*dto.GetFundMonthsResponse, error)
	GetFundNAVHistories(context.Context, dto.GetFundNAVHistoriesRequest) (*dto.NAVHistories, error)
	GetFundPortfolio(context.Context, dto.GetFundPortfolioRequest) (*dto.PortfolioFund, error)
	GetFundCertificates(context.Context, dto.GetFundCertificatesRequest) (*[]dto.FundCertificate, error)
	GetGlobalNewsList(context.Context, dto.GetGlobalNewsListRequest) (*dto.GlobalNewsPage, error)
	GetGlobalNewsDetail(context.Context, dto.GetGlobalNewsDetailRequest) (*dto.GlobalNewsDetail, error)
	GetGoldChartData(context.Context, dto.GetChartDataRequest) (*dto.ChartDataResponse, error)
	GetGoldProviderData(context.Context) (*[]dto.CommodityIndexData, error)
	GetGoldData(context.Context) (*[]dto.CommodityIndexData, error)
	GetIndexRealtime(context.Context, dto.GetIndexRealtimeRequest) (*[]dto.IndexRealtime, error)
	GetMacroData(context.Context, dto.GetMacroDataRequest) (*[]dto.MacroData, error)
	GetMarketData(context.Context, dto.GetMarketDataRequest) (*[]dto.MarketData, error)
	GetMetalProviderData(context.Context) (*[]dto.CommodityIndexData, error)
	GetStockEvents(context.Context, dto.GetStockEventsRequest) (*[]dto.StockEventResponse, error)
	GetPriceHistoriesChart(context.Context, dto.GetPriceHistoriesChartRequest) (*dto.PriceHistoriesChart, error)
	GetRecommendationReports(context.Context, dto.GetRecommendationReportsRequest) (*dto.RecommendationReportsResponse, error)
	GetSilverChartData(context.Context, dto.GetChartDataRequest) (*dto.ChartDataResponse, error)
	GetSilverData(context.Context) (*[]dto.CommodityIndexData, error)
	GetStockRealtime(context.Context, dto.GetStockRealtimeRequest) (*dto.StockRealtimeResponse, error)
	GetTradingEconomicsData(context.Context, dto.GetTradingEconomicsDataRequest) (*[]dto.TradingEconomicsData, error)
}

type Dependencies struct {
	BaseURL           string
	Headers           func(authenticated bool, hasBody bool) map[string]string
	RequireCapability func(core.Capability) error
	Send              func(context.Context, string, transport.HTTPRequest) (transport.HTTPResponse, error)
}

type service struct {
	dependencies Dependencies
	realtime     RealtimeService
}

type envelope[T any] struct {
	Status       *int64  `json:"status,omitempty"`
	ErrorCode    *string `json:"error_code,omitempty"`
	Message      *string `json:"message,omitempty"`
	PopupMessage *string `json:"popup_message,omitempty"`
	Title        *string `json:"title,omitempty"`
	Data         *T      `json:"data,omitempty"`
	Result       *T      `json:"result,omitempty"`
}

func NewService(dependencies Dependencies, realtimeServices ...RealtimeService) Service {
	var realtimeService RealtimeService
	if len(realtimeServices) > 0 {
		realtimeService = realtimeServices[0]
	}
	return &service{dependencies: dependencies, realtime: realtimeService}
}

func (s *service) Realtime() RealtimeService { return s.realtime }

func do[T any](s *service, ctx context.Context, capability core.Capability, method, path string, query url.Values) (*T, error) {
	if err := s.dependencies.RequireCapability(capability); err != nil {
		return nil, err
	}
	response, err := s.dependencies.Send(ctx, string(capability), transport.HTTPRequest{
		Method:  method,
		URL:     httpx.URL(s.dependencies.BaseURL, path, query),
		Headers: s.dependencies.Headers(true, false),
	})
	if err != nil {
		return nil, err
	}
	env, err := httpx.DecodeResponse[envelope[T]]("fhsc", string(capability), "decode FHSC marketdata response", response)
	if err != nil {
		return nil, err
	}
	if code := stringValue(env.ErrorCode); code != "" && code != "0" {
		return nil, sdkerrors.BrokerRejected("fhsc", string(capability), code, firstNonEmpty(env.Message, env.PopupMessage, env.Title), httpx.RawPayload(response))
	}
	if env.Result != nil {
		return env.Result, nil
	}
	if env.Data != nil {
		return env.Data, nil
	}
	result := new(T)
	return result, nil
}

func (s *service) GetBankInterestRates(ctx context.Context) (*dto.JSONObject, error) {
	return do[dto.JSONObject](s, ctx, CapabilityGetBankInterestRates, "GET", "/market/financial-data/bank-interest-rates", url.Values{})
}

func (s *service) GetCompanyFinancialAnalysis(ctx context.Context, request dto.GetCompanyFinancialAnalysisRequest) (*[]dto.FinancialAnalysisEntry, error) {
	query := url.Values{}
	set(query, "symbol", request.Symbol)
	setOptional(query, "period", request.Period)
	return do[[]dto.FinancialAnalysisEntry](s, ctx, CapabilityGetCompanyFinancialAnalysis, "GET", "/market/company-financial/analysis", query)
}

func (s *service) GetCompanyFinancialOverview(ctx context.Context, request dto.GetCompanyFinancialOverviewRequest) (*dto.FinancialOverview, error) {
	query := url.Values{}
	set(query, "symbol", request.Symbol)
	return do[dto.FinancialOverview](s, ctx, CapabilityGetCompanyFinancialOverview, "GET", "/market/company-financial/overview", query)
}

func (s *service) GetTopTrendingCryptos(ctx context.Context) (*[]dto.CryptoCurrency, error) {
	return do[[]dto.CryptoCurrency](s, ctx, CapabilityGetTopTrendingCryptos, "GET", "/market/financial-data/cryptos/top-trending", url.Values{})
}

func (s *service) GetEconomicCalendarEvents(ctx context.Context, request dto.GetEconomicCalendarEventsRequest) (*[]dto.EconomicCalendarEvent, error) {
	query := url.Values{}
	setOptionalInt(query, "weeks", request.Weeks)
	setOptional(query, "country", request.Country)
	return do[[]dto.EconomicCalendarEvent](s, ctx, CapabilityGetEconomicCalendarEvents, "GET", "/market/financial-data/economic-calendar-events", query)
}

func (s *service) GetFinancialData(ctx context.Context) (*[]dto.FinancialData, error) {
	return do[[]dto.FinancialData](s, ctx, CapabilityGetFinancialData, "GET", "/market/financial-data", url.Values{})
}

func (s *service) GetFinancialStatementV2(ctx context.Context, request dto.GetFinancialStatementV2Request) (*[]dto.FinancialMetricValue, error) {
	query := url.Values{}
	set(query, "symbol", request.Symbol)
	set(query, "type", request.Type)
	setOptional(query, "period", request.Period)
	return do[[]dto.FinancialMetricValue](s, ctx, CapabilityGetFinancialStatementV2, "GET", "/market/v2/financial-statement/statement", query)
}

func (s *service) GetGrowthBenchmark(ctx context.Context, request dto.GetGrowthBenchmarkRequest) (*[]dto.GrowthBenchmark, error) {
	query := url.Values{}
	set(query, "fund-names", request.FundNames)
	set(query, "amount", strconv.FormatInt(request.Amount, 10))
	set(query, "period", request.Period)
	return do[[]dto.GrowthBenchmark](s, ctx, CapabilityGetGrowthBenchmark, "GET", "/fund-trading/public/fund-certificates/benchmark/growth", query)
}

func (s *service) GetNAVBenchmark(ctx context.Context, request dto.GetNAVBenchmarkRequest) (*[]dto.NAVBenchmark, error) {
	query := url.Values{}
	set(query, "fund-names", request.FundNames)
	setOptional(query, "period", request.Period)
	setOptional(query, "from-month", request.FromMonth)
	setOptional(query, "to-month", request.ToMonth)
	return do[[]dto.NAVBenchmark](s, ctx, CapabilityGetNAVBenchmark, "GET", "/fund-trading/public/fund-certificates/benchmark/nav", query)
}

func (s *service) GetFundCompanies(ctx context.Context) (*[]dto.FundCompany, error) {
	return do[[]dto.FundCompany](s, ctx, CapabilityGetFundCompanies, "GET", "/fund-trading/public/fund-companies", url.Values{})
}

func (s *service) GetFundSuggestions(ctx context.Context, request dto.GetFundSuggestionsRequest) (*[]dto.SuggestedFund, error) {
	path := "/fund-trading/public/fund-certificates/" + escaped(request.FundName) + "/suggestions"
	return do[[]dto.SuggestedFund](s, ctx, CapabilityGetFundSuggestions, "GET", path, url.Values{})
}

func (s *service) GetFundMonths(ctx context.Context, request dto.GetFundMonthsRequest) (*dto.GetFundMonthsResponse, error) {
	path := "/market/funds/" + escaped(request.Fund) + "/months"
	return do[dto.GetFundMonthsResponse](s, ctx, CapabilityGetFundMonths, "GET", path, url.Values{})
}

func (s *service) GetFundNAVHistories(ctx context.Context, request dto.GetFundNAVHistoriesRequest) (*dto.NAVHistories, error) {
	query := url.Values{}
	setOptional(query, "period", request.Period)
	path := "/fund-trading/public/fund-certificates/" + escaped(request.FundName) + "/nav-histories"
	return do[dto.NAVHistories](s, ctx, CapabilityGetFundNAVHistories, "GET", path, query)
}

func (s *service) GetFundPortfolio(ctx context.Context, request dto.GetFundPortfolioRequest) (*dto.PortfolioFund, error) {
	query := url.Values{}
	setOptional(query, "month", request.Month)
	path := "/market/funds/" + escaped(request.Fund) + "/portfolio"
	return do[dto.PortfolioFund](s, ctx, CapabilityGetFundPortfolio, "GET", path, query)
}

func (s *service) GetFundCertificates(ctx context.Context, request dto.GetFundCertificatesRequest) (*[]dto.FundCertificate, error) {
	query := url.Values{}
	set(query, "fund-type", request.FundType)
	if request.FundCompanyID != nil {
		set(query, "fund-company-id", strconv.FormatInt(*request.FundCompanyID, 10))
	}
	return do[[]dto.FundCertificate](s, ctx, CapabilityGetFundCertificates, "GET", "/fund-trading/public/fund-certificates", query)
}

func (s *service) GetGlobalNewsList(ctx context.Context, request dto.GetGlobalNewsListRequest) (*dto.GlobalNewsPage, error) {
	query := url.Values{}
	setOptional(query, "category", request.Category)
	setOptionalInt(query, "page", request.Page)
	setOptionalInt(query, "page_size", request.PageSize)
	return do[dto.GlobalNewsPage](s, ctx, CapabilityGetGlobalNewsList, "GET", "/market/financial-data/global-news", query)
}

func (s *service) GetGlobalNewsDetail(ctx context.Context, request dto.GetGlobalNewsDetailRequest) (*dto.GlobalNewsDetail, error) {
	path := "/market/financial-data/global-news/" + strconv.FormatInt(request.ID, 10)
	return do[dto.GlobalNewsDetail](s, ctx, CapabilityGetGlobalNewsDetail, "GET", path, url.Values{})
}

func (s *service) GetGoldChartData(ctx context.Context, request dto.GetChartDataRequest) (*dto.ChartDataResponse, error) {
	return getChartData(s, ctx, CapabilityGetGoldChartData, "/market/financial-data/gold-chart", request)
}

func (s *service) GetGoldProviderData(ctx context.Context) (*[]dto.CommodityIndexData, error) {
	return do[[]dto.CommodityIndexData](s, ctx, CapabilityGetGoldProviderData, "GET", "/market/financial-data/gold-providers", url.Values{})
}

func (s *service) GetGoldData(ctx context.Context) (*[]dto.CommodityIndexData, error) {
	return do[[]dto.CommodityIndexData](s, ctx, CapabilityGetGoldData, "GET", "/market/financial-data/gold", url.Values{})
}

func (s *service) GetIndexRealtime(ctx context.Context, request dto.GetIndexRealtimeRequest) (*[]dto.IndexRealtime, error) {
	query := url.Values{}
	set(query, "index", request.Index)
	return do[[]dto.IndexRealtime](s, ctx, CapabilityGetIndexRealtime, "GET", "/market/index-realtime", query)
}

func (s *service) GetMacroData(ctx context.Context, request dto.GetMacroDataRequest) (*[]dto.MacroData, error) {
	query := url.Values{}
	set(query, "type", request.Type)
	set(query, "country", request.Country)
	setOptional(query, "period", request.Period)
	return do[[]dto.MacroData](s, ctx, CapabilityGetMacroData, "GET", "/market/financial-data/macro", query)
}

func (s *service) GetMarketData(ctx context.Context, request dto.GetMarketDataRequest) (*[]dto.MarketData, error) {
	query := url.Values{}
	set(query, "type", request.Type)
	setOptionalInt(query, "limit", request.Limit)
	return do[[]dto.MarketData](s, ctx, CapabilityGetMarketData, "GET", "/market/financial-data/market", query)
}

func (s *service) GetMetalProviderData(ctx context.Context) (*[]dto.CommodityIndexData, error) {
	return do[[]dto.CommodityIndexData](s, ctx, CapabilityGetMetalProviderData, "GET", "/market/financial-data/metal-providers", url.Values{})
}

func (s *service) GetStockEvents(ctx context.Context, request dto.GetStockEventsRequest) (*[]dto.StockEventResponse, error) {
	query := url.Values{}
	setOptional(query, "stock", request.Stock)
	setOptional(query, "stocks", request.Stocks)
	setOptional(query, "from_date", request.FromDate)
	setOptional(query, "to_date", request.ToDate)
	return do[[]dto.StockEventResponse](s, ctx, CapabilityGetStockEvents, "GET", "/market/news", query)
}

func (s *service) GetPriceHistoriesChart(ctx context.Context, request dto.GetPriceHistoriesChartRequest) (*dto.PriceHistoriesChart, error) {
	query := url.Values{}
	set(query, "symbol", request.Symbol)
	set(query, "resolution", request.Resolution)
	set(query, "from", strconv.FormatInt(request.From, 10))
	set(query, "to", strconv.FormatInt(request.To, 10))
	return do[dto.PriceHistoriesChart](s, ctx, CapabilityGetPriceHistoriesChart, "GET", "/market/price-histories-chart", query)
}

func (s *service) GetRecommendationReports(ctx context.Context, request dto.GetRecommendationReportsRequest) (*dto.RecommendationReportsResponse, error) {
	path := "/market/recommendation-reports/" + escaped(request.Symbol)
	return do[dto.RecommendationReportsResponse](s, ctx, CapabilityGetRecommendationReports, "GET", path, url.Values{})
}

func (s *service) GetSilverChartData(ctx context.Context, request dto.GetChartDataRequest) (*dto.ChartDataResponse, error) {
	return getChartData(s, ctx, CapabilityGetSilverChartData, "/market/financial-data/silver-chart", request)
}

func (s *service) GetSilverData(ctx context.Context) (*[]dto.CommodityIndexData, error) {
	return do[[]dto.CommodityIndexData](s, ctx, CapabilityGetSilverData, "GET", "/market/financial-data/silver", url.Values{})
}

func (s *service) GetStockRealtime(ctx context.Context, request dto.GetStockRealtimeRequest) (*dto.StockRealtimeResponse, error) {
	query := url.Values{}
	setOptional(query, "symbol", request.Symbol)
	setOptional(query, "symbols", request.Symbols)
	setOptional(query, "exchange", request.Exchange)
	return do[dto.StockRealtimeResponse](s, ctx, CapabilityGetStockRealtime, "GET", "/market/stock-realtime", query)
}

func (s *service) GetTradingEconomicsData(ctx context.Context, request dto.GetTradingEconomicsDataRequest) (*[]dto.TradingEconomicsData, error) {
	query := url.Values{}
	set(query, "country", request.Country)
	setOptional(query, "category", request.Category)
	setOptionalInt(query, "year", request.Year)
	return do[[]dto.TradingEconomicsData](s, ctx, CapabilityGetTradingEconomicsData, "GET", "/market/financial-data/trading-economics", query)
}

func getChartData(s *service, ctx context.Context, capability core.Capability, path string, request dto.GetChartDataRequest) (*dto.ChartDataResponse, error) {
	query := url.Values{}
	setOptionalInt(query, "days", request.Days)
	return do[dto.ChartDataResponse](s, ctx, capability, "GET", path, query)
}

func escaped(value string) string { return url.PathEscape(value) }

func set(query url.Values, key, value string) { query.Set(key, value) }

func setOptional(query url.Values, key, value string) {
	if value != "" {
		query.Set(key, value)
	}
}

func setOptionalInt(query url.Values, key string, value int64) {
	if value > 0 {
		query.Set(key, strconv.FormatInt(value, 10))
	}
}

func stringValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func firstNonEmpty(values ...*string) string {
	for _, value := range values {
		if value != nil && strings.TrimSpace(*value) != "" {
			return *value
		}
	}
	return ""
}
