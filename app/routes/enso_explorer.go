package routes

// func Queries(j map[string]*jdb.JDB, col string) *ExplorerQueries {
// 	return &ExplorerQueries{
// 		j,
// 		&BlockchainStatus{},
// 		col,
// 	}
// }

// func ENSOroutes(eq *ExplorerQueries, r *mux.Router) *mux.Router {
// 	//info := Queries(j, "info")
// 	//info := Queries(j.JDBclient("explorer"),"info")
// 	//r.StrictSlash(true)
// 	//n := s.PathPrefix("/n").Subrouter()
// 	//n.HandleFunc("/{coin}/nodes", explorersCollection.CoinNodesHandler).Methods("GET")
// 	//n.HandleFunc("/{coin}/{nodeip}", explorersCollection.nodeHandler).Methods("GET")

// 	b := r.PathPrefix("/explorer").Subrouter()
// 	b.HandleFunc("/{coin}/status", eq.ViewStatus).Methods("GET")
// 	b.HandleFunc("/{coin}/blocks/{per}/{page}", eq.ViewBlocks).Methods("GET")
// 	b.HandleFunc("/{coin}/lastblock", eq.LastBlock).Methods("GET")
// 	b.HandleFunc("/{coin}/block/{id}", eq.ViewBlock).Methods("GET")
// 	b.HandleFunc("/{coin}/tx/{txid}", eq.ViewTx).Methods("GET")
// 	b.HandleFunc("/{coin}/mempool", eq.ViewRawMemPool).Methods("GET")
// 	b.HandleFunc("/{coin}/mining", eq.ViewMiningInfo).Methods("GET")
// 	b.HandleFunc("/{coin}/info", eq.ViewInfo).Methods("GET")
// 	b.HandleFunc("/{coin}/peers", eq.ViewPeers).Methods("GET")
// 	//b.HandleFunc("/{coin}/market", explorersCollection.ViewMarket).Methods("GET")
// 	return r
// }
