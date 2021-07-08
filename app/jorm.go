package app

import (
	"github.com/comhttp/jorm-server/app/cfg"
	"github.com/comhttp/jorm-server/app/jorm/coin"
	"github.com/comhttp/jorm-server/app/hnd"
	"github.com/comhttp/jorm-server/pkg/utl"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"net/http"
	"time"
)

type JORM struct {
	Coins coin.Coins
	WWW         *http.Server
}

func NewJORM() *JORM {
	err := cfg.CFG.Read("conf", "conf", &cfg.C)
	utl.ErrorLog(err)
	j := &JORM{
	}
	j.WWW = &http.Server{
		Handler:      handler(),
		Addr:         ":" + cfg.C.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return j
}

func handler() http.Handler {
	r := mux.NewRouter()
	s := r.Host("jorm.okno.rs").Subrouter()
	s.StrictSlash(true)

	//s.HandleFunc("/", h.HomeHandler)

	//f := s.PathPrefix("/f").Subrouter()
	//f.HandleFunc("/addcoin", h.AddCoinHandler).Methods("POST")
	//f.HandleFunc("/addnode", h.AddNodeHandler).Methods("POST")

	a := s.PathPrefix("/a").Subrouter()
	a.HandleFunc("/coins", hnd.CoinsHandler).Methods("GET")
	a.HandleFunc("/{coin}/nodes", hnd.CoinNodesHandler).Methods("GET")
	a.HandleFunc("/{coin}/{nodeip}", hnd.NodeHandler).Methods("GET")

	b := s.PathPrefix("/b").Subrouter()
	b.HandleFunc("/{coin}/blocks/{per}/{page}", hnd.ViewBlocks).Methods("GET")
	b.HandleFunc("/{coin}/lastblock", hnd.LastBlock).Methods("GET")
	b.HandleFunc("/{coin}/block/{id}", hnd.ViewBlock).Methods("GET")
	b.HandleFunc("/{coin}/tx/{txid}", hnd.ViewTx).Methods("GET")

	b.HandleFunc("/{coin}/mempool", hnd.ViewRawMemPool).Methods("GET")
	b.HandleFunc("/{coin}/mining", hnd.ViewMiningInfo).Methods("GET")
	b.HandleFunc("/{coin}/info", hnd.ViewInfo).Methods("GET")
	b.HandleFunc("/{coin}/peers", hnd.ViewPeers).Methods("GET")
	//b.HandleFunc("/{coin}/market", hnd.ViewMarket).Methods("GET")

	j := s.PathPrefix("/j").Subrouter()

	j.PathPrefix("/").Handler(hnd.ViewJSON())

	j.Headers("Access-Control-Allow-Origin", "*")

	e := s.PathPrefix("/e").Subrouter()
	//e.HandleFunc("/{coin}/blocks/{per}/{page}", h.ViewBlocks).Methods("GET")
	//e.HandleFunc("/{coin}/lastblock", h.LastBlock).Methods("GET")
	e.HandleFunc("/{sec}/{coin}/{type}/{file}", hnd.ViewJSONfolder)
	//e.HandleFunc("/{sec}/{coin}/{app}/{type}/{file}", h.ViewJSONfolder)
	//e.HandleFunc("/{coin}/hash/{blockhash}", h.ViewHash).Methods("GET")
	//e.HandleFunc("/{coin}/tx/{txid}", h.ViewTx).Methods("GET")

	//a.HandleFunc("/", o.goodBye).Methods("GET")
	e.Headers("Access-Control-Allow-Origin", "*")
	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))
}
