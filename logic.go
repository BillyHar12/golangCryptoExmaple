package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func (gc *GameCrypto) update() {
	for {
		resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var data []struct {
			ID           string  `json:"id"`
			CurrentPrice float64 `json:"current_price"`
			MarketCap    float64 `json:"market_cap"`
			Volume24h    float64 `json:"total_volume"`
			Change1h     float64 `json:"price_change_percentage_1h_in_currency"`
			Change24h    float64 `json:"price_change_percentage_24h_in_currency"`
			Change7d     float64 `json:"price_change_percentage_7d_in_currency"`
			Change30d    float64 `json:"price_change_percentage_30d_in_currency"`
			Change60d    float64 `json:"price_change_percentage_60d_in_currency"`
			Change90d    float64 `json:"price_change_percentage_90d_in_currency"`
			AllTimeHigh  float64 `json:"ath"`
		}
		if err := json.Unmarshal(body, &data); err != nil {
			panic(err)
		}

		for _, d := range data {
			gc.DB.Where("coin = ?", d.ID).Assign(GameCrypto{
				Coin:        d.ID,
				Price:       uint(d.CurrentPrice * 100),
				MarketCap:   uint(d.MarketCap),
				Volume24h:   uint(d.Volume24h),
				Change1h:    uint(d.Change1h),
				Change24h:   uint(d.Change24h),
				Change7d:    uint(d.Change7d),
				Change30d:   uint(d.Change30d),
				Change60d:   uint(d.Change60d),
				Change90d:   uint(d.Change90d),
				AllTimeHigh: uint(d.AllTimeHigh),
			}).FirstOrCreate(&GameCrypto{})

			time.Sleep(1 * time.Second)
		}

		time.Sleep(5 * time.Minute)
	}
}
