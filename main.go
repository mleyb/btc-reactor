package main

import (
	"context"
	"fmt"
	_ "github.com/aws/aws-lambda-go/lambda"
	"github.com/mleyb/btc-reactor/models"
	"github.com/mleyb/btc-reactor/util"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	//lambda.Start(HandleRequest)

	var prices models.CurrentPrice


	fmt.Println("test")

	prices = util.Prices()

	fmt.Sprintf("Prices %s", prices)

	//log.Println("OK")
}

// https://api.coindesk.com/v1/bpi/currentprice.json
