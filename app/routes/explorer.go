package routes

// func ExplorerRoutes(okno strapi.StrapiRestClient, apis map[string][]handlers.API, jormCommands *handlers.JormCmds, exJDBs map[string]*explorer.ExplorerJDBs, r *fiber.App) {

// 	ja := &handlers.JormAPI{
// 		OKNO: okno,
// 		Apis: apis,
// 		// JORMcommands: jormCommands,
// 		ExJDB: exJDBs,
// 		// Coin:         c.Params("coin"),
// 		JORMcommands: &handlers.JormCmds{
// 			CMDs: jormCommands.CMDs,
// 			// Command: "lastblock",
// 		},
// 	}

// 	jormExplorer := r.Group("/explorers")

// 	jormExplorer.Get("/", func(c *fiber.Ctx) error {
// 		err := c.SendString("And the API is UP!")
// 		return err
// 	})
// 	jormExplorerCoin := jormExplorer.Group("/:coin")

// 	jormExplorerCoin.Get("/status", ja.ViewStatus())
// 	jormExplorerCoin.Get("/lastblock", ja.LastBlock())
// 	jormExplorerCoin.Get("/:type/:per/:page", ja.ViewTypes())

// 	jormExplorerCoin.Get("/:type/:id", ja.ViewType())

// 	//info := Queries(j, "info")
// 	//info := Queries(j.JDBclient("explorer"),"info")
// 	//r.StrictSlash(true)
// 	//n := s.PathPrefix("/n").Subrouter()
// 	//n.HandleFunc("/{coin}/nodes", explorersCollection.CoinNodesHandler).Methods("GET")
// 	//n.HandleFunc("/{coin}/{nodeip}", explorersCollection.nodeHandler).Methods("GET")

// 	// b := r.PathPrefix("/explorer").Subrouter()
// 	// b.HandleFunc("/{coin}/status", eq.ViewStatus).Methods("GET")
// 	// b.HandleFunc("/{coin}/blocks/{per}/{page}", eq.ViewBlocks).Methods("GET")
// 	// b.HandleFunc("/{coin}/lastblock", eq.LastBlock).Methods("GET")
// 	// b.HandleFunc("/{coin}/block/{id}", eq.ViewBlock).Methods("GET")
// 	// b.HandleFunc("/{coin}/tx/{txid}", eq.ViewTx).Methods("GET")
// 	// b.HandleFunc("/{coin}/mempool", eq.ViewRawMemPool).Methods("GET")
// 	// b.HandleFunc("/{coin}/mining", eq.ViewMiningInfo).Methods("GET")
// 	// b.HandleFunc("/{coin}/info", eq.ViewInfo).Methods("GET")
// 	// b.HandleFunc("/{coin}/peers", eq.ViewPeers).Methods("GET")
// 	//b.HandleFunc("/{coin}/market", explorersCollection.ViewMarket).Methods("GET")
// 	// return r
// }
