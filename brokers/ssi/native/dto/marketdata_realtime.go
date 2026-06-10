package dto

import "github.com/vnbrokers/vnbrokers-go/internal/numberic"

type QuoteEvent struct {
	TradingDate string `json:"TradingDate"`
	Time        string `json:"Time"`
	Exchange    string `json:"Exchange"`
	Symbol      string `json:"Symbol"`
	RType       string `json:"RType"`

	AskPrice1  float64 `json:"AskPrice1"`
	AskPrice2  float64 `json:"AskPrice2"`
	AskPrice3  float64 `json:"AskPrice3"`
	AskPrice4  float64 `json:"AskPrice4"`
	AskPrice5  float64 `json:"AskPrice5"`
	AskPrice6  float64 `json:"AskPrice6"`
	AskPrice7  float64 `json:"AskPrice7"`
	AskPrice8  float64 `json:"AskPrice8"`
	AskPrice9  float64 `json:"AskPrice9"`
	AskPrice10 float64 `json:"AskPrice10"`

	AskVol1  float64 `json:"AskVol1"`
	AskVol2  float64 `json:"AskVol2"`
	AskVol3  float64 `json:"AskVol3"`
	AskVol4  float64 `json:"AskVol4"`
	AskVol5  float64 `json:"AskVol5"`
	AskVol6  float64 `json:"AskVol6"`
	AskVol7  float64 `json:"AskVol7"`
	AskVol8  float64 `json:"AskVol8"`
	AskVol9  float64 `json:"AskVol9"`
	AskVol10 float64 `json:"AskVol10"`

	BidPrice1  float64 `json:"BidPrice1"`
	BidPrice2  float64 `json:"BidPrice2"`
	BidPrice3  float64 `json:"BidPrice3"`
	BidPrice4  float64 `json:"BidPrice4"`
	BidPrice5  float64 `json:"BidPrice5"`
	BidPrice6  float64 `json:"BidPrice6"`
	BidPrice7  float64 `json:"BidPrice7"`
	BidPrice8  float64 `json:"BidPrice8"`
	BidPrice9  float64 `json:"BidPrice9"`
	BidPrice10 float64 `json:"BidPrice10"`

	BidVol1  float64 `json:"BidVol1"`
	BidVol2  float64 `json:"BidVol2"`
	BidVol3  float64 `json:"BidVol3"`
	BidVol4  float64 `json:"BidVol4"`
	BidVol5  float64 `json:"BidVol5"`
	BidVol6  float64 `json:"BidVol6"`
	BidVol7  float64 `json:"BidVol7"`
	BidVol8  float64 `json:"BidVol8"`
	BidVol9  float64 `json:"BidVol9"`
	BidVol10 float64 `json:"BidVol10"`

	TradingSession string `json:"TradingSession"`
}

type ForeignRoomEvent struct {
	RType       string  `json:"RType"`
	TradingDate string  `json:"TradingDate"`
	Time        string  `json:"Time"`
	Isin        string  `json:"Isin"`
	Symbol      string  `json:"Symbol"`
	TotalRoom   float64 `json:"TotalRoom"`
	CurrentRoom float64 `json:"CurrentRoom"`
	BuyVol      float64 `json:"BuyVol"`
	SellVol     float64 `json:"SellVol"`
	BuyVal      float64 `json:"BuyVal"`
	SellVal     float64 `json:"SellVal"`
	MarketID    string  `json:"MarketId"`
	Exchange    string  `json:"Exchange"`
}

type TradingStatusEvent struct {
	RType          string `json:"RType"`
	MarketID       string `json:"MarketId"`
	TradingDate    string `json:"TradingDate"`
	Time           string `json:"Time"`
	Symbol         string `json:"Symbol"`
	TradingSession string `json:"TradingSession"`
	TradingStatus  string `json:"TradingStatus"`
	Exchange       string `json:"Exchange"`
}

