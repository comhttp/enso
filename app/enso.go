package app

import (
	"github.com/comhttp/enso/pkg/utl"
	"github.com/comhttp/jorm/mod/coins"
	"github.com/comhttp/jorm/mod/explorers"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type ENSO struct {
	Coins    coins.Coins
	WWW      *http.Server
	JDB      *jdb.JDB
	Explorer *explorers.Explorer
}

func NewENSO() *ENSO {
	err := cfg.CFG.Read("conf", "conf", &cfg.C)
	utl.ErrorLog(err)
	e := &ENSO{
		JDB: jdb.NewJDB(cfg.C.JDBservers),
	}
	//e.Explorer = explorers.GetExplorer(e.JDB)
	e.WWW = &http.Server{
		Handler:      handler(e),
		Addr:         ":" + cfg.C.Port["enso"],
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return e
}

func handler(e *ENSO) http.Handler {
	s := mux.NewRouter()
	//s := r.Host("enso.okno.rs").Subrouter()
	s.StrictSlash(true)

	//s.HandleFunc("/", h.HomeHandler)

	//f := s.PathPrefix("/f").Subrouter()
	//f.HandleFunc("/addcoin", h.AddCoinHandler).Methods("POST")
	//f.HandleFunc("/addnode", h.AddNodeHandler).Methods("POST")

	c := s.PathPrefix("/coins").Subrouter()
	c.HandleFunc("/", e.CoinsHandler).Methods("GET")
	c.HandleFunc("/{coin}", e.CoinHandler).Methods("GET")

	i := s.PathPrefix("/info").Subrouter()
	i.HandleFunc("/all", e.allCoinsHandler).Methods("GET")
	i.HandleFunc("/node", e.nodeCoinsHandler).Methods("GET")
	i.HandleFunc("/rest", e.restCoinsHandler).Methods("GET")
	i.HandleFunc("/algo", e.algoCoinsHandler).Methods("GET")
	i.HandleFunc("/words", e.coinsWordsHandler).Methods("GET")
	i.HandleFunc("/usable", e.usableCoinsHandler).Methods("GET")
	i.HandleFunc("/bin", e.coinsBinHandler).Methods("GET")

	n := s.PathPrefix("/n").Subrouter()
	//n.HandleFunc("/{coin}/nodes", e.CoinNodesHandler).Methods("GET")
	n.HandleFunc("/{coin}/{nodeip}", e.nodeHandler).Methods("GET")

	b := s.PathPrefix("/e").Subrouter()
	b.HandleFunc("/{coin}/status", e.ViewStatus).Methods("GET")
	b.HandleFunc("/{coin}/blocks/{per}/{page}", e.ViewBlocks).Methods("GET")
	//b.HandleFunc("/{coin}/lastblock", hnd.LastBlock).Methods("GET")
	b.HandleFunc("/{coin}/block/{id}", e.ViewBlock).Methods("GET")
	b.HandleFunc("/{coin}/tx/{txid}", e.ViewTx).Methods("GET")
	//
	b.HandleFunc("/{coin}/mempool", e.ViewRawMemPool).Methods("GET")
	b.HandleFunc("/{coin}/mining", e.ViewMiningInfo).Methods("GET")
	b.HandleFunc("/{coin}/info", e.ViewInfo).Methods("GET")
	b.HandleFunc("/{coin}/peers", e.ViewPeers).Methods("GET")
	//b.HandleFunc("/{coin}/market", e.ViewMarket).Methods("GET")

	//j := s.PathPrefix("/j").Subrouter()

	//j.PathPrefix("/").Handler(e.ViewJSON())

	//j.Headers("Access-Control-Allow-Origin", "*")

	//f.HandleFunc("/{sec}/{coin}/{type}/{file}", e.ViewJSONfolder)
	//e.HandleFunc("/{sec}/{coin}/{app}/{type}/{file}", h.ViewJSONfolder)

	//a.HandleFunc("/", o.goodBye).Methods("GET")
	//f.Headers("Access-Control-Allow-Origin", "*")
	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(s, utl.DefaultErrorHandler)))
}
