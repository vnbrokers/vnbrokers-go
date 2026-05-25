package marketdata

type SubscribeSymbolRequest struct {
	Symbol  string
	Symbols []string
}

func (r SubscribeSymbolRequest) SymbolList() []string {
	if len(r.Symbols) > 0 {
		out := make([]string, len(r.Symbols))
		copy(out, r.Symbols)
		return out
	}
	if r.Symbol == "" {
		return nil
	}
	return []string{r.Symbol}
}
