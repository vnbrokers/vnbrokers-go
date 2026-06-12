package dto

import (
	"encoding/json"
	"testing"

	"github.com/shopspring/decimal"
)

func TestQuoteEventDecodesCurrentWireFields(t *testing.T) {
	var event QuoteEvent
	if err := json.Unmarshal([]byte(`{
		"T":"q",
		"bid":[
			{"price":23.35,"qtty":38550},
			{"price":23.3,"qtty":73910},
			{"price":23.25,"qtty":44840}
		],
		"boardId":"G1",
		"isin":"VN000000HPG4",
		"marketId":"STO",
		"multicastReceiveTime":{"Nanos":197458188,"Seconds":1781238547},
		"offer":[
			{"price":23.4,"qtty":7040},
			{"price":23.45,"qtty":5280},
			{"price":23.5,"qtty":15710}
		],
		"symbol":"HPG",
		"time":{"Nanos":197000000,"Seconds":1781238547}
	}`), &event); err != nil {
		t.Fatal(err)
	}

	if event.T != "q" || event.MarketID != "STO" || event.BoardID != "G1" || event.ISIN != "VN000000HPG4" || event.Symbol != "HPG" {
		t.Fatalf("event = %+v", event)
	}
	if event.MulticastReceiveTime.Nanos != 197458188 || event.MulticastReceiveTime.Seconds != 1781238547 {
		t.Fatalf("multicast receive time = %+v", event.MulticastReceiveTime)
	}
	if event.Time.Nanos != 197000000 || event.Time.Seconds != 1781238547 {
		t.Fatalf("time = %+v", event.Time)
	}
	if len(event.Bid) != 3 || event.Bid[0].Price == nil || event.Bid[0].Price.String() != "23.35" || event.Bid[0].Quantity == nil || event.Bid[0].Quantity.String() != "38550" {
		t.Fatalf("bid = %+v", event.Bid)
	}
	if len(event.Offer) != 3 || event.Offer[0].Price == nil || event.Offer[0].Price.String() != "23.4" || event.Offer[0].Quantity == nil || event.Offer[0].Quantity.String() != "7040" {
		t.Fatalf("offer = %+v", event.Offer)
	}
}

func TestTradeEventDecodesCurrentWireFields(t *testing.T) {
	var event TradeEvent
	if err := json.Unmarshal([]byte(`{
		"T":"t",
		"boardId":"G1",
		"grossTradeAmount":585.321425,
		"highestPrice":16.6,
		"isin":"VN000000TPB0",
		"lowestPrice":16,
		"marketId":"STO",
		"matchPrice":16.4,
		"matchQtty":100,
		"multicastReceiveTime":{"Nanos":338028896,"Seconds":1781244600},
		"openPrice":16.05,
		"symbol":"TPB",
		"time":{"Nanos":338000000,"Seconds":1781244600},
		"totalVolumeTraded":3586250,
		"tradingSessionId":"40"
	}`), &event); err != nil {
		t.Fatal(err)
	}

	if event.T != "t" || event.MarketID != "STO" || event.BoardID != "G1" || event.ISIN != "VN000000TPB0" || event.Symbol != "TPB" {
		t.Fatalf("event = %+v", event)
	}
	if event.TradingSessionID != "40" {
		t.Fatalf("trading session id = %q", event.TradingSessionID)
	}
	if event.MulticastReceiveTime.Nanos != 338028896 || event.MulticastReceiveTime.Seconds != 1781244600 {
		t.Fatalf("multicast receive time = %+v", event.MulticastReceiveTime)
	}
	if event.Time.Nanos != 338000000 || event.Time.Seconds != 1781244600 {
		t.Fatalf("time = %+v", event.Time)
	}
	assertDecimalString(t, "match price", event.MatchPrice, "16.4")
	assertDecimalString(t, "match quantity", event.MatchQuantity, "100")
	assertDecimalString(t, "gross trade amount", event.GrossTradeAmount, "585.321425")
	assertDecimalString(t, "highest price", event.HighestPrice, "16.6")
	assertDecimalString(t, "lowest price", event.LowestPrice, "16")
	assertDecimalString(t, "open price", event.OpenPrice, "16.05")
	assertDecimalString(t, "total volume traded", event.TotalVolumeTraded, "3586250")
}

