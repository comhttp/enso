package routes

// func exchangesRoutes(eq *ExchangeQueries, r *mux.Router) *mux.Router {
// 	//cq := j.CollectionQueries("coin").(CoinsQueries)
// 	//s := r.Host("enso.okno.rs").Subrouter()
// 	r.StrictSlash(true)

// 	//s.HandleFunc("/", h.HomeHandler)

// 	//f := s.PathPrefix("/f").Subrouter()
// 	//f.HandleFunc("/addcoin", h.AddCoinHandler).Methods("POST")
// 	//f.HandleFunc("/addnode", h.AddNodeHandler).Methods("POST")

// 	e := r.PathPrefix("/exchanges").Subrouter()
// 	e.HandleFunc("/", eq.ExchangesHandler).Methods("GET")
// 	e.HandleFunc("/{exchange}", eq.ExchangeHandler).Methods("GET")
// 	e.HandleFunc("/{exchange}/markets", eq.MarketsHandler).Methods("GET")
// 	e.HandleFunc("/{exchange}/markets/{market}", eq.MarketHandler).Methods("GET")

// 	return r
// }
