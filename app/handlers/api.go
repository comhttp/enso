package handlers

import (
	"fmt"
	"strings"

	explorer "github.com/comhttp/explorer/app"
	"github.com/comhttp/jorm/mod/coin"
	"github.com/comhttp/jorm/pkg/strapi"
	"github.com/comhttp/jorm/pkg/utl"
)

type JormCmds struct {
	CMDs     map[string](func(vars map[string]interface{}) map[string]interface{})
	Command  string
	Variable string
}

type API struct {
	UrlFormat string
	Endpoint  string
	Commands  map[string]string `json:"commands"`
}

type JormAPI struct {
	OKNO         strapi.StrapiRestClient
	Apis         map[string][]API
	JORMcommands *JormCmds
	ExJDB        map[string]*explorer.ExplorerJDBs
	Coin         string
}

// func (j *jormAPI)getAPIdata(okno strapi.StrapiRestClient, apis []API, jormCommands *jormCmd, exJDB *explorer.ExplorerJDBs, coinSlug, command, variable string) (data map[string]interface{}) {
func (ja *JormAPI) getAPIdata(vars map[string]interface{}) (data map[string]interface{}) {
	fmt.Println("varsvarsvarsvarsvarsvars :   ", vars)
	fmt.Println("coinAPIs :   ", ja.Coin)
	fmt.Println("command :   ", ja.JORMcommands.Command)
	fmt.Println("variable :   ", ja.JORMcommands.Variable)
	if ja.JORMcommands.Variable != "" {
		ja.JORMcommands.Variable = "/" + ja.JORMcommands.Variable
	}
	for _, api := range ja.Apis[ja.Coin] {
		var coinRaw []*coin.Coin
		err := ja.OKNO.Get("coins", ja.Coin, &coinRaw)
		url := api.Endpoint + api.Commands[ja.JORMcommands.Command] + ja.JORMcommands.Variable
		switch api.UrlFormat {
		case "command":
			data = ja.JORMcommands.CMDs[ja.JORMcommands.Command](vars)
		case "endpoint_command":
			url = api.Endpoint + api.Commands[ja.JORMcommands.Command] + ja.JORMcommands.Variable
			err = utl.GetSource(url, &data)
		case "endpoint_coin_command":
			url = api.Endpoint + "/" + ja.Coin + api.Commands[ja.JORMcommands.Command] + ja.JORMcommands.Variable
			err = utl.GetSource(url, &data)
		case "endpoint_command_coin":
			url = api.Endpoint + "/" + api.Commands[ja.JORMcommands.Command] + "/" + ja.Coin + "/" + ja.JORMcommands.Variable
			err = utl.GetSource(url, &data)
		case "endpoint_symbol_command":
			url = api.Endpoint + "/" + strings.ToLower(coinRaw[0].Symbol) + api.Commands[ja.JORMcommands.Command] + ja.JORMcommands.Variable
			err = utl.GetSource(url, &data)
		default:
			url = api.Endpoint + api.Commands[ja.JORMcommands.Command] + ja.JORMcommands.Variable
			err = utl.GetSource(url, &data)
		}
		if err == nil {
			break
		}
	}
	return data
}
