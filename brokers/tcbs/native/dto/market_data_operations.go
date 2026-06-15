package dto

type GetDerivativeTickersRequest struct{}
type GetDerivativeTickersResponse struct {
	Data []DerivativeTicker `json:"data"`
}

type GetStockTickersRequest struct {
	Tickers string
	Index   float64
}
type GetStockTickersResponse struct {
	Data        []StockTicker `json:"data"`
	TradingDate string        `json:"tradingDate"`
}

type GetStockForeignRoomsRequest struct{ Index float64 }
type GetStockForeignRoomsResponse struct {
	Data        []StockForeignRoom `json:"data"`
	TradingDate string             `json:"tradingDate"`
}

type GetStockPutThroughsRequest struct{ Floor float64 }
type GetStockPutThroughsResponse struct {
	BuyAdvertisements  []PutThroughAdvertisement `json:"buyAdv"`
	SellAdvertisements []PutThroughAdvertisement `json:"sellAdv"`
	Matches            []PutThroughMatch         `json:"match"`
}

type GetStockTradeHistoryRequest struct {
	Ticker    string
	Page      float64
	Size      float64
	HeadIndex float64
}
type GetStockTradeHistoryResponse struct {
	Page          float64                  `json:"page"`
	Size          float64                  `json:"size"`
	HeadIndex     float64                  `json:"headIndex"`
	NumberOfItems float64                  `json:"numberOfItems"`
	Total         float64                  `json:"total"`
	Ticker        string                   `json:"ticker"`
	Date          string                   `json:"d"`
	Data          []StockTradeHistoryEntry `json:"data"`
}

type GetStockSupplyDemand15MinutesRequest struct {
	Ticker     string
	TimeWindow string
	TWindow    string
	Type       string
}
type GetStockSupplyDemand15MinutesResponse struct {
	Ticker string                            `json:"ticker"`
	Data   []StockSupplyDemand15MinutesEntry `json:"data"`
}

type GetStockSupplyDemandDailyRequest struct {
	Ticker string
	Type   string
}
type GetStockSupplyDemandDailyResponse struct {
	Ticker string                   `json:"ticker"`
	Data   []StockSupplyDemandEntry `json:"data"`
}

type GetStockSupplyDemandMonthlyRequest struct {
	Ticker     string
	TimeWindow string
	Type       string
}
type GetStockSupplyDemandMonthlyResponse struct {
	Ticker string                   `json:"ticker"`
	Data   []StockSupplyDemandEntry `json:"data"`
}
