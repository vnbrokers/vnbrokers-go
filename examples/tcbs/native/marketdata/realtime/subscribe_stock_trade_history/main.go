package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
	"github.com/vnbrokers/vnbrokers-go/internal/env"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, env.Duration("TCBS_STREAM_DURATION", 8*time.Hour))
	defer cancel()

	broker := vnbrokers.NewTCBS(tcbs.Config{AccessToken: env.RequiredString("TCBS_ACCESS_TOKEN")})
	tickers := "ACB,BID,BSR,CTG,FPT,GAS,GVR,HDB,HPG,LPB,MBB,MSN,MWG,PLX,SAB,SHB,SSB,SSI,STB,TCB,TPB,VCB,VHM,VIB,VIC,VJC,VNM,VPB,VPL,VRE"
	request := nativedto.SubscribeStockTradeHistoryRequest{Tickers: env.List("TCBS_TICKERS", tickers)}
	subscription, err := broker.Native().MarketData().Realtime().SubscribeStockTradeHistory(ctx, request)
	if err != nil {
		panic(err)
	}
	defer subscription.Close()

	for {
		select {
		case event, ok := <-subscription.Events():
			if !ok {
				return
			}
			fmt.Printf("event: %+v\n", event)
		case status, ok := <-subscription.Status():
			if !ok {
				return
			}
			fmt.Printf("status: %s\n", status)
		case err, ok := <-subscription.Errors():
			if !ok {
				return
			}
			fmt.Printf("error: %v\n", err)
		case <-ctx.Done():
			return
		}
	}
}
