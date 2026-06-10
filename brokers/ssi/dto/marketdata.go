package dto

type SecuritiesRequest struct {
	Market    string `json:"market"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}

type SecuritiesResponse struct {
	Data        []Security `json:"data"`
	Message     string     `json:"message"`
	Status      string     `json:"status"`
	TotalRecord int        `json:"totalRecord"`
}

type Security struct {
	Market      string `json:"Market"`
	Symbol      string `json:"Symbol"`
	StockName   string `json:"StockName"`
	StockEnName string `json:"StockEnName"`
}

type SecuritiesDetailsRequest struct {
	Market    string `json:"market"`
	Symbol    string `json:"symbol"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}

type SecuritiesDetailsResponse struct {
	Data        []SecurityDetailsGroup `json:"data"`
	Message     string                 `json:"message"`
	Status      string                 `json:"status"`
	TotalRecord int                    `json:"totalRecord"`
}

type SecurityDetailsGroup struct {
	RType        string           `json:"RType"`
	ReportDate   string           `json:"ReportDate"`
	TotalNoSym   string           `json:"TotalNoSym"`
	RepeatedInfo []SecurityDetail `json:"RepeatedInfo"`
}

type SecurityDetail struct {
	Isin               any    `json:"Isin"`
	Symbol             string `json:"Symbol"`
	SymbolName         string `json:"SymbolName"`
	SymbolEngName      string `json:"SymbolEngName"`
	SecType            string `json:"SecType"`
	MarketID           string `json:"MarketId"`
	Exchange           string `json:"Exchange"`
	Issuer             any    `json:"Issuer"`
	LotSize            string `json:"LotSize"`
	IssueDate          string `json:"IssueDate"`
	MaturityDate       string `json:"MaturityDate"`
	FirstTradingDate   string `json:"FirstTradingDate"`
	LastTradingDate    string `json:"LastTradingDate"`
	ContractMultiplier string `json:"ContractMultiplier"`
	SettlMethod        string `json:"SettlMethod"`
	Underlying         any    `json:"Underlying"`
	PutOrCall          any    `json:"PutOrCall"`
	ExercisePrice      string `json:"ExercisePrice"`
	ExerciseStyle      string `json:"ExerciseStyle"`
	ExcerciseRatio     string `json:"ExcerciseRatio"`
	ListedShare        string `json:"ListedShare"`
	TickPrice1         any    `json:"TickPrice1"`
	TickIncrement1     any    `json:"TickIncrement1"`
	TickPrice2         any    `json:"TickPrice2"`
	TickIncrement2     any    `json:"TickIncrement2"`
	TickPrice3         any    `json:"TickPrice3"`
	TickIncrement3     any    `json:"TickIncrement3"`
	TickPrice4         any    `json:"TickPrice4"`
	TickIncrement4     any    `json:"TickIncrement4"`
}

type IndexComponentsRequest struct {
	IndexCode string `json:"indexCode"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}

type IndexComponentsResponse struct {
	Data        []IndexComponentGroup `json:"data"`
	Message     string                `json:"message"`
	Status      string                `json:"status"`
	TotalRecord int                   `json:"totalRecord"`
}

type IndexComponentGroup struct {
	IndexCode      string               `json:"IndexCode"`
	IndexName      string               `json:"IndexName"`
	Exchange       string               `json:"Exchange"`
	TotalSymbolNo  string               `json:"TotalSymbolNo"`
	IndexComponent []IndexComponentItem `json:"IndexComponent"`
}

type IndexComponentItem struct {
	Isin        string `json:"Isin"`
	StockSymbol string `json:"StockSymbol"`
}

type IndexListRequest struct {
	Exchange  string `json:"exchange"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}

type IndexListResponse struct {
	Data        []IndexItem `json:"data"`
	Message     string      `json:"message"`
	Status      string      `json:"status"`
	TotalRecord int         `json:"totalRecord"`
}

type IndexItem struct {
	IndexCode string `json:"IndexCode"`
	IndexName string `json:"IndexName"`
	Exchange  string `json:"Exchange"`
}

