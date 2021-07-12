package n

import (
	"fmt"
	"github.com/comhttp/enso/app/jorm/a"
	"github.com/comhttp/jorm/coins"
	"path/filepath"

	"github.com/comhttp/enso/app/cfg"
	"github.com/comhttp/enso/pkg/utl"
)

// BitNoded data
type BitNoded struct {
	Coin     string          `json:"coin"`
	BitNodes []BitNodeStatus `json:"bitnodes"`
}

// NodeStatus stores current data for a node
type BitNodeStatus struct {
	Live           bool        `json:"live"`
	IP             string      `json:"ip"`
	GetInfo        interface{} `json:"getinfo"`
	GetPeerInfo    interface{} `json:"getpeerinfo"`
	GetRawMemPool  interface{} `json:"getrawmempool"`
	GetMiningInfo  interface{} `json:"getmininginfo"`
	GetNetworkInfo interface{} `json:"getnetworkinfo"`
	GeoIP          interface{} `json:"geoip"`
}

type Nodes []NodeInfo

// NodeInfo stores info retrieved via geoip about a node
type NodeInfo struct {
	IP            string  `json:"ip"`
	Port          int64   `json:"port"`
	Host          string  `json:"host"`
	Rdns          string  `json:"rdns"`
	ASN           string  `json:"asn"`
	ISP           string  `json:"isp"`
	CountryName   string  `json:"countryname"`
	CountryCode   string  `json:"countrycode"`
	RegionName    string  `json:"regionname"`
	RegionCode    string  `json:"regioncode"`
	City          string  `json:"city"`
	Postcode      string  `json:"postcode"`
	ContinentName string  `json:"continentname"`
	ContinentCode string  `json:"continentcode"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Zipcode       string  `json:"zipcode"`
	Timezone      string  `json:"timezone"`
	LastSeen      string  `json:"lastseen"`
	Live          bool    `json:"live"`
}

// GetBitNodes updates the data about all of the coins in the database
func GetBitNodes(coins coins.Coins) {
	var b []string
	bns := make(map[string]BitNoded)
	for _, coin := range coins.C {
		var bn BitNoded
		if utl.FileExists(filepath.FromSlash(cfg.Path + "nodes/" + coin)) {

			b = append(b, coin)
			bitNodes := a.BitNodes{}
			if err := cfg.CFG.Read("nodes", coin, &bitNodes); err != nil {
				fmt.Println("Error", err)
			}
			for _, bitnode := range bitNodes {
				bitNode := GetBitNodeStatus(bitnode)
				nds := GetNodes(bitNode)
				for _, n := range nds {
					if n.IP[:3] == "10." {
						n.IP = "212.62.35.158"
					}
					//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/nodes/"+coin), n.IP, n)
				}
				if bitnode.IP[:3] == "10." {
					bitnode.IP = "212.62.35.158"
				}
				//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/bitnodes/"+coin), bitnode.IP, bitNode)
				//
				//fmt.Println("--------------------")
				//fmt.Println("bitNodes", nds)
				//fmt.Println("--------------------")

				//bn.Coin = coin
				bn.BitNodes = append(bn.BitNodes, *bitNode)
			}
			bns[coin] = bn
			//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/info/nodes/"+coin), "bitnodes", bn)

			//data, err := jdb.JDB.ReadAll(filepath.FromSlash(cfg.C.Out + "/nodes/" + coin))
			//utl.ErrorLog(err)
			//nodes := make([][]byte, len(data))
			//for i := range data {
			//	nodes[i] = []byte(data[i])
			//}

			//ns := make(Nodes, len(nodes))
			//
			//for i := range nodes {
			//	if err := json.Unmarshal(nodes[i], &ns[i]); err != nil {
			//		fmt.Println("Error", err)
			//	}
			//}
			//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/info/nodes/"+coin), "nodes", ns)
		}
	}

	//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/info"), "bitnoded", b)
	//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/info"), "bitnodestat", bns)
}