type TradeEvent struct {
	RType           string  `json:"RType"`
	TradingDate     string  `json:"TradingDate"`
	Time            string  `json:"Time"`
	Isin            string  `json:"Isin"`
	Symbol          string  `json:"Symbol"`
	Ceiling         float64 `json:"Ceiling"`
	Floor           float64 `json:"Floor"`
	RefPrice        float64 `json:"RefPrice"`
	AvgPrice        float64 `json:"AvgPrice"`
	PriorVal        float64 `json:"PriorVal"`
	LastPrice       float64 `json:"LastPrice"`
	LastVol         float64 `json:"LastVol"`
	TotalVal        float64 `json:"TotalVal"`
	TotalVol        float64 `json:"TotalVol"`
	MarketID        string  `json:"MarketId"`
	Exchange        string  `json:"Exchange"`
	TradingSession  string  `json:"TradingSession"`
	TradingStatus   string  `json:"TradingStatus"`
	Change          float64 `json:"Change"`
	RatioChange     float64 `json:"RatioChange"`
	EstMatchedPrice float64 `json:"EstMatchedPrice"`
	Highest         float64 `json:"Highest"`
	Lowest          float64 `json:"Lowest"`
	Side            string  `json:"Side"`
}

type SnapshotEvent struct {
	RType       string `json:"RType"`
	TradingDate string `json:"TradingDate"`
	Time        string `json:"Time"`
	Isin        string `json:"Isin"`
	Symbol      string `json:"Symbol"`

	Ceiling   float64             `json:"Ceiling"`
	Floor     float64             `json:"Floor"`
	RefPrice  float64             `json:"RefPrice"`
	Open      float64             `json:"Open"`
	High      float64             `json:"High"`
	Low       float64             `json:"Low"`
	Close     float64             `json:"Close"`
	AvgPrice  numberic.NaNFloat64 `json:"AvgPrice"`
	PriorVal  float64             `json:"PriorVal"`
	LastPrice float64             `json:"LastPrice"`
	LastVol   float64             `json:"LastVol"`
	TotalVal  float64             `json:"TotalVal"`
	TotalVol  float64             `json:"TotalVol"`

	BidPrice1  float64 `json:"BidPrice1"`
	BidPrice2  float64 `json:"BidPrice2"`
	BidPrice3  float64 `json:"BidPrice3"`
	BidPrice4  float64 `json:"BidPrice4"`
	BidPrice5  float64 `json:"BidPrice5"`
	BidPrice6  float64 `json:"BidPrice6"`
	BidPrice7  float64 `json:"BidPrice7"`
	BidPrice8  float64 `json:"BidPrice8"`
	BidPrice9  float64 `json:"BidPrice9"`
	BidPrice10 float64 `json:"BidPrice10"`

	BidVol1  float64 `json:"BidVol1"`
	BidVol2  float64 `json:"BidVol2"`
	BidVol3  float64 `json:"BidVol3"`
	BidVol4  float64 `json:"BidVol4"`
	BidVol5  float64 `json:"BidVol5"`
	BidVol6  float64 `json:"BidVol6"`
	BidVol7  float64 `json:"BidVol7"`
	BidVol8  float64 `json:"BidVol8"`
	BidVol9  float64 `json:"BidVol9"`
	BidVol10 float64 `json:"BidVol10"`

	AskPrice1  float64 `json:"AskPrice1"`
	AskPrice2  float64 `json:"AskPrice2"`
	AskPrice3  float64 `json:"AskPrice3"`
	AskPrice4  float64 `json:"AskPrice4"`
	AskPrice5  float64 `json:"AskPrice5"`
	AskPrice6  float64 `json:"AskPrice6"`
	AskPrice7  float64 `json:"AskPrice7"`
	AskPrice8  float64 `json:"AskPrice8"`
	AskPrice9  float64 `json:"AskPrice9"`
	AskPrice10 float64 `json:"AskPrice10"`

	AskVol1  float64 `json:"AskVol1"`
	AskVol2  float64 `json:"AskVol2"`
	AskVol3  float64 `json:"AskVol3"`
	AskVol4  float64 `json:"AskVol4"`
	AskVol5  float64 `json:"AskVol5"`
	AskVol6  float64 `json:"AskVol6"`
	AskVol7  float64 `json:"AskVol7"`
	AskVol8  float64 `json:"AskVol8"`
	AskVol9  float64 `json:"AskVol9"`
	AskVol10 float64 `json:"AskVol10"`

	MarketID        string  `json:"MarketId"`
	Exchange        string  `json:"Exchange"`
	TradingSession  string  `json:"TradingSession"`
	TradingStatus   string  `json:"TradingStatus"`
	Change          float64 `json:"Change"`
	RatioChange     float64 `json:"RatioChange"`
	EstMatchedPrice float64 `json:"EstMatchedPrice"`
	Side            string  `json:"Side"`
	CloseQtty       float64 `json:"CloseQtty"`
}

