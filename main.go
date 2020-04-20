package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	//	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/mleyb/btc-reactor/util"
	"log"
)

func handler(ctx context.Context, event events.CloudWatchEvent) {

	//lc, _ := lambdacontext.FromContext(ctx)
	//lc.ClientContext.
	prices := util.Prices()

	log.Println(prices.Bpi.GBP.Rate)
}

func main() {
	lambda.Start(handler)
}

// https://api.coindesk.com/v1/bpi/currentprice.json
