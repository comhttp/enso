package app

import (
	"github.com/comhttp/enso/app/cfg"
	"github.com/comhttp/enso/pkg/utl"
	"github.com/comhttp/jorm/coins"
	"github.com/comhttp/jorm/jdb"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type ENSO struct {
	Coins coins.Coins
	WWW   *http.Server
	JDB   *jdb.JDB
}

func NewENSO() *ENSO {
	err := cfg.CFG.Read("conf", "conf", &cfg.C)
	utl.ErrorLog(err)
	e := &ENSO{
		JDB: jdb.NewJDB(cfg.C.JDBservers),
	}
	e.WWW = &http.Server{
		Handler:      handler(e),
		Addr:         ":" + cfg.C.Port["enso"],
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return e
}

func handler(e *ENSO) http.Handler {
	r := mux.NewRouter()
	s := r.Host("enso.okno.rs").Subrouter()
	s.StrictSlash(true)

	//s.HandleFunc("/", h.HomeHandler)

	//f := s.PathPrefix("/f").Subrouter()
	//f.HandleFunc("/addcoin", h.AddCoinHandler).Methods("POST")
	//f.HandleFunc("/addnode", h.AddNodeHandler).Methods("POST")

	c := s.PathPrefix("/coins").Subrouter()
	c.HandleFunc("/", e.CoinsHandler).Methods("GET")
	c.HandleFunc("/{coin}", e.CoinHandler).Methods("GET")

	i := c.PathPrefix("/info").Subrouter()
	i.HandleFunc("/all", e.allCoinsHandler).Methods("GET")
	i.HandleFunc("/node", e.nodeCoinsHandler).Methods("GET")
	i.HandleFunc("/rest", e.restCoinsHandler).Methods("GET")
	i.HandleFunc("/algo", e.algoCoinsHandler).Methods("GET")
	i.HandleFunc("/words", e.coinsWordsHandler).Methods("GET")
	i.HandleFunc("/usable", e.usableCoinsHandler).Methods("GET")
	i.HandleFunc("/bin", e.coinsBinHandler).Methods("GET")

	//a.HandleFunc("/{coin}/nodes", e.CoinNodesHandler).Methods("GET")
	//a.HandleFunc("/{coin}/{nodeip}", e.NodeHandler).Methods("GET")

	//b := s.PathPrefix("/b").Subrouter()
	//b.HandleFunc("/{coin}/blocks/{per}/{page}", hnd.ViewBlocks).Methods("GET")
	//b.HandleFunc("/{coin}/lastblock", hnd.LastBlock).Methods("GET")
	//b.HandleFunc("/{coin}/block/{id}", hnd.ViewBlock).Methods("GET")
	//b.HandleFunc("/{coin}/tx/{txid}", hnd.ViewTx).Methods("GET")
	//
	//b.HandleFunc("/{coin}/mempool", hnd.ViewRawMemPool).Methods("GET")
	//b.HandleFunc("/{coin}/mining", hnd.ViewMiningInfo).Methods("GET")
	//b.HandleFunc("/{coin}/info", hnd.ViewInfo).Methods("GET")
	//b.HandleFunc("/{coin}/peers", hnd.ViewPeers).Methods("GET")
	//b.HandleFunc("/{coin}/market", hnd.ViewMarket).Methods("GET")

	//j := s.PathPrefix("/j").Subrouter()

	//j.PathPrefix("/").Handler(e.ViewJSON())

	//j.Headers("Access-Control-Allow-Origin", "*")

	//f := s.PathPrefix("/e").Subrouter()
	//e.HandleFunc("/{coin}/blocks/{per}/{page}", h.ViewBlocks).Methods("GET")
	//e.HandleFunc("/{coin}/lastblock", h.LastBlock).Methods("GET")
	//f.HandleFunc("/{sec}/{coin}/{type}/{file}", e.ViewJSONfolder)
	//e.HandleFunc("/{sec}/{coin}/{app}/{type}/{file}", h.ViewJSONfolder)
	//e.HandleFunc("/{coin}/hash/{blockhash}", h.ViewHash).Methods("GET")
	//e.HandleFunc("/{coin}/tx/{txid}", h.ViewTx).Methods("GET")

	//a.HandleFunc("/", o.goodBye).Methods("GET")
	//f.Headers("Access-Control-Allow-Origin", "*")
	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))
}
