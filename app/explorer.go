package app

import (
	"encoding/json"
	"fmt"
	"github.com/comhttp/jorm/mod/explorers"
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
	ex := explorers.GetExplorer(e.JDB,v["coin"])
	lastblock := ex.Blocks - 1
	fmt.Println("lastblocklastblocklastblock", lastblock)

	lb := map[string]interface{}{
		"currentPage": page,
		"pageCount":   lastblock / per,
		"blocks":      e.Explorer.GetBlocks(e.JDB, v["coin"], per, page),
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
	out, err := json.Marshal(explorers.GetBlock(e.JDB, v["coin"], v["id"]))
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
	out, err := json.Marshal(explorers.GetTx(e.JDB, v["coin"], v["txid"]))
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (e *ENSO) ViewAddr(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	var block interface{}
	block = explorers.GetBlock(e.JDB, v["coin"], v["id"])
	out, err := json.Marshal(block)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (e *ENSO) ViewRawMemPool(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	rawMemPool := explorers.GetMemPool(e.JDB, v["coin"])
	out, err := json.Marshal(rawMemPool)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (e *ENSO) ViewMiningInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	miningInfo := explorers.GetMiningInfo(e.JDB, v["coin"])

	out, err := json.Marshal(miningInfo)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (e *ENSO) ViewInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	info := explorers.GetInfo(e.JDB, v["coin"])
	out, err := json.Marshal(info)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (e *ENSO) ViewPeers(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	peers := explorers.GetPeers(e.JDB, v["coin"])
	out, err := json.Marshal(peers)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
