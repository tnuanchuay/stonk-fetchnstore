package stock_fetcher

import (
	"fmt"
	"stock_fetchernstore/internal/fapi_cli"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	DEFAULT_INTERVAL = "1d"
	DEFAULT_RANGE    = "1y"
)

func Fetch() {
	facli := fapi_cli.Get()
	symbols := viper.GetStringSlice("symbols")
	for _, symbol := range symbols {
		result, err := facli.GetTicker(symbol, DEFAULT_INTERVAL, DEFAULT_RANGE)
		if err != nil {
			log.Error(err)
			continue
		}

		fmt.Println(result.Chart.Result[0].Timestamp)
		// fmt.Println(result.Chart.Result[0].Indicators.Quote[0].Close)

		break
	}
}

// func FetchSymbolOnDate()