func TestTradeExtraEventDecodesCurrentWireFields(t *testing.T) {
	var event TradeExtraEvent
	if err := json.Unmarshal([]byte(`{
		"T":"te",
		"avgPrice":23.459,
		"boardId":"G1",
		"grossTradeAmount":107.168245,
		"highestPrice":23.55,
		"isin":"VN000000HPG4",
		"lowestPrice":23.35,
		"marketId":"STO",
		"matchPrice":23.35,
		"matchQtty":20,
		"multicastReceiveTime":{"Nanos":924429049,"Seconds":1781238485},
		"openPrice":23.55,
		"side":"SELL",
		"symbol":"HPG",
		"time":{"Nanos":924000000,"Seconds":1781238485},
		"totalVolumeTraded":456830,
		"tradingSessionId":"40"
	}`), &event); err != nil {
		t.Fatal(err)
	}

	if event.T != "te" || event.MarketID != "STO" || event.BoardID != "G1" || event.ISIN != "VN000000HPG4" || event.Symbol != "HPG" {
		t.Fatalf("event = %+v", event)
	}
	if event.Side != "SELL" || event.TradingSessionID != "40" {
		t.Fatalf("side/session = %q/%q", event.Side, event.TradingSessionID)
	}
	if event.MulticastReceiveTime.Nanos != 924429049 || event.MulticastReceiveTime.Seconds != 1781238485 {
		t.Fatalf("multicast receive time = %+v", event.MulticastReceiveTime)
	}
	if event.Time.Nanos != 924000000 || event.Time.Seconds != 1781238485 {
		t.Fatalf("time = %+v", event.Time)
	}
	assertDecimalString(t, "average price", event.AveragePrice, "23.459")
	assertDecimalString(t, "gross trade amount", event.GrossTradeAmount, "107.168245")
	assertDecimalString(t, "highest price", event.HighestPrice, "23.55")
	assertDecimalString(t, "lowest price", event.LowestPrice, "23.35")
	assertDecimalString(t, "match price", event.MatchPrice, "23.35")
	assertDecimalString(t, "match quantity", event.MatchQuantity, "20")
	assertDecimalString(t, "open price", event.OpenPrice, "23.55")
	assertDecimalString(t, "total volume traded", event.TotalVolumeTraded, "456830")
}

func TestExpectedPriceEventDecodesCurrentWireFields(t *testing.T) {
	var event ExpectedPriceEvent
	if err := json.Unmarshal([]byte(`{
		"T":"e",
		"marketId":"DVX",
		"boardId":"G1",
		"symbol":"41I1G1000",
		"isin":"VN41I1G10000",
		"closePrice":28.45,
		"expectedTradePrice":28.45,
		"expectedTradeQuantity":133780,
		"time":{"Seconds":1779694639,"Nanos":736000000}
	}`), &event); err != nil {
		t.Fatal(err)
	}

	if event.T != "e" || event.MarketID != "DVX" || event.BoardID != "G1" || event.Symbol != "41I1G1000" || event.ISIN != "VN41I1G10000" {
		t.Fatalf("event = %+v", event)
	}
	assertDecimalString(t, "close price", event.ClosePrice, "28.45")
	assertDecimalString(t, "expected trade price", event.ExpectedTradePrice, "28.45")
	if event.ExpectedTradeQuantity != 133780 {
		t.Fatalf("expected trade quantity = %d", event.ExpectedTradeQuantity)
	}
	if event.Time.Seconds != 1779694639 || event.Time.Nanos != 736000000 {
		t.Fatalf("time = %+v", event.Time)
	}
}

func TestOHLCEventDecodesCurrentWireFields(t *testing.T) {
	var event OHLCEvent
	if err := json.Unmarshal([]byte(`{
		"T":"b",
		"time":1757992500,
		"open":30.4,
		"high":30.4,
		"low":30.25,
		"close":30.3,
		"volume":1398200,
		"symbol":"HPG",
		"resolution":"15",
		"lastUpdated":1757993014,
		"type":"STOCK"
	}`), &event); err != nil {
		t.Fatal(err)
	}

	if event.T != "b" || event.Time != 1757992500 {
		t.Fatalf("type/time = %q/%d", event.T, event.Time)
	}
	assertDecimalString(t, "open", event.Open, "30.4")
	assertDecimalString(t, "high", event.High, "30.4")
	assertDecimalString(t, "low", event.Low, "30.25")
	assertDecimalString(t, "close", event.Close, "30.3")
	if event.Volume != 1398200 {
		t.Fatalf("volume = %d", event.Volume)
	}
	if event.Symbol != "HPG" || event.Resolution != "15" || event.LastUpdated != 1757993014 || event.Type != "STOCK" {
		t.Fatalf("event = %+v", event)
	}
}

