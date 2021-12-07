package routes

import (
	"github.com/comhttp/enso/app/handlers"
	"github.com/comhttp/jorm/mod/coin"
	"github.com/gofiber/fiber/v2"
)

func CoinsRoutes(cq *coin.CoinsQueries, r *fiber.App) {
	jormCoins := r.Group("/coins")
	jormCoins.Get("/", handlers.CoinsHandler(cq))

	jormCoins.Get("/all", handlers.AllCoinsHandler(cq))
	jormCoins.Get("/node", handlers.NodeCoinsHandler(cq))
	jormCoins.Get("/rest", handlers.RestCoinsHandler(cq))
	jormCoins.Get("/algo", handlers.AlgoCoinsHandler(cq))
	jormCoins.Get("/words", handlers.CoinsWordsHandler(cq))
	jormCoins.Get("/usable", handlers.UsableCoinsHandler(cq))
	jormCoins.Get("/bin", handlers.CoinsBinHandler(cq))

	jormCoin := r.Group("/coin")
	jormSingleCoin := jormCoin.Group("/:coin")
	jormSingleCoin.Get("/", handlers.CoinHandler(cq))
	jormSingleCoin.Get("/logo/:size", handlers.LogoHandler(cq))

	jormNodes := jormSingleCoin.Group("/nodes")
	// jormNodes.Get("/", handlers.CoinNodesHandler(cq))
	jormNodes.Get("/:ip", handlers.NodeHandler(cq))
	// n := r.PathPrefix("/nodes").Subrouter()
	// //n.HandleFunc("/{coin}/nodes", cq.CoinNodesHandler).Methods("GET")
	// n.HandleFunc("/{coin}/{nodeip}", cq.nodeHandler).Methods("GET")

	jsonCoin := jormCoins.Group("/json")
	jsonCoin.Get("/algo", handlers.JsonAlgoCoinsHandler(cq))

	return
}
