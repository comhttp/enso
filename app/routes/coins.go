package routes

import (
	"github.com/comhttp/enso/app/handlers"
	explorer "github.com/comhttp/explorer/app"
	"github.com/comhttp/jorm/mod/coin"
	"github.com/comhttp/jorm/pkg/strapi"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/gofiber/fiber/v2"
)

func CoinsRoutes(cqs *coin.CoinsQueries, okno strapi.StrapiRestClient, apis map[string][]handlers.API, jormCommands *handlers.JormCmds, exJDBs map[string]*explorer.ExplorerJDBs, r *fiber.App) {
	ja := &handlers.JormAPI{
		OKNO: okno,
		Apis: apis,
		// JORMcommands: jormCommands,
		ExJDB: exJDBs,
		CQ:    cqs,
		// Coin:         c.Params("coin"),
		JORMcommands: &handlers.JormCmds{
			CMDs: jormCommands.CMDs,
			// Command: "lastblock",
		},
	}

	app := r.Group("/system")
	app.Get("/dashboard", monitor.New())

	jormCoins := r.Group("/coins")
	jormCoins.Get("/", handlers.CoinsHandler(ja.CQ))

	jormCoins.Get("/all", handlers.AllCoinsHandler(ja.CQ))
	jormCoins.Get("/node", handlers.NodeCoinsHandler(ja.CQ))
	jormCoins.Get("/rest", handlers.RestCoinsHandler(ja.CQ))
	jormCoins.Get("/algo", handlers.AlgoCoinsHandler(ja.CQ))
	jormCoins.Get("/words", handlers.CoinsWordsHandler(ja.CQ))
	jormCoins.Get("/usable", handlers.UsableCoinsHandler(ja.CQ))
	jormCoins.Get("/bin", handlers.CoinsBinHandler(ja.CQ))

	jormCoin := r.Group("/coin")
	jormSingleCoin := jormCoin.Group("/:coin")
	jormSingleCoin.Get("/", handlers.CoinHandler(ja.CQ))
	jormSingleCoin.Get("/logo/:size", handlers.LogoHandler(ja.CQ))

	jormNodes := jormSingleCoin.Group("/nodes")
	// jormNodes.Get("/", handlers.CoinNodesHandler(cq))
	jormNodes.Get("/:ip", handlers.NodeHandler(ja.CQ))
	// n := r.PathPrefix("/nodes").Subrouter()
	// //n.HandleFunc("/{coin}/nodes", cq.CoinNodesHandler).Methods("GET")
	// n.HandleFunc("/{coin}/{nodeip}", cq.nodeHandler).Methods("GET")

	jsonCoin := jormCoins.Group("/json")
	jsonCoin.Get("/algo", handlers.JsonAlgoCoinsHandler(ja.CQ))

	jormExplorer := jormSingleCoin.Group("/chain")

	jormExplorer.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})
	jormExplorer.Get("/status", ja.ViewStatus())
	jormExplorer.Get("/lastblock", ja.LastBlock())
	jormExplorer.Get("/:type/:per/:page", ja.ViewTypes())

	jormExplorer.Get("/:type/:id", ja.ViewType())

	return
}
