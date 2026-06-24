// Code generated from Finhay OpenAPI v2; DO NOT EDIT.

package dto

type AvailableTradeResult struct {
	// Buying power in VND — the max amount available to BUY {symbol} at quotePrice. Use for BUY: check pp0 >= quantity * quotePrice.
	Pp0 int64 `json:"pp0,omitempty"`
	// Max shares after settlement (purchasing power after settlement)
	Ppse int64 `json:"ppse,omitempty"`
	// Max quantity cap from the broker
	Maxqtty int64 `json:"maxqtty,omitempty"`
	// Number of shares of {symbol} available to SELL. Use for SELL: check trade >= quantity.
	Trade int64 `json:"trade,omitempty"`
	// Account cash balance (VND)
	Balance int64 `json:"balance,omitempty"`
	// Cash pending transfer (VND)
	CashPendingSend int64 `json:"cash_pending_send,omitempty"`
	// Mortgage value (VND)
	Mortgage int64 `json:"mortgage,omitempty"`
	// Margin rate
	Marginrate float64 `json:"marginrate,omitempty"`
	// Margin ratio loan
	Mrratioloan *float64 `json:"mrratioloan,omitempty"`
}

type CancelOrderRequest struct {
	// Extended sub-account ID (use SUB_ACCOUNT_EXT_ORDER — must end in `.4`)
	SubAccount string `json:"sub_account"`
}

type CreateOrderRequest struct {
	// Extended sub-account ID (must end in `.4` — the order-execution account)
	SubAccount string `json:"sub_account,omitempty"`
	// Order side Enum: BUY, SELL
	Side string `json:"side"`
	// Stock symbol (e.g. HPG, VNM, FPT)
	Symbol string `json:"symbol"`
	// Number of shares
	Quantity int64 `json:"quantity"`
	// Order type. Determines which price field to use. Enum: LIMIT, MARKET
	TypeValue string `json:"type"`
	// Limit price in VND. Set when type=LIMIT, null when type=MARKET.
	LimitPrice *int64 `json:"limit_price,omitempty"`
	// Market price type. Set when type=MARKET, null when type=LIMIT. Enum: MP, ATO, ATC, MAK, MOK, MTL, PLO, FOK, FAK
	MarketPrice *string `json:"market_price,omitempty"`
	// Securities type. Default STOCK for most orders. Enum: STOCK, BOND, FUND_CERTIFICATE, WARRANT, ETF
	StockType string `json:"stock_type,omitempty"`
}

type ExchangeSessionInfo struct {
	// Exchange code Enum: HOSE, HNX, UPCOM, HCX
	Exchange string `json:"exchange,omitempty"`
	// Current session: `OPEN` = ATO (opening auction), `PROGRESS` = continuous matching, `BREAK` = lunch break, `CLOSED` = market closed, `PRE_CLOSED` = ATC (closing auction), `PUT_THROUGH` = put-through session, `POST_SESSION` = HNX post-session, `CA` = periodic auction Enum: OPEN, PROGRESS, BREAK, CLOSED, PRE_CLOSED, PUT_THROUGH, POST_SESSION, CA
	ExchangeSession string `json:"exchange_session,omitempty"`
	// Order types available in the current session
	AvailableOrderTypes []string `json:"available_order_types,omitempty"`
}

