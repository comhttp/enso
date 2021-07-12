package a

import (
	"fmt"

	"github.com/comhttp/enso/app/cfg"
	"github.com/comhttp/enso/pkg/utl"
)

// BitNodes is array of bitnodes addresses
type BitNodes []BitNode

// BitNodeSrc is a node's address
type BitNode struct {
	IP   string `json:"ip"`
	Port int64  `json:"p"`
	Jrc  *utl.Endpoint
}

func RPCSRC(c string) (b *BitNode) {
	bitNodes := BitNodes{}
	if err := cfg.CFG.Read("nodes", c, &bitNodes); err != nil {
		fmt.Println("Errdor", err)
	}
	for _, bn := range bitNodes {
		b = &bn
		b.Jrc = utl.NewClient(cfg.C.RPC.Username, cfg.C.RPC.Password, b.IP, b.Port)
		if b.Jrc != nil {
			fmt.Println("bitb:", b.IP)
			break
		}

	}
	fmt.Println("b:", b)
	return
}
