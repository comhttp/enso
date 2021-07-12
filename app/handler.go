package app

import (
	"fmt"
	"github.com/comhttp/enso/app/cfg"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"

	"encoding/json"
	"github.com/comhttp/jorm/coins"

	//"github.com/comhttp/enso/pkg/utl"

	"github.com/tdewolff/minify"
	mjson "github.com/tdewolff/minify/json"
)

type home struct {
	D []coins.Coin
	C coins.Coins
}

// HomeHandler handles a request for (?)
//func HomeHandler(w http.ResponseWriter, r *http.Request) {
//	coins := coin.ReadAllCoins()
//	var bitnoded []coin.Coin
//	for _, coin := range coins.C {
//		if utl.FileExists(cfg.Path + "/jorm/" + coin.Slug + "/info/bitnodes") {
//			bitnoded = append(bitnoded, coin)
//		}
//		//fmt.Println("coin", coin)
//	}
//	data := home{
//		D: bitnoded,
//		C: coins,
//	}
//	tpl.TemplateHandler(cfg.Path+"/templates/comhttp").ExecuteTemplate(w, "base_gohtml", data)
//}

// CoinsHandler handles a request for coin data
func (e *ENSO) CoinHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	out, err := json.Marshal(coins.GetCoin(e.JDB, v["coin"]))
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (e *ENSO) CoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(coins.GetCoins(e.JDB))
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (e *ENSO) restCoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(coins.GetRestCoins(e.JDB))
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (e *ENSO) algoCoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(coins.GetAlgoCoins(e.JDB))
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (e *ENSO) coinsWordsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(coins.GetCoinsWords(e.JDB))
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (e *ENSO) usableCoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(coins.GetUsableCoins(e.JDB))
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (e *ENSO) allCoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(coins.GetAllCoins(e.JDB))
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (e *ENSO) nodeCoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(coins.GetNodeCoins(e.JDB))
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (e *ENSO) coinsBinHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(coins.GetCoinsBin(e.JDB))
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinNodesHandler handles a request for (?)
func (e *ENSO) CoinNodesHandler(w http.ResponseWriter, r *http.Request) {

}

// NodeHandler handles a request for (?)
func (e *ENSO) NodeHandler(w http.ResponseWriter, r *http.Request) {

}

// NodeHandler handles a request for (?)
func (e *ENSO) ViewJSON() http.Handler {
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), mjson.Minify)

	return http.StripPrefix("/j", m.Middleware(http.FileServer(http.Dir(cfg.Path+cfg.C.Out))))
}

// NodeHandler handles a request for (?)
func (e *ENSO) ViewJSONfolder(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	m := minify.New()
	//height, err := strconv.ParseUint(v["file"], 10, 64)
	//if err != nil {
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), mjson.Minify)
	path := v["sec"] + "/" + v["coin"] + "/" + v["type"]
	http.StripPrefix("/e/"+path, m.Middleware(http.FileServer(http.Dir(cfg.Path+"/www/"+path)))).ServeHTTP(w, r)
	//} else {
	//	index := map[uint64]string{}
	//if err := jdb.JDB.Read("/www/data/"+v["coin"]+"/index", v["type"], &index); err != nil {
	//	fmt.Println("Error", err)
	//}

	//fmt.Println("index[height]", index[height])
	//out := map[string]interface{}{}
	//if err := jdb.JDB.Read(cfg.C.Out+"/"+v["coin"]+"/"+v["type"], index[height], &out); err != nil {
	//	fmt.Println("Error", err)
	//}
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), mjson.Minify)
	//json.NewEncoder(w).Encode(out)
	json.NewEncoder(w).Encode("out")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//}
}
