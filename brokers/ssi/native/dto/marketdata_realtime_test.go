package dto

import (
	"encoding/json"
	"testing"
)

func TestQuoteEventUnmarshal(t *testing.T) {
	data := []byte(`{
		"TradingDate":"14/08/2023",
		"Time":"14:00:28",
		"Exchange":"HOSE",
		"Symbol":"ACB",
		"RType":"X-QUOTE",
		"AskPrice1":22950.0,
		"AskPrice10":0.0,
		"AskVol1":109100.0,
		"AskVol10":0.0,
		"BidPrice1":22900.0,
		"BidPrice10":0.0,
		"BidVol1":290900.0,
		"BidVol10":0.0,
		"TradingSession":"LO"
	}`)

	var event QuoteEvent
	if err := json.Unmarshal(data, &event); err != nil {
		t.Fatalf("unmarshal QuoteEvent: %v", err)
	}

	if event.TradingDate != "14/08/2023" || event.Time != "14:00:28" {
		t.Fatalf("unexpected event time: date=%q time=%q", event.TradingDate, event.Time)
	}
	if event.Exchange != "HOSE" || event.Symbol != "ACB" || event.RType != "X-QUOTE" {
		t.Fatalf("unexpected event identity: exchange=%q symbol=%q rtype=%q", event.Exchange, event.Symbol, event.RType)
	}
	if event.AskPrice1 != 22950 || event.AskVol1 != 109100 {
		t.Fatalf("unexpected ask level: price=%v volume=%v", event.AskPrice1, event.AskVol1)
	}
	if event.BidPrice1 != 22900 || event.BidVol1 != 290900 {
		t.Fatalf("unexpected bid level: price=%v volume=%v", event.BidPrice1, event.BidVol1)
	}
	if event.AskPrice10 != 0 || event.AskVol10 != 0 || event.BidPrice10 != 0 || event.BidVol10 != 0 {
		t.Fatal("expected zero values at level 10")
	}
	if event.TradingSession != "LO" {
		t.Fatalf("unexpected trading session: %q", event.TradingSession)
	}
}

func TestForeignRoomEventUnmarshal(t *testing.T) {
	data := []byte(`{
		"RType":"R",
		"TradingDate":"14/08/2023",
		"Time":"14:24:02",
		"Isin":"SSI",
		"Symbol":"SSI",
		"TotalRoom":1501130137.0,
		"CurrentRoom":806887173.0,
		"BuyVol":863352.0,
		"SellVol":825308.0,
		"BuyVal":25123543200.0,
		"SellVal":24016462800.0,
		"MarketId":"HOSE",
		"Exchange":"HOSE"
	}`)

	var event ForeignRoomEvent
	if err := json.Unmarshal(data, &event); err != nil {
		t.Fatalf("unmarshal ForeignRoomEvent: %v", err)
	}

	if event.RType != "R" || event.TradingDate != "14/08/2023" || event.Time != "14:24:02" {
		t.Fatalf("unexpected event metadata: rtype=%q date=%q time=%q", event.RType, event.TradingDate, event.Time)
	}
	if event.Isin != "SSI" || event.Symbol != "SSI" || event.MarketID != "HOSE" || event.Exchange != "HOSE" {
		t.Fatalf("unexpected security identity: isin=%q symbol=%q market=%q exchange=%q", event.Isin, event.Symbol, event.MarketID, event.Exchange)
	}
	if event.TotalRoom != 1501130137 || event.CurrentRoom != 806887173 {
		t.Fatalf("unexpected room values: total=%v current=%v", event.TotalRoom, event.CurrentRoom)
	}
	if event.BuyVol != 863352 || event.SellVol != 825308 {
		t.Fatalf("unexpected volumes: buy=%v sell=%v", event.BuyVol, event.SellVol)
	}
	if event.BuyVal != 25123543200 || event.SellVal != 24016462800 {
		t.Fatalf("unexpected values: buy=%v sell=%v", event.BuyVal, event.SellVal)
	}
}

