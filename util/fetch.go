package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/mleyb/btc-reactor/models"
)

const url = "https://api.coindesk.com/v1/bpi/currentprice.json"

// Prices fetches the current prices index
func Prices() models.Price {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	price := models.Price{}

	jsonErr := json.Unmarshal(body, &price)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return price
}
