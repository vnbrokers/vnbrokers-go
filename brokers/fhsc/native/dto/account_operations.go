package dto

type GetAccountSummaryRequest struct{ SubAccountID string }

type GetUserAssetsSummaryRequest struct {
	UserID       int64
	CacheControl string
}

type GetPnLTodayRequest struct {
	UserID       int64
	SubAccountID string
}

type GetPortfolioRequest struct {
	SubAccountID string
	CacheControl string
}

type GetAvailableTradeRequest struct {
	SubAccountID string
	OrderSide    string
	Symbol       string
	QuotePrice   int64
}

type GetUserRightsRequest struct {
	SubAccountID string
	FromDate     string
	ToDate       string
	CatType      string
	IsCom        string
	Symbol       string
	Status       string
}

type GetMarketSessionRequest struct{ Exchange string }