func TestTradingStatusEventUnmarshal(t *testing.T) {
	data := []byte(`{
		"RType":"F",
		"MarketId":"HOSE",
		"TradingDate":"14/08/2023",
		"Time":"13:00:00",
		"Symbol":"SSI",
		"TradingSession":"LO",
		"TradingStatus":"N",
		"Exchange":"HOSE"
	}`)

	var event TradingStatusEvent
	if err := json.Unmarshal(data, &event); err != nil {
		t.Fatalf("unmarshal TradingStatusEvent: %v", err)
	}

	if event.RType != "F" || event.MarketID != "HOSE" {
		t.Fatalf("unexpected event type: rtype=%q market=%q", event.RType, event.MarketID)
	}
	if event.TradingDate != "14/08/2023" || event.Time != "13:00:00" || event.Symbol != "SSI" {
		t.Fatalf("unexpected event identity: date=%q time=%q symbol=%q", event.TradingDate, event.Time, event.Symbol)
	}
	if event.TradingSession != "LO" || event.TradingStatus != "N" || event.Exchange != "HOSE" {
		t.Fatalf("unexpected trading state: session=%q status=%q exchange=%q", event.TradingSession, event.TradingStatus, event.Exchange)
	}
}

func TestTradeEventUnmarshal(t *testing.T) {
	data := []byte(`{
		"RType":"X-TRADE",
		"TradingDate":"14/08/2023",
		"Time":"14:07:53",
		"Isin":"ACB",
		"Symbol":"ACB",
		"Ceiling":24500.0,
		"Floor":21300.0,
		"RefPrice":22900.0,
		"AvgPrice":22927.93,
		"PriorVal":22900.0,
		"LastPrice":22950.0,
		"LastVol":100.0,
		"TotalVal":201713015000.0,
		"TotalVol":8797700.0,
		"MarketId":"HOSE",
		"Exchange":"HOSE",
		"TradingSession":"LO",
		"TradingStatus":"N",
		"Change":50.0,
		"RatioChange":0.22,
		"EstMatchedPrice":22900.0,
		"Highest":23050,
		"Lowest":22800,
		"Side":"SD"
	}`)

	var event TradeEvent
	if err := json.Unmarshal(data, &event); err != nil {
		t.Fatalf("unmarshal TradeEvent: %v", err)
	}

	if event.RType != "X-TRADE" || event.Symbol != "ACB" || event.Isin != "ACB" {
		t.Fatalf("unexpected trade identity: rtype=%q symbol=%q isin=%q", event.RType, event.Symbol, event.Isin)
	}
	if event.TradingDate != "14/08/2023" || event.Time != "14:07:53" || event.MarketID != "HOSE" || event.Exchange != "HOSE" {
		t.Fatalf("unexpected trade metadata: date=%q time=%q market=%q exchange=%q", event.TradingDate, event.Time, event.MarketID, event.Exchange)
	}
	if event.Ceiling != 24500 || event.Floor != 21300 || event.RefPrice != 22900 || event.AvgPrice != 22927.93 {
		t.Fatalf("unexpected reference prices: ceiling=%v floor=%v ref=%v avg=%v", event.Ceiling, event.Floor, event.RefPrice, event.AvgPrice)
	}
	if event.PriorVal != 22900 || event.LastPrice != 22950 || event.LastVol != 100 {
		t.Fatalf("unexpected last trade: prior=%v last=%v vol=%v", event.PriorVal, event.LastPrice, event.LastVol)
	}
	if event.TotalVal != 201713015000 || event.TotalVol != 8797700 {
		t.Fatalf("unexpected totals: value=%v volume=%v", event.TotalVal, event.TotalVol)
	}
	if event.TradingSession != "LO" || event.TradingStatus != "N" || event.Side != "SD" {
		t.Fatalf("unexpected session state: session=%q status=%q side=%q", event.TradingSession, event.TradingStatus, event.Side)
	}
	if event.Change != 50 || event.RatioChange != 0.22 || event.EstMatchedPrice != 22900 || event.Highest != 23050 || event.Lowest != 22800 {
		t.Fatalf("unexpected price movement: change=%v ratio=%v est=%v high=%v low=%v", event.Change, event.RatioChange, event.EstMatchedPrice, event.Highest, event.Lowest)
	}
}