type OrderBookEntry struct {
	// Strategy ID if the order is linked to a strategy
	StrategyID *string `json:"strategy_id,omitempty"`
	Custodycd  string  `json:"custodycd,omitempty"`
	// Transaction date
	Txdate string `json:"txdate,omitempty"`
	Custid string `json:"custid,omitempty"`
	// Account number
	Afacctno string `json:"afacctno,omitempty"`
	// Order ID
	Orderid   string `json:"orderid,omitempty"`
	Odorderid string `json:"odorderid,omitempty"`
	// Transaction time
	Txtime string `json:"txtime,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	// Whether cancel is allowed
	Allowcancel string `json:"allowcancel,omitempty"`
	// Whether amendment is allowed
	Allowamend string `json:"allowamend,omitempty"`
	// Side code (raw)
	SideCode string `json:"side_code,omitempty"`
	// Side display name (BUY/SELL)
	Side  string `json:"side,omitempty"`
	Price int64  `json:"price,omitempty"`
	// Price type code
	Pricetype string `json:"pricetype,omitempty"`
	// Order channel code
	ViaCode string `json:"via_code,omitempty"`
	// Order channel name
	Via string `json:"via,omitempty"`
	// Order quantity
	Qtty int64 `json:"qtty,omitempty"`
	// Executed quantity
	Execqtty int64 `json:"execqtty,omitempty"`
	// Executed amount
	Execamt int64 `json:"execamt,omitempty"`
	// Executed price
	Execprice int64 `json:"execprice,omitempty"`
	// Remaining quantity
	Remainqtty int64 `json:"remainqtty,omitempty"`
	// Remaining amount
	Remainamt int64 `json:"remainamt,omitempty"`
	// Status code (raw)
	StatusCode string `json:"status_code,omitempty"`
	// Status display name
	Status   string `json:"status,omitempty"`
	Tlname   string `json:"tlname,omitempty"`
	Username string `json:"username,omitempty"`
	// HOSE session info
	Hosesession   string `json:"hosesession,omitempty"`
	Cancelqtty    int64  `json:"cancelqtty,omitempty"`
	Adjustqtty    int64  `json:"adjustqtty,omitempty"`
	Isdisposal    string `json:"isdisposal,omitempty"`
	Rootorderid   string `json:"rootorderid,omitempty"`
	Timetype      string `json:"timetype,omitempty"`
	Timetypevalue string `json:"timetypevalue,omitempty"`
	// Feedback message from exchange
	Feedbackmsg            string  `json:"feedbackmsg,omitempty"`
	Quoteqtty              int64   `json:"quoteqtty,omitempty"`
	Limitprice             float64 `json:"limitprice,omitempty"`
	Odtimestamp            string  `json:"odtimestamp,omitempty"`
	MatchtypeCode          string  `json:"matchtype_code,omitempty"`
	Producttypename        string  `json:"producttypename,omitempty"`
	AfacctnoExt            string  `json:"afacctno_ext,omitempty"`
	SideSideUnderscore     string  `json:"side_,omitempty"`
	ViaViaUnderscore       string  `json:"via_,omitempty"`
	StatusStatusUnderscore string  `json:"status_,omitempty"`
	Matchtype              string  `json:"matchtype_,omitempty"`
}

type OrderResponse struct {
	// Exchange order ID
	OrderID         string `json:"order_id,omitempty"`
	AccountID       string `json:"account_id,omitempty"`
	TransactionDate string `json:"transaction_date,omitempty"`
	Symbol          string `json:"symbol,omitempty"`
	// Enum: BUY, SELL
	OrderSide       string `json:"order_side,omitempty"`
	OrderQuantity   int64  `json:"order_quantity,omitempty"`
	LimitPrice      int64  `json:"limit_price,omitempty"`
	MarketPrice     string `json:"market_price,omitempty"`
	ExecuteQuantity int64  `json:"execute_quantity,omitempty"`
	ExecutePrice    int64  `json:"execute_price,omitempty"`
	// Initial status after placement (usually RECEIVED or SENT) Enum: RECEIVED, SENT, MATCHED, CANCELLED, REJECTED, FAILED
	OrderStatus   string  `json:"order_status,omitempty"`
	FeeAmount     float64 `json:"fee_amount,omitempty"`
	TaxAmount     float64 `json:"tax_amount,omitempty"`
	ExecuteAmount int64   `json:"execute_amount,omitempty"`
	// Enum: LO, MP, ATO, ATC, MAK, MOK, MTL, PLO, FOK, FAK
	OrderType string `json:"order_type,omitempty"`
	// Error code if order was rejected. See error-codes.md.
	Code *string `json:"code,omitempty"`
	// Human-readable rejection reason
	RejectedReason *string `json:"rejected_reason,omitempty"`
	// EVEN = round lot (≥100 shares), ODD = odd lot (1-99 shares) Enum: EVEN, ODD
	Lot string `json:"lot,omitempty"`
}

type UpdateOrderRequest struct {
	// New quantity
	Quantity int64 `json:"quantity,omitempty"`
	// New limit price in VND
	Price int64 `json:"price,omitempty"`
}
