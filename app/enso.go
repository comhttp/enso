package app

import (
	"fmt"

	"github.com/comhttp/enso/app/handlers"
	"github.com/comhttp/enso/pkg/utl"
	explorer "github.com/comhttp/explorer/app"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/strapi"
	"github.com/gofiber/fiber/v2"
)

type ENSO struct {
	// Coins    coins.Coins
	Router *fiber.App
	APIs   map[string][]handlers.API
	// Explorer *explorers.Explorer
	jdbServers   map[string]string
	ExJDBs       map[string]*explorer.ExplorerJDBs
	BitNoded     []string
	jormCommands *handlers.JormCmds
	config       cfg.Config
	okno         strapi.StrapiRestClient
}

func NewENSO(path string) *ENSO {
	e := &ENSO{
		ExJDBs: make(map[string]*explorer.ExplorerJDBs),
	}
	//e.Explorer = explorers.GetExplorer(e.JDB)
	e.config.Path = path
	c, _ := cfg.NewCFG(e.config.Path, nil)
	err := c.Read("conf", "conf", &e.config)
	utl.ErrorLog(err)
	jdbServers := make(map[string]string)
	// err = c.Read("conf", "jdbs", &jdbServers)
	// utl.ErrorLog(err)

	// e.ExJDBs = explorer.InitExplorerJDBs(jdbServers, "", "")

	e.okno = strapi.New(e.config.Strapi)

	bitnodes := e.okno.GetAll("nodes", "bitnode=true&")
	bitnodedCoins := make(map[string]int)
	for _, bitnode := range bitnodes {
		bitnodedCoins[bitnode["coin"].(map[string]interface{})["slug"].(string)] = bitnodedCoins[bitnode["coin"].(map[string]interface{})["slug"].(string)] + 1
	}
	jdbs := e.okno.GetAll("services", "type=jdb&")
	for _, jdb := range jdbs {
		jdbServers[jdb["slug"].(string)] = jdb["server"].(map[string]interface{})["tailscale"].(string) + ":" + fmt.Sprint(jdb["port"].(float64))
	}

	for bitnodedCoin, _ := range bitnodedCoins {
		e.ExJDBs[bitnodedCoin] = explorer.InitExplorerJDBs(jdbServers, "", bitnodedCoin)
		e.BitNoded = append(e.BitNoded, bitnodedCoin)
	}
	e.jdbServers = jdbServers
	// e.WWW = &http.Server{
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	// e.WWW = &http.Server{
	// 	Addr:         ":" + e.config.Port["enso"],
	// 	Handler:      handler(e),
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// coinAPIs := make(map[string]interface{})
	// fmt.Println("coinAPIs :   ", coinAPIs)
	e.getAPIs()
	return e
}