func TestSnapshotEventUnmarshal(t *testing.T) {
	data := []byte(`{
		"RType":"X",
		"TradingDate":"15/08/2023",
		"Time":"10:28:41",
		"Isin":"TAR",
		"Symbol":"TAR",
		"Ceiling":23400.0,
		"Floor":19200.0,
		"RefPrice":21300.0,
		"Open":21300.0,
		"High":22300.0,
		"Low":21300.0,
		"Close":21800.0,
		"AvgPrice": 21784.0,
		"PriorVal":21300.0,
		"LastPrice":21800.0,
		"LastVol":100.0,
		"TotalVal":27946100000.0,
		"TotalVol":1282900.0,
		"BidPrice1":21800.0,
		"BidPrice10":20900.0,
		"BidVol1":18800.0,
		"BidVol10":26100.0,
		"AskPrice1":21900.0,
		"AskPrice10":22800.0,
		"AskVol1":50300.0,
		"AskVol10":17000.0,
		"MarketId":"HNX",
		"Exchange":"HNX",
		"TradingSession":"LO",
		"TradingStatus":"N",
		"Change":500.0,
		"RatioChange":2.35,
		"EstMatchedPrice":0.0
	}`)

	var event SnapshotEvent
	if err := json.Unmarshal(data, &event); err != nil {
		t.Fatalf("unmarshal SnapshotEvent: %v", err)
	}

	if event.RType != "X" || event.Symbol != "TAR" || event.Isin != "TAR" {
		t.Fatalf("unexpected snapshot identity: rtype=%q symbol=%q isin=%q", event.RType, event.Symbol, event.Isin)
	}
	if event.TradingDate != "15/08/2023" || event.Time != "10:28:41" || event.MarketID != "HNX" || event.Exchange != "HNX" {
		t.Fatalf("unexpected snapshot metadata: date=%q time=%q market=%q exchange=%q", event.TradingDate, event.Time, event.MarketID, event.Exchange)
	}
	if event.Ceiling != 23400 || event.Floor != 19200 || event.RefPrice != 21300 {
		t.Fatalf("unexpected reference levels: ceiling=%v floor=%v ref=%v", event.Ceiling, event.Floor, event.RefPrice)
	}
	if event.Open != 21300 || event.High != 22300 || event.Low != 21300 || event.Close != 21800 || event.AvgPrice != 21784.0 {
		t.Fatalf("unexpected OHLC data: open=%v high=%v low=%v close=%v avg=%v", event.Open, event.High, event.Low, event.Close, event.AvgPrice)
	}
	if event.LastPrice != 21800 || event.LastVol != 100 || event.TotalVal != 27946100000 || event.TotalVol != 1282900 {
		t.Fatalf("unexpected trade totals: last=%v vol=%v totalVal=%v totalVol=%v", event.LastPrice, event.LastVol, event.TotalVal, event.TotalVol)
	}
	if event.BidPrice1 != 21800 || event.BidPrice10 != 20900 || event.BidVol1 != 18800 || event.BidVol10 != 26100 {
		t.Fatalf("unexpected bid book: p1=%v p10=%v v1=%v v10=%v", event.BidPrice1, event.BidPrice10, event.BidVol1, event.BidVol10)
	}
	if event.AskPrice1 != 21900 || event.AskPrice10 != 22800 || event.AskVol1 != 50300 || event.AskVol10 != 17000 {
		t.Fatalf("unexpected ask book: p1=%v p10=%v v1=%v v10=%v", event.AskPrice1, event.AskPrice10, event.AskVol1, event.AskVol10)
	}
	if event.TradingSession != "LO" || event.TradingStatus != "N" {
		t.Fatalf("unexpected trading state: session=%q status=%q", event.TradingSession, event.TradingStatus)
	}
	if event.Change != 500 || event.RatioChange != 2.35 || event.EstMatchedPrice != 0 {
		t.Fatalf("unexpected price movement: change=%v ratio=%v est=%v", event.Change, event.RatioChange, event.EstMatchedPrice)
	}
}

func TestSnapshotEventUnmarshalNaNAvgPrice(t *testing.T) {
	data := []byte(`{
		"RType":"X",
		"AvgPrice": "NaN"
	}`)

	var event SnapshotEvent
	if err := json.Unmarshal(data, &event); err != nil {
		t.Fatalf("unmarshal SnapshotEvent: %v", err)
	}

	if event.AvgPrice != 0 {
		t.Fatalf("unexpected OHLC data: avg=%v", event.AvgPrice)
	}
}

