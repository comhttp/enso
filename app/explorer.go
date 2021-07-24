package app

import (
	"encoding/json"
	"fmt"
	"github.com/comhttp/jorm/mod/explorer"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (e *ENSO) ViewStatus(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	out, err := json.Marshal(e.Explorer.Status[v["coin"]])
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (e *ENSO) ViewBlocks(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	per, _ := strconv.Atoi(v["per"])
	page, _ := strconv.Atoi(v["page"])
	ex := explorer.GetExplorer(e.JDB)
	lastblock := ex.Status[v["coin"]].Blocks - 1
	fmt.Println("lastblocklastblocklastblock", lastblock)

	lb := map[string]interface{}{
		"currentPage": page,
		"pageCount":   lastblock / per,
		"blocks":      e.Explorer.GetBlocks(v["coin"], per, page),
		"lastBlock":   lastblock,
	}

	out, err := json.Marshal(lb)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (e *ENSO) LastBlock(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	out, err := json.Marshal(e.Explorer.Status[v["coin"]].Blocks)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (e *ENSO) ViewBlock(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	out, err := json.Marshal(e.Explorer.GetBlock(v["coin"], v["id"]))
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

//func ViewBlockHeight(w http.ResponseWriter, r *http.Request) {
//	v := mux.Vars(r)
//	bh := v["blockheight"]
//	// node := Node{}
//	bhi, _ := strconv.Atoi(bh)
//	block := a.RPCSRC(v["coin"]).GetBlockByHeight(bhi)
//	out, err := json.Marshal(block)
//	if err != nil {
//		fmt.Println("Error encoding JSON")
//		return
//	}
//	w.Write([]byte(out))
//}
//
//func ViewHash(w http.ResponseWriter, r *http.Request) {
//	v := mux.Vars(r)
//	bh := v["blockhash"]
//	block := (a.RPCSRC(v["coin"]).GetBlock(bh)).(map[string]interface{})
//	h := strconv.FormatInt(block["height"].(int64), 10)
//	http.Redirect(w, r, "/b/"+v["coin"]+"/block/"+h, 301)
//}

func (e *ENSO) ViewTx(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	out, err := json.Marshal(e.Explorer.GetTx(v["coin"], v["txid"]))
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (e *ENSO) ViewAddr(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	var block interface{}
	block = e.Explorer.GetBlock(v["coin"], v["id"])
	out, err := json.Marshal(block)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (e *ENSO) ViewRawMemPool(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	rawMemPool := e.Explorer.GetMemPool(v["coin"])
	out, err := json.Marshal(rawMemPool)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (e *ENSO) ViewMiningInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	miningInfo := e.Explorer.GetMiningInfo(v["coin"])

	out, err := json.Marshal(miningInfo)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (e *ENSO) ViewInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	info := e.Explorer.GetInfo(v["coin"])
	out, err := json.Marshal(info)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (e *ENSO) ViewPeers(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	peers := e.Explorer.GetPeers(v["coin"])
	out, err := json.Marshal(peers)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
