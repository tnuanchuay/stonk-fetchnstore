package fapi_cli

import (
	"github.com/spf13/viper"
	ficli "github.com/tspn/fapi-client"
)

func Get() *ficli.FinanceApiClient {
	baseUrl := viper.GetString("financeApi.baseUrl")
	apiKey := viper.GetString("financeApi.apiKey")

	return ficli.New(baseUrl, apiKey)
}