func TestMarketIndexEventUnmarshal(t *testing.T) {
	data := []byte(`{
		"IndexId":"VN30",
		"IndexValEst":1200.03,
		"IndexValue":1238.76,
		"PriorIndexValue":1226.16,
		"TradingDate":"02/04/2021",
		"Time":"11:28:13",
		"TotalTrade":0.0,
		"TotalQtty":191838100.0,
		"TotalValue":7289093000000.0,
		"IndexName":"VN30",
		"Advances":25,
		"NoChanges":2,
		"Declines":3,
		"Ceilings":0,
		"Floors":0,
		"Change":12.6,
		"RatioChange":1.03,
		"TotalQttyPt":2064000.0,
		"TotalValuePt":244251000000.0,
		"Exchange":"HOSE",
		"AllQty":193902100.0,
		"AllValue":7533344000000.0,
		"IndexType":"Main",
		"TradingSession":null,
		"MarketId":null,
		"RType":"MI",
		"TotalQttyOd":0.0,
		"TotalValueOd":0.0
	}`)

	var event MarketIndexEvent
	if err := json.Unmarshal(data, &event); err != nil {
		t.Fatalf("unmarshal MarketIndexEvent: %v", err)
	}

	if event.IndexID != "VN30" || event.IndexName != "VN30" || event.RType != "MI" {
		t.Fatalf("unexpected index identity: id=%q name=%q rtype=%q", event.IndexID, event.IndexName, event.RType)
	}
	if event.IndexValEst != 1200.03 || event.IndexValue != 1238.76 || event.PriorIndexValue != 1226.16 {
		t.Fatalf("unexpected index values: est=%v value=%v prior=%v", event.IndexValEst, event.IndexValue, event.PriorIndexValue)
	}
	if event.TradingDate != "02/04/2021" || event.Time != "11:28:13" || event.Exchange != "HOSE" || event.IndexType != "Main" {
		t.Fatalf("unexpected metadata: date=%q time=%q exchange=%q type=%q", event.TradingDate, event.Time, event.Exchange, event.IndexType)
	}
	if event.TotalQtty != 191838100 || event.TotalValue != 7289093000000 || event.AllQty != 193902100 || event.AllValue != 7533344000000 {
		t.Fatalf("unexpected volume totals: totalQty=%v totalValue=%v allQty=%v allValue=%v", event.TotalQtty, event.TotalValue, event.AllQty, event.AllValue)
	}
	if event.Advances != 25 || event.NoChanges != 2 || event.Declines != 3 || event.Ceilings != 0 || event.Floors != 0 {
		t.Fatalf("unexpected breadth values: advances=%v noChanges=%v declines=%v ceilings=%v floors=%v", event.Advances, event.NoChanges, event.Declines, event.Ceilings, event.Floors)
	}
	if event.Change != 12.6 || event.RatioChange != 1.03 || event.TotalQttyPt != 2064000 || event.TotalValuePt != 244251000000 {
		t.Fatalf("unexpected change values: change=%v ratio=%v qtyPt=%v valuePt=%v", event.Change, event.RatioChange, event.TotalQttyPt, event.TotalValuePt)
	}
	if event.TradingSession != nil || event.MarketID != nil {
		t.Fatalf("expected nil nullable fields: tradingSession=%v marketID=%v", event.TradingSession, event.MarketID)
	}
	if event.TotalQttyOd != 0 || event.TotalValueOd != 0 || event.TotalTrade != 0 {
		t.Fatalf("unexpected order values: totalQttyOd=%v totalValueOd=%v totalTrade=%v", event.TotalQttyOd, event.TotalValueOd, event.TotalTrade)
	}
}

func TestOHLCVEventUnmarshal(t *testing.T) {
	data := []byte(`{
		"RType":"B",
		"Symbol":"X26",
		"Time":"14:28:33",
		"Open":16000,
		"High":16000,
		"Low":16000,
		"Close":16000,
		"Volume":5000,
		"Value":0
	}`)

	var event OHLCVEvent
	if err := json.Unmarshal(data, &event); err != nil {
		t.Fatalf("unmarshal OHLCVEvent: %v", err)
	}

	if event.RType != "B" || event.Symbol != "X26" || event.Time != "14:28:33" {
		t.Fatalf("unexpected OHLCV identity: rtype=%q symbol=%q time=%q", event.RType, event.Symbol, event.Time)
	}
	if event.Open != 16000 || event.High != 16000 || event.Low != 16000 || event.Close != 16000 {
		t.Fatalf("unexpected OHLC values: open=%v high=%v low=%v close=%v", event.Open, event.High, event.Low, event.Close)
	}
	if event.Volume != 5000 || event.Value != 0 {
		t.Fatalf("unexpected volume values: volume=%v value=%v", event.Volume, event.Value)
	}
}

