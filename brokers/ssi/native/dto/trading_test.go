package dto

import (
	"encoding/json"
	"testing"
)

func TestTradingRequestTypesFollowJSONExamples(t *testing.T) {
	var accessToken AccessTokenRequest
	if err := json.Unmarshal([]byte(`{
		"consumerID":"consumer",
		"consumerSecret":"secret",
		"twoFactorType":1,
		"code":"123456",
		"isSave":false
	}`), &accessToken); err != nil {
		t.Fatalf("unmarshal access token request: %v", err)
	}

	if accessToken.TwoFactorType != 1 {
		t.Fatalf("twoFactorType = %d", accessToken.TwoFactorType)
	}
	if accessToken.IsSave {
		t.Fatal("isSave = true")
	}

	_ = NewOrderRequest{}
	_ = VSDCashDWRequest{}
	_ = StockTransferRequest{}
	_ = RightsCreateRequest{}
	_ = FCONewOrderRequest{}
}

func TestOrderBookResponseUsesExampleNumberTypes(t *testing.T) {
	var response OrderBookResponse
	if err := json.Unmarshal([]byte(`{
		"message":"Success",
		"status":200,
		"data":{
			"account":"1184418",
			"orders":[{
				"price":1000,
				"quantity":100,
				"filledQty":0,
				"cancelQty":0,
				"avgPrice":0,
				"isForcesell":"F",
				"isShortsell":"F",
				"lastErrorEvent":null
			}]
		}
	}`), &response); err != nil {
		t.Fatalf("unmarshal order book response: %v", err)
	}

	order := response.Data.Orders[0]
	if order.Price != 1000 || order.Quantity != 100 || order.FilledQty != 0 {
		t.Fatalf("unexpected order quantities: %+v", order)
	}

	_ = AuditOrderBookResponse{}
	_ = OrderHistoryResponse{}
	_ = CashInAdvanceAmountResponse{}
	_ = StockTransferResponse{}
	_ = RightsCreateResponse{}
	_ = FCOListResponse{}
}
