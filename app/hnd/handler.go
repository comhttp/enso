package hnd

import (
	"fmt"
	"github.com/comhttp/jorm-server/app/cfg"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"strconv"

	"encoding/json"
	"github.com/comhttp/jorm-server/app/jorm/a"
	"github.com/comhttp/jorm-server/app/jorm/coin"

	//"github.com/comhttp/jorm-server/pkg/utl"

	"github.com/tdewolff/minify"
	mjson "github.com/tdewolff/minify/json"
)

type home struct {
	D []coin.Coin
	C coin.Coins
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

// ChatHandler handles a request for (?)
func ChatHandler(w http.ResponseWriter, r *http.Request) {
	//data := home{}
	//tpl.TemplateHandler(cfg.Path+"/templates/chat").ExecuteTemplate(w, "chat_gohtml", data)
}

// AddCoinHandler handles a request for adding coin data
func AddCoinHandler(w http.ResponseWriter, r *http.Request) {
	//name := r.FormValue("coin")
	//coin := coin.Coin{
	//	Name: name,
	//	Slug: utl.MakeSlug(name),
	//}

	//fmt.Println("name", name)
	//fmt.Println("coin", coin)

	//jdb.DB.Write(cfg.Web+"/coins", coin.Slug, coin)
	http.Redirect(w, r, "/", 302)
}

// AddNodeHandler handles a request for adding node data
func AddNodeHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.FormValue("ip")
	p := r.FormValue("port")
	//c := utl.MakeSlug(r.FormValue("coin"))

	var bitNodes a.BitNodes
	//jdb.DB.Read(c, "bitnodes", &bitNodes)

	port, err := strconv.ParseInt(p, 10, 64)
	if err == nil {
		// What is this supposed to be printing exactly?
		// fmt.Printf("%d of type %T", p, p)
	}

	fmt.Println("ip", ip)
	fmt.Println("port", port)

	bitNode := a.BitNode{
		IP:   ip,
		Port: port,
	}

	bitNodes = append(bitNodes, bitNode)

	//jdb.DB.Write(c, "bitnodes", bitNodes)
	http.Redirect(w, r, "/", 302)
}

// CoinsHandler handles a request for coin data
func CoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	//if err != nil {
	//	fmt.Println("Error encoding JSON")
	//	return
	//}
	//w.Write([]byte(out))
}

// CoinNodesHandler handles a request for (?)
func CoinNodesHandler(w http.ResponseWriter, r *http.Request) {

}

// NodeHandler handles a request for (?)
func NodeHandler(w http.ResponseWriter, r *http.Request) {

}

// NodeHandler handles a request for (?)
func ViewJSON() http.Handler {
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), mjson.Minify)

	return http.StripPrefix("/j", m.Middleware(http.FileServer(http.Dir(cfg.Path+cfg.C.Out))))
}

// NodeHandler handles a request for (?)
func ViewJSONfolder(w http.ResponseWriter, r *http.Request) {
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
