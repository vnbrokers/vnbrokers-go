package dnse

import (
	"context"
	"net/url"
	"time"

	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type MarketDataSymbolsService struct {
	broker *Broker
}

func (s *MarketDataSymbolsService) List(
	ctx context.Context,
	symbol string,
	marketID string,
	securityGroupID string,
	indexName string,
	limit int,
	page int,
) ([]domain.Symbol, error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataSymbolsList); err != nil {
		return nil, err
	}
	params := url.Values{}
	if symbol != "" {
		params.Set("symbol", symbol)
	}
	if marketID != "" {
		params.Set("marketId", marketID)
	}
	if securityGroupID != "" {
		params.Set("securityGroupId", securityGroupID)
	}
	if indexName != "" {
		params.Set("indexName", indexName)
	}
	if limit == 0 {
		limit = s.broker.config.MarketDataSymbolLimit
	}
	params.Set("limit", stringify(limit))
	if page != 0 {
		params.Set("page", stringify(page))
	}
	response, err := s.broker.send(ctx, "marketdata.symbols.list", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/instruments?" + params.Encode()),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return nil, err
	}
	return MapSymbols(expectObject(response.Body)), nil
}

func (s *MarketDataSymbolsService) SecurityDefinition(
	ctx context.Context,
	symbol string,
	boardID string,
) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataSymbolsList); err != nil {
		return domain.RawPayload{}, err
	}
	if boardID == "" {
		boardID = s.broker.config.MarketDataBoardID
	}
	params := url.Values{}
	params.Set("boardId", boardID)
	response, err := s.broker.send(ctx, "marketdata.symbols.secdef", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/price/" + url.PathEscape(symbol) + "/secdef?" + params.Encode()),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}

func (s *MarketDataSymbolsService) WorkingDates(ctx context.Context) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataSymbolsList); err != nil {
		return domain.RawPayload{}, err
	}
	response, err := s.broker.send(ctx, "marketdata.symbols.working_dates", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/market/working-dates"),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}

type MarketDataQuotesService struct {
	broker *Broker
}

func (s *MarketDataQuotesService) Get(ctx context.Context, symbol string, boardID string) (domain.Quote, error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataQuotes); err != nil {
		return domain.Quote{}, err
	}
	if boardID == "" {
		boardID = s.broker.config.MarketDataBoardID
	}
	params := url.Values{}
	params.Set("boardId", boardID)
	response, err := s.broker.send(ctx, "marketdata.quotes.get", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/price/" + url.PathEscape(symbol) + "/trades/latest?" + params.Encode()),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return domain.Quote{}, err
	}
	return MapQuote(symbol, expectObject(response.Body)), nil
}

func (s *MarketDataQuotesService) PriceTrades(
	ctx context.Context,
	symbol string,
	fromTime int64,
	toTime int64,
	boardID string,
	limit int,
) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataQuotes); err != nil {
		return domain.RawPayload{}, err
	}
	if boardID == "" {
		boardID = s.broker.config.MarketDataBoardID
	}
	params := url.Values{}
	params.Set("boardId", boardID)
	params.Set("from", stringify(fromTime))
	params.Set("to", stringify(toTime))
	if limit != 0 {
		params.Set("limit", stringify(limit))
	}
	response, err := s.broker.send(ctx, "marketdata.quotes.trades", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/price/" + url.PathEscape(symbol) + "/trades?" + params.Encode()),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}

func (s *MarketDataQuotesService) ClosePrice(
	ctx context.Context,
	symbol string,
	boardID string,
) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataQuotes); err != nil {
		return domain.RawPayload{}, err
	}
	if boardID == "" {
		boardID = s.broker.config.MarketDataBoardID
	}
	params := url.Values{}
	params.Set("boardId", boardID)
	response, err := s.broker.send(ctx, "marketdata.quotes.close", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/price/" + url.PathEscape(symbol) + "/close?" + params.Encode()),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}

type MarketDataCandlesService struct {
	broker *Broker
}

func (s *MarketDataCandlesService) Get(
	ctx context.Context,
	symbol string,
	interval string,
	fromTime int64,
	toTime int64,
	marketType string,
) ([]domain.Candle, error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataCandles); err != nil {
		return nil, err
	}
	if toTime == 0 {
		toTime = time.Now().UTC().Unix()
	}
	if fromTime == 0 {
		fromTime = toTime - s.broker.config.CandleLookbackSeconds
	}
	if marketType == "" {
		marketType = s.broker.config.CandleMarketType
	}
	params := url.Values{}
	params.Set("symbol", symbol)
	params.Set("type", marketType)
	params.Set("resolution", interval)
	params.Set("from", stringify(fromTime))
	params.Set("to", stringify(toTime))
	response, err := s.broker.send(ctx, "marketdata.candles.get", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/price/ohlc?" + params.Encode()),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return nil, err
	}
	return MapCandles(symbol, interval, expectObject(response.Body)), nil
}