type DailyOhlcRequest struct {
	Symbol    string `json:"symbol"`
	FromDate  string `json:"fromDate"`
	ToDate    string `json:"toDate"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
	Ascending bool   `json:"ascending"`
}

type DailyOhlcResponse struct {
	Data        []OhlcRecord `json:"data"`
	Message     string       `json:"message"`
	Status      string       `json:"status"`
	TotalRecord int          `json:"totalRecord"`
}

type IntradayOhlcRequest struct {
	Symbol     string `json:"symbol"`
	FromDate   string `json:"fromDate"`
	ToDate     string `json:"toDate"`
	PageIndex  int    `json:"pageIndex"`
	PageSize   int    `json:"pageSize"`
	Ascending  bool   `json:"ascending"`
	Resolution int    `json:"resolution"`
}

type IntradayOhlcResponse struct {
	Data        []OhlcRecord `json:"data"`
	Message     string       `json:"message"`
	Status      string       `json:"status"`
	TotalRecord int          `json:"totalRecord"`
}

type OhlcRecord struct {
	Symbol      string `json:"Symbol"`
	Market      string `json:"Market"`
	TradingDate string `json:"TradingDate"`
	Time        any    `json:"Time"`
	Open        string `json:"Open"`
	High        string `json:"High"`
	Low         string `json:"Low"`
	Close       string `json:"Close"`
	Volume      string `json:"Volume"`
	Value       string `json:"Value"`
}

type DailyIndexRequest struct {
	IndexID   string `json:"indexId"`
	FromDate  string `json:"fromDate"`
	ToDate    string `json:"toDate"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
	Ascending bool   `json:"ascending"`
}

type DailyIndexResponse struct {
	Data        []DailyIndexRecord `json:"data"`
	Message     string             `json:"message"`
	Status      string             `json:"status"`
	TotalRecord int                `json:"totalRecord"`
}

type DailyIndexRecord struct {
	IndexID        string `json:"IndexId"`
	IndexValue     string `json:"IndexValue"`
	TradingDate    string `json:"TradingDate"`
	Time           any    `json:"Time"`
	Change         string `json:"Change"`
	RatioChange    string `json:"RatioChange"`
	TotalTrade     string `json:"TotalTrade"`
	TotalMatchVol  string `json:"TotalMatchVol"`
	TotalMatchVal  string `json:"TotalMatchVal"`
	TypeIndex      any    `json:"TypeIndex"`
	IndexName      string `json:"IndexName"`
	Advances       string `json:"Advances"`
	NoChanges      string `json:"NoChanges"`
	Declines       string `json:"Declines"`
	Ceilings       string `json:"Ceilings"`
	Floors         string `json:"Floors"`
	TotalDealVol   string `json:"TotalDealVol"`
	TotalDealVal   string `json:"TotalDealVal"`
	TotalVol       string `json:"TotalVol"`
	TotalVal       string `json:"TotalVal"`
	TradingSession string `json:"TradingSession"`
	Market         string `json:"Market"`
	Exchange       string `json:"Exchange"`
}

type DailyStockPriceRequest struct {
	Symbol    string `json:"symbol"`
	FromDate  string `json:"fromDate"`
	ToDate    string `json:"toDate"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
	Market    string `json:"market"`
}

type DailyStockPriceResponse struct {
	Data        []DailyStockPriceRecord `json:"data"`
	Message     string                  `json:"message"`
	Status      string                  `json:"status"`
	TotalRecord int                     `json:"totalRecord"`
}

type DailyStockPriceRecord struct {
	TradingDate         string `json:"TradingDate"`
	PriceChange         string `json:"PriceChange"`
	PerPriceChange      string `json:"PerPriceChange"`
	CeilingPrice        string `json:"CeilingPrice"`
	FloorPrice          string `json:"FloorPrice"`
	RefPrice            string `json:"RefPrice"`
	OpenPrice           string `json:"OpenPrice"`
	HighestPrice        string `json:"HighestPrice"`
	LowestPrice         string `json:"LowestPrice"`
	ClosePrice          string `json:"ClosePrice"`
	AveragePrice        string `json:"AveragePrice"`
	ClosePriceAdjusted  string `json:"ClosePriceAdjusted"`
	TotalMatchVol       string `json:"TotalMatchVol"`
	TotalMatchVal       string `json:"TotalMatchVal"`
	TotalDealVal        string `json:"TotalDealVal"`
	TotalDealVol        string `json:"TotalDealVol"`
	ForeignBuyVolTotal  string `json:"ForeignBuyVolTotal"`
	ForeignCurrentRoom  string `json:"ForeignCurrentRoom"`
	ForeignSellVolTotal string `json:"ForeignSellVolTotal"`
	ForeignBuyValTotal  string `json:"ForeignBuyValTotal"`
	ForeignSellValTotal string `json:"ForeignSellValTotal"`
	TotalBuyTrade       string `json:"TotalBuyTrade"`
	TotalBuyTradeVol    string `json:"TotalBuyTradeVol"`
	TotalSellTrade      string `json:"TotalSellTrade"`
	TotalSellTradeVol   string `json:"TotalSellTradeVol"`
	NetBuySellVol       string `json:"NetBuySellVol"`
	NetBuySellVal       string `json:"NetBuySellVal"`
	TotalTradedVol      string `json:"TotalTradedVol"`
	TotalTradedValue    string `json:"TotalTradedValue"`
	Symbol              string `json:"Symbol"`
	Time                any    `json:"Time"`
}
