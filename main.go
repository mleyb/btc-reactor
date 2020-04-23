package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mleyb/btc-reactor/aws"
	"github.com/mleyb/btc-reactor/env"
	"github.com/mleyb/btc-reactor/util"
)

func handler(ctx context.Context, event events.CloudWatchEvent) {

	//lc, _ := lambdacontext.FromContext(ctx)
	//lc.ClientContext.
	prices := util.Prices()

	if prices.Bpi.GBP.RateFloat > env.Limit() {
		aws.PublishSMS("BTC price is over limit", env.PhoneNumber())
	}
}

func main() {
	lambda.Start(handler)
}

// https://api.coindesk.com/v1/bpi/currentprice.json