func TestOHLCClosedEventDecodesCurrentWireFields(t *testing.T) {
	var event OHLCClosedEvent
	if err := json.Unmarshal([]byte(`{
		"T":"bc",
		"time":1757992500,
		"open":30.4,
		"high":30.4,
		"low":30.25,
		"close":30.3,
		"volume":1398200,
		"symbol":"HPG",
		"resolution":"15",
		"lastUpdated":1757993014,
		"type":"STOCK"
	}`), &event); err != nil {
		t.Fatal(err)
	}

	if event.T != "bc" || event.Time != 1757992500 || event.Volume != 1398200 {
		t.Fatalf("event = %+v", event)
	}
	if event.Symbol != "HPG" || event.Resolution != "15" || event.LastUpdated != 1757993014 || event.Type != "STOCK" {
		t.Fatalf("event = %+v", event)
	}
	assertDecimalString(t, "open", event.Open, "30.4")
	assertDecimalString(t, "high", event.High, "30.4")
	assertDecimalString(t, "low", event.Low, "30.25")
	assertDecimalString(t, "close", event.Close, "30.3")
}

func assertDecimalString(t *testing.T, name string, value *decimal.Decimal, want string) {
	t.Helper()
	if value == nil || value.String() != want {
		t.Fatalf("%s = %v", name, value)
	}
}

func TestForeignEventDecodesCurrentWireFields(t *testing.T) {
	var event ForeignEvent
	if err := json.Unmarshal([]byte(`{
		"T":"f",
		"boardId":"G1",
		"buyTradedAmount":1342640000,
		"buyVolume":88900,
		"foreignInvestorTypeCode":"10",
		"foreignerBuyPossibleQuantity":1141121020,
		"foreignerOrderLimitQuantity":856057160,
		"marketId":"STO",
		"multicastReceiveTime":{"Nanos":102356876,"Seconds":1781246780},
		"sellTradedAmount":20130175000,
		"sellVolume":1339100,
		"symbol":"TCH",
		"totalBuyTradedAmount":1342640000,
		"totalBuyVolume":88900,
		"totalSellTradedAmount":20131525000,
		"totalSellVolume":1339190,
		"tradingSessionId":"40",
		"transactTime":"064600004"
	}`), &event); err != nil {
		t.Fatal(err)
	}

	if event.T != "f" || event.BoardID != "G1" || event.MarketID != "STO" || event.Symbol != "TCH" {
		t.Fatalf("event = %+v", event)
	}
	if event.ForeignInvestorTypeCode != "10" || event.TradingSessionID != "40" || event.TransactTime != "064600004" {
		t.Fatalf("event = %+v", event)
	}
	if event.BuyVolume != 88900 || event.SellVolume != 1339100 || event.TotalBuyVolume != 88900 || event.TotalSellVolume != 1339190 {
		t.Fatalf("volumes = %d/%d/%d/%d", event.BuyVolume, event.SellVolume, event.TotalBuyVolume, event.TotalSellVolume)
	}
	if event.ForeignerBuyPossibleQuantity != 1141121020 || event.ForeignerOrderLimitQuantity != 856057160 {
		t.Fatalf("foreign quantities = %d/%d", event.ForeignerBuyPossibleQuantity, event.ForeignerOrderLimitQuantity)
	}
	if event.MulticastReceiveTime.Nanos != 102356876 || event.MulticastReceiveTime.Seconds != 1781246780 {
		t.Fatalf("multicast receive time = %+v", event.MulticastReceiveTime)
	}
	assertDecimalString(t, "buy traded amount", event.BuyTradedAmount, "1342640000")
	assertDecimalString(t, "sell traded amount", event.SellTradedAmount, "20130175000")
	assertDecimalString(t, "total buy traded amount", event.TotalBuyTradedAmount, "1342640000")
	assertDecimalString(t, "total sell traded amount", event.TotalSellTradedAmount, "20131525000")
}

func TestMarketIndexEventAllowsNullableCounts(t *testing.T) {
	var event MarketIndexEvent
	if err := json.Unmarshal([]byte(`{"T":"mi","indexName":"VNINDEX","fluctuationLowerLimitIssueCount":null,"fluctuationUpperLimitIssueCount":7}`), &event); err != nil {
		t.Fatal(err)
	}
	if event.FluctuationLowerLimitIssueCount != nil {
		t.Fatalf("lower-limit count = %v", event.FluctuationLowerLimitIssueCount)
	}
	if event.FluctuationUpperLimitIssueCount == nil || *event.FluctuationUpperLimitIssueCount != 7 {
		t.Fatalf("upper-limit count = %v", event.FluctuationUpperLimitIssueCount)
	}
}

