package main

import (
	stock_fetcher "stock_fetchernstore"
	"stock_fetchernstore/internal/applog"
	"stock_fetchernstore/internal/config"
)

func main() {
	config.Load()
	applog.Init()
	stock_fetcher.Fetch()
}
