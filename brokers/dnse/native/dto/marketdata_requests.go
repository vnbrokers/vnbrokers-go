package dto

type GetTradeHistoryRequest struct {
	Symbol, BoardID string
	From, To        int64
	Limit           int
}
type GetInstrumentDetailsRequest struct {
	Symbol, MarketID, SecurityGroupID, IndexName string
	Limit, Page                                  int
}
type GetInstrumentsRequest struct {
	Symbol      string
	Limit, Page int
}
type GetLatestQuotesRequest struct{ Symbol, BoardID string }
type GetLatestTradesRequest struct{ Symbol, BoardID string }
type GetOHLCRequest struct {
	Symbol, Type, Resolution string
	From, To                 int64
}
type GetClosePriceRequest struct{ Symbol, BoardID string }
type GetQuoteHistoryRequest struct {
	Symbol, BoardID string
	From, To        int64
	Limit           int
}
type GetSecurityDefinitionRequest struct{ Symbol, BoardID string }
type GetWorkingDatesRequest struct{}