func TestOddLotEventUnmarshal(t *testing.T) {
	data := []byte(`{
		"RType":"OL",
		"TradingDate":"18/02/2025",
		"Time":"13:55:03",
		"StockNo":2027,
		"Symbol":"MBB",
		"Ceiling":24200,
		"Floor":21100,
		"RefPrice":22650,
		"Open":22650,
		"High":22950,
		"Low":22600,
		"LastPrice":22750,
		"LastVol":9193,
		"TotalVal":185028289999.99728,
		"TotalVol":8135000,
		"BidPrice1":22700,
		"BidPrice2":22650,
		"BidPrice3":22600,
		"BidVol1":1108,
		"BidVol2":1630,
		"BidVol3":2016,
		"AskPrice1":22750,
		"AskPrice2":22800,
		"AskPrice3":22850,
		"AskVol1":132,
		"AskVol2":548,
		"AskVol3":297,
		"Exchange":"HOSE",
		"TradingSession":"LO",
		"TradingStatus":"H",
		"Change":100,
		"RatioChange":0.44,
		"StockType":"Stock"
	}`)

	var event OddLotEvent
	if err := json.Unmarshal(data, &event); err != nil {
		t.Fatalf("unmarshal OddLotEvent: %v", err)
	}

	if event.RType != "OL" || event.Symbol != "MBB" || event.StockNo != 2027 {
		t.Fatalf("unexpected odd-lot identity: rtype=%q symbol=%q stockNo=%v", event.RType, event.Symbol, event.StockNo)
	}
	if event.TradingDate != "18/02/2025" || event.Time != "13:55:03" || event.Exchange != "HOSE" {
		t.Fatalf("unexpected odd-lot metadata: date=%q time=%q exchange=%q", event.TradingDate, event.Time, event.Exchange)
	}
	if event.Ceiling != 24200 || event.Floor != 21100 || event.RefPrice != 22650 {
		t.Fatalf("unexpected reference levels: ceiling=%v floor=%v ref=%v", event.Ceiling, event.Floor, event.RefPrice)
	}
	if event.Open != 22650 || event.High != 22950 || event.Low != 22600 || event.LastPrice != 22750 || event.LastVol != 9193 {
		t.Fatalf("unexpected trade levels: open=%v high=%v low=%v last=%v lastVol=%v", event.Open, event.High, event.Low, event.LastPrice, event.LastVol)
	}
	if event.TotalVal != 185028289999.99728 || event.TotalVol != 8135000 {
		t.Fatalf("unexpected totals: totalVal=%v totalVol=%v", event.TotalVal, event.TotalVol)
	}
	if event.BidPrice1 != 22700 || event.BidPrice2 != 22650 || event.BidPrice3 != 22600 || event.BidVol1 != 1108 || event.BidVol2 != 1630 || event.BidVol3 != 2016 {
		t.Fatalf("unexpected bid book: p1=%v p2=%v p3=%v v1=%v v2=%v v3=%v", event.BidPrice1, event.BidPrice2, event.BidPrice3, event.BidVol1, event.BidVol2, event.BidVol3)
	}
	if event.AskPrice1 != 22750 || event.AskPrice2 != 22800 || event.AskPrice3 != 22850 || event.AskVol1 != 132 || event.AskVol2 != 548 || event.AskVol3 != 297 {
		t.Fatalf("unexpected ask book: p1=%v p2=%v p3=%v v1=%v v2=%v v3=%v", event.AskPrice1, event.AskPrice2, event.AskPrice3, event.AskVol1, event.AskVol2, event.AskVol3)
	}
	if event.TradingSession != "LO" || event.TradingStatus != "H" || event.StockType != "Stock" {
		t.Fatalf("unexpected state: session=%q status=%q stockType=%q", event.TradingSession, event.TradingStatus, event.StockType)
	}
	if event.Change != 100 || event.RatioChange != 0.44 {
		t.Fatalf("unexpected movement: change=%v ratio=%v", event.Change, event.RatioChange)
	}
}
