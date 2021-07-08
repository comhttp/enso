package explorer

import (
	"fmt"
	"github.com/comhttp/jorm-server/app/cfg"
	"github.com/comhttp/jorm-server/app/jdb"
	"github.com/comhttp/jorm-server/app/jorm/a"
	"github.com/comhttp/jorm-server/app/jorm/coin"
	"github.com/comhttp/jorm-server/app/jorm/n"
	"github.com/comhttp/jorm-server/pkg/utl"
	"path/filepath"
)

type Explorer struct {
	Status map[string]uint64 `json:"status"`
}

// GetExplorer updates the data from blockchain of a coin in the database
func GetExplorer(j *jdb.JDB, c coin.Coins) {
	var b []string

	for _, coin := range c.C {
		var bn n.BitNoded
		fmt.Println("Coin is BitNode:", coin)
		if utl.FileExists(filepath.FromSlash(cfg.Path + "nodes/" + coin)) {
			b = append(b, coin)
			bitNodes := a.BitNodes{}
			//if err := jdb.JDB.Read(filepath.FromSlash("/nodes/"), coin, &bitNodes); err != nil {
			//	fmt.Println("Error", err)
			//}
			fmt.Println("Coin is BitNode:", coin)
			for _, bitnode := range bitNodes {

				fmt.Println("bitnodebitnodebitnode", cfg.C.RPC.Username, cfg.C.RPC.Password, bitnode.IP, bitnode.Port)

				err := GetBlockchain(bitnode)
				if err != nil {

				}

				bn.Coin = coin
			}
			//bns = append(bns, bn)
			//nodes := jdb.ReadData(filepath.FromSlash(cfg.C.Out + "/nodes"))
			//ns := make(n.Nodes, len(nodes))
			//
			//for i := range nodes {
			//	if err := json.Unmarshal(nodes[i], &ns[i]); err != nil {
			//		fmt.Println("Error", err)
			//	}
			//}
		}
	}

	//et := mod.Cache{Data: e}
	//jdb.JDB.Write(cfg.Web, "explorer", et)
}

// GetExplorer returns the full set of information about a block
func GetBlockchain(b a.BitNode) (err error) {

	fmt.Println("aaaaaaaaa", b)

	if utl.FileExists(cfg.Path + cfg.C.Out + "/info/explorer") {
		e := Explorer{}
		//if err := jdb.JDB.Read(cfg.C.Out+"/info", "explorer", &e); err != nil {
		//	fmt.Println("Error", err)
		//}
		fmt.Println("PortPortPort", b.Port)
		fmt.Println("IPIPIPIP", b.IP)
		blockCount := b.GetBlockCount()

		fmt.Println("blockCountblockCountblockCountblockCount", blockCount)

		if e.Status != nil && blockCount >= int(e.Status["blocks"]) {
			for {
				//e.block(a, www)
				fmt.Println("BlocksPre", e.Status["blocks"])
				blocksIndex := map[uint64]string{}
				//if err := jdb.JDB.Read(cfg.C.Out+"/index", "blocks", &blocksIndex); err != nil {
				//	fmt.Println("Error", err)
				//}
				e.Status["blocks"]++
				blockRaw := b.GetBlockByHeight(int(e.Status["blocks"]))
				if blockRaw != nil && blockRaw != "" {
					blockHash := blockRaw.(map[string]interface{})["hash"].(string)
					blocksIndex[e.Status["blocks"]] = blockHash
					//jdb.JDB.Write(cfg.C.Out+"/blocks", blockHash, blockRaw)
					block := (blockRaw).(map[string]interface{})
					if e.Status["blocks"] != 0 {
						for _, t := range (block["tx"]).([]interface{}) {
							e.tx(b, t.(string))
						}
					}
					fmt.Println("BlocksPosle", e.Status["blocks"])
					//jdb.JDB.Write(cfg.C.Out+"/info", "explorer", e)
					//jdb.JDB.Write(cfg.C.Out+"/index", "blocks", blocksIndex)
				} else {
					break
				}
			}
		}
	} else {
		//e := &Explorer{
		//	Status: map[string]uint64{"blocks": 0, "txs": 0, "addresses": 0},
		//}
		//jdb.JDB.Write(cfg.C.Out+"/info", "explorer", e)
		//jdb.JDB.Write(cfg.C.Out+"/explorer/index", "blocks", map[uint64]string{})
		//jdb.JDB.Write(cfg.C.Out+"/explorer/index", "txs", map[uint64]string{})
		//jdb.JDB.Write(cfg.C.Out+"/explorer/index", "addresses", map[uint64]string{})

		fmt.Println("ExplorerExplorerExplorerExplorer", cfg.C.Out+"/info")

	}
	return
}

func (e *Explorer) tx(a a.BitNode, txid string) {
	txRaw := a.GetTx(txid)
	txsIndex := map[uint64]string{}
	//if err := jdb.JDB.Read(cfg.C.Out+"/index", "txs", &txsIndex); err != nil {
	//	fmt.Println("Error", err)
	//}
	e.Status["txs"]++
	txsIndex[e.Status["txs"]] = txid
	//fmt.Println("txid", txid)
	//go jdb.JDB.Write(cfg.C.Out+"/txs", txid, txRaw)
	if txRaw != nil {
		tx := (txRaw).(map[string]interface{})
		if tx["vout"] != nil {
			for _, nRaw := range tx["vout"].([]interface{}) {
				if nRaw.(map[string]interface{})["scriptPubKey"] != nil {
					scriptPubKey := nRaw.(map[string]interface{})["scriptPubKey"].(map[string]interface{})
					if scriptPubKey["addresses"] != nil {
						for _, address := range (scriptPubKey["addresses"]).([]interface{}) {
							e.addr(cfg.C.Out, address.(string))
						}
					}
				}
			}
		}
	}
	//jdb.JDB.Write(cfg.C.Out+"/index", "txs", txsIndex)
	return
}

func (e *Explorer) addr(www, address string) {
	addressesIndex := map[uint64]string{}
	//if err := jdb.JDB.Read(www+"/index", "addresses", &addressesIndex); err != nil {
	//	fmt.Println("Error", err)
	//}
	//addressData := new(interface{})
	//if err := jdb.JDB.Read(www, "explorer", &e); err != nil {
	//	fmt.Println("Error", err)
	//}
	//addressData := address
	e.Status["addresses"]++
	addressesIndex[e.Status["addresses"]] = address
	//go jdb.JDB.Write(www+"/addresses", address, addressData)
	//jdb.JDB.Write(www+"/index", "addresses", addressesIndex)
	//fmt.Println("address", address)
	return
}

func (e *Explorer) status() {
	//s =
}