type MarketIndexEvent struct {
	IndexID         string  `json:"IndexId"`
	IndexValEst     float64 `json:"IndexValEst"`
	IndexValue      float64 `json:"IndexValue"`
	PriorIndexValue float64 `json:"PriorIndexValue"`
	TradingDate     string  `json:"TradingDate"`
	Time            string  `json:"Time"`
	TotalTrade      float64 `json:"TotalTrade"`
	TotalQtty       float64 `json:"TotalQtty"`
	TotalValue      float64 `json:"TotalValue"`
	IndexName       string  `json:"IndexName"`
	Advances        float64 `json:"Advances"`
	NoChanges       float64 `json:"NoChanges"`
	Declines        float64 `json:"Declines"`
	Ceilings        float64 `json:"Ceilings"`
	Floors          float64 `json:"Floors"`
	Change          float64 `json:"Change"`
	RatioChange     float64 `json:"RatioChange"`
	TotalQttyPt     float64 `json:"TotalQttyPt"`
	TotalValuePt    float64 `json:"TotalValuePt"`
	Exchange        string  `json:"Exchange"`
	AllQty          float64 `json:"AllQty"`
	AllValue        float64 `json:"AllValue"`
	IndexType       string  `json:"IndexType"`
	TradingSession  *string `json:"TradingSession"`
	MarketID        *string `json:"MarketId"`
	RType           string  `json:"RType"`
	TotalQttyOd     float64 `json:"TotalQttyOd"`
	TotalValueOd    float64 `json:"TotalValueOd"`
}

type OHLCVEvent struct {
	RType       string  `json:"RType"`
	Symbol      string  `json:"Symbol"`
	TradingDate string  `json:"TradingDate"`
	Time        string  `json:"Time"`
	Open        float64 `json:"Open"`
	High        float64 `json:"High"`
	Low         float64 `json:"Low"`
	Close       float64 `json:"Close"`
	Volume      float64 `json:"Volume"`
	Value       float64 `json:"Value"`
}

type OddLotEvent struct {
	RType          string  `json:"RType"`
	TradingDate    string  `json:"TradingDate"`
	Time           string  `json:"Time"`
	StockNo        float64 `json:"StockNo"`
	Symbol         string  `json:"Symbol"`
	Ceiling        float64 `json:"Ceiling"`
	Floor          float64 `json:"Floor"`
	RefPrice       float64 `json:"RefPrice"`
	Open           float64 `json:"Open"`
	High           float64 `json:"High"`
	Low            float64 `json:"Low"`
	LastPrice      float64 `json:"LastPrice"`
	LastVol        float64 `json:"LastVol"`
	TotalVal       float64 `json:"TotalVal"`
	TotalVol       float64 `json:"TotalVol"`
	BidPrice1      float64 `json:"BidPrice1"`
	BidPrice2      float64 `json:"BidPrice2"`
	BidPrice3      float64 `json:"BidPrice3"`
	BidVol1        float64 `json:"BidVol1"`
	BidVol2        float64 `json:"BidVol2"`
	BidVol3        float64 `json:"BidVol3"`
	AskPrice1      float64 `json:"AskPrice1"`
	AskPrice2      float64 `json:"AskPrice2"`
	AskPrice3      float64 `json:"AskPrice3"`
	AskVol1        float64 `json:"AskVol1"`
	AskVol2        float64 `json:"AskVol2"`
	AskVol3        float64 `json:"AskVol3"`
	Exchange       string  `json:"Exchange"`
	TradingSession string  `json:"TradingSession"`
	TradingStatus  string  `json:"TradingStatus"`
	Change         float64 `json:"Change"`
	RatioChange    float64 `json:"RatioChange"`
	StockType      string  `json:"StockType"`
}
