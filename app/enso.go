package app

import (
	"fmt"

	"github.com/comhttp/enso/app/handlers"
	"github.com/comhttp/enso/app/routes"
	"github.com/comhttp/enso/pkg/utl"
	explorer "github.com/comhttp/explorer/app"
	"github.com/comhttp/jorm/mod/coin"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/strapi"
	"github.com/gofiber/fiber/v2"
)

type ENSO struct {
	// Coins    coins.Coins
	Router *fiber.App
	APIs   map[string][]handlers.API
	// Explorer *explorers.Explorer
	jdbServers map[string]string
	ExJDBs     map[string]*explorer.ExplorerJDBs

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

func (e *ENSO) ENSOrouter() {
	e.Router = fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "OKNO",
		AppName:       "ENSO",
	})

	coins, err := jdb.NewJDB(e.jdbServers["jdbcoins"])
	utl.ErrorLog(err)

	cq := coin.Queries(coins, "")

	routes.CoinsRoutes(cq, e.okno, e.APIs, e.getJORMcommands(), e.ExJDBs, e.Router)

	// routes.Explorer(e.okno, e.APIs, e.getJORMcommands(), e.ExJDBs, e.Router)
	//s := r.Host("enso.okno.rs").Subrouter()
	// r.StrictSlash(true)
	// return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))

	// log.Fatal().Err(e.Router.Listen(":3000"))
	// log.Fatal(e.Router.Listen(":3000"))

}

func (e *ENSO) getAPIs() {
	e.APIs = make(map[string][]handlers.API)

	apis := e.okno.GetAll("apis", "")
	for _, api := range apis {
		coinsRaw := api["coin"].([]interface{})
		if len(coinsRaw) > 0 {
			for _, coinRaw := range coinsRaw {
				coin := coinRaw.(map[string]interface{})["slug"]
				apiCommands := make(map[string]string)
				for _, command := range api["commands"].([]interface{}) {
					cmd := command.(map[string]interface{})
					apiCommands[cmd["type"].(string)] = cmd["command"].(string)
				}
				e.APIs[coin.(string)] = append(e.APIs[coin.(string)],
					handlers.API{
						UrlFormat: api["urlformat"].(string),
						Commands:  apiCommands,
						Endpoint:  api["url"].(string),
					})
			}
		}
	}
	return
}

func (e *ENSO) getJORMcommands() *handlers.JormCmds {
	e.jormCommands = &handlers.JormCmds{}
	e.jormCommands.CMDs = make(map[string](func(vars map[string]interface{}) map[string]interface{}))
	e.jormCommands.CMDs["status"] = func(vars map[string]interface{}) map[string]interface{} {
		status := e.ExJDBs[vars["coin"].(string)].GetExplorer(vars["coin"].(string))
		d := make(map[string]interface{})
		d["blocks"] = status.Blocks
		d["addresses"] = status.Addresses
		d["txs"] = status.Txs
		return d
	}
	e.jormCommands.CMDs["lastblock"] = func(vars map[string]interface{}) map[string]interface{} {
		lastblock := e.ExJDBs[vars["coin"].(string)].GetLastBlock(vars["coin"].(string))
		d := make(map[string]interface{})
		d["lastblock"] = lastblock
		return d
	}
	e.jormCommands.CMDs["blocks"] = func(vars map[string]interface{}) map[string]interface{} {
		blocks := e.ExJDBs[vars["coin"].(string)].GetBlocks(vars["coin"].(string), vars["per"].(int), vars["page"].(int))
		d := make(map[string]interface{})
		d["blocks"] = blocks
		return d
	}

	e.jormCommands.CMDs["block"] = func(vars map[string]interface{}) map[string]interface{} {
		block := e.ExJDBs[vars["coin"].(string)].GetBlock(vars["coin"].(string), vars["id"].(string))
		return block
	}
	// e.jormCommands["blocks"] = func(vars map[string]interface{}) map[string]interface{} {
	// 	lastblock := e.ExJDBs[vars["coin"]].ViewBlocks(vars["coin"])
	// 	d := make(map[string]interface{})
	// 	d["lastblock"] = lastblock
	// 	return d
	// }
	return e.jormCommands
}