func TestMarketIndexEventDecodesCurrentWireFields(t *testing.T) {
	var event MarketIndexEvent
	if err := json.Unmarshal([]byte(`{
		"T":"mi",
		"blkTrdAccTrdVal":41.69273115,
		"blkTrdAccTrdVol":1738007,
		"changedRatio":0.91,
		"changedValue":17.74,
		"contauctAccTrdVal":2369.54154935,
		"contauctAccTrdVol":85122929,
		"currencyCode":"VND",
		"grossTradeAmount":2411.2342805,
		"highestValueIndexes":1969.19,
		"indexName":"VN30",
		"indexTypeCode":"101",
		"lowestValueIndexes":1947.28,
		"marketId":"STO",
		"marketIndexClass":"HSX",
		"multicastReceiveTime":{"Nanos":109572169,"Seconds":1781234600},
		"priorValueIndexes":1947.28,
		"totalVolumeTraded":86860936,
		"tradingSessionId":"40",
		"transactTime":{"Nanos":63000000,"Seconds":1781234600},
		"valueIndexes":1965.02
	}`), &event); err != nil {
		t.Fatal(err)
	}
	if event.IndexName != "VN30" || event.MarketID != "STO" || event.CurrencyCode != "VND" {
		t.Fatalf("event = %+v", event)
	}
	if event.ValueIndexes == nil || event.ValueIndexes.String() != "1965.02" {
		t.Fatalf("value indexes = %v", event.ValueIndexes)
	}
	if event.GrossTradeAmount == nil || event.GrossTradeAmount.String() != "2411.2342805" {
		t.Fatalf("gross trade amount = %v", event.GrossTradeAmount)
	}
	if event.TotalVolumeTraded != 86860936 || event.TransactTime.Seconds != 1781234600 {
		t.Fatalf("volume/time = %d/%+v", event.TotalVolumeTraded, event.TransactTime)
	}
}

func TestEstimatedMarketIndexEventDecodesCurrentWireFields(t *testing.T) {
	var event EstimatedMarketIndexEvent
	if err := json.Unmarshal([]byte(`{
		"T":"emi",
		"action":"estimated_market_index_update",
		"marketIndex":{
			"changedRatio":0.74,
			"changedValue":14.5,
			"fluctuationDownIssueCount":5,
			"fluctuationSteadinessIssueCount":2,
			"fluctuationUpIssueCount":23,
			"grossTradeAmount":6478.76,
			"indexName":"VN30",
			"time":"2026-06-12 13:51:28.242",
			"totalVolumeTraded":234239000,
			"valueIndexes":1961.78
		},
		"timestamp":1781247088242
	}`), &event); err != nil {
		t.Fatal(err)
	}
	if event.T != "emi" || event.Action != "estimated_market_index_update" || event.Timestamp != 1781247088242 {
		t.Fatalf("event = %+v", event)
	}
	index := event.MarketIndex
	if index.IndexName != "VN30" || index.Time != "2026-06-12 13:51:28.242" {
		t.Fatalf("market index = %+v", index)
	}
	assertDecimalString(t, "changed ratio", index.ChangedRatio, "0.74")
	assertDecimalString(t, "changed value", index.ChangedValue, "14.5")
	assertDecimalString(t, "gross trade amount", index.GrossTradeAmount, "6478.76")
	assertDecimalString(t, "value indexes", index.ValueIndexes, "1961.78")
	if index.FluctuationDownIssueCount != 5 || index.FluctuationSteadinessIssueCount != 2 || index.FluctuationUpIssueCount != 23 {
		t.Fatalf("fluctuation counts = %d/%d/%d", index.FluctuationDownIssueCount, index.FluctuationSteadinessIssueCount, index.FluctuationUpIssueCount)
	}
	if index.TotalVolumeTraded != 234239000 {
		t.Fatalf("total volume traded = %d", index.TotalVolumeTraded)
	}
}

func TestBrokerOrderEventDecodesWireFieldNames(t *testing.T) {
	var event BrokerOrderEvent
	if err := json.Unmarshal([]byte(`{
		"id":596,
		"side":"NS",
		"accountNo":"0001179019",
		"symbol":"41I1G5000",
		"orderType":"LO",
		"price":1920.0,
		"quantity":5,
		"fillQuantity":2,
		"canceledQuantity":0,
		"leaveQuantity":3,
		"orderStatus":"PartiallyFilled",
		"loanPackageId":2278,
		"marketType":"DERIVATIVE",
		"transDate":"2026-04-06T00:00:00Z",
		"createdDate":"2026-04-13T04:24:05.274Z",
		"modifiedDate":"2026-04-13T04:28:27.749Z"
	}`), &event); err != nil {
		t.Fatal(err)
	}
	if event.ID != 596 || event.AccountNo != "0001179019" || event.Symbol != "41I1G5000" {
		t.Fatalf("event = %+v", event)
	}
	if event.Price == nil || event.Price.String() != "1920" {
		t.Fatalf("price = %v", event.Price)
	}
	if event.Quantity == nil || event.Quantity.String() != "5" || event.FillQuantity == nil || event.FillQuantity.String() != "2" {
		t.Fatalf("quantities = %v/%v", event.Quantity, event.FillQuantity)
	}
	if event.MarketType != "DERIVATIVE" || event.OrderStatus != "PartiallyFilled" || event.LoanPackageID != 2278 {
		t.Fatalf("event = %+v", event)
	}
}
