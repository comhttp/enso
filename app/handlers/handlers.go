package handlers

import (
	"fmt"
	"strconv"

	explorer "github.com/comhttp/explorer/app"
	"github.com/comhttp/jorm/pkg/strapi"

	"github.com/gofiber/fiber/v2"
)

func (ja *JormAPI) ViewStatus() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(ja.ExJDB[c.Params("coin")].GetExplorer(c.Params("coin")))
	}
}

func (ja *JormAPI) ViewBlocks(okno strapi.StrapiRestClient, apis map[string][]API, jormCommands *JormCmds, exJDBs map[string]*explorer.ExplorerJDBs) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var blockCount int
		per, _ := strconv.Atoi(c.Params("per"))
		page, _ := strconv.Atoi(c.Params("page"))
		ja := &JormAPI{
			JORMcommands: &JormCmds{
				// CMDs: ,
				Command: "lastblock",
			},
		}
		data := make(map[string]interface{})
		vars := map[string]interface{}{
			"coin": ja.Coin,
			"per":  per,
			"page": page,
		}
		rawData := ja.getAPIdata(vars)
		if rawData != nil {
			if rawData["height"] != nil {
				blockCount = int(rawData["height"].(float64))
			}
		}
		startBlock := blockCount - per*page
		minusBlockStart := int(startBlock + per)
		data["explorer"] = exJDBs[c.Params("coin")].GetExplorer(c.Params("coin"))
		var blocks []map[string]interface{}
		for ibh := minusBlockStart; ibh >= startBlock; {
			ja.JORMcommands.Command = "block"
			ja.JORMcommands.Variable = fmt.Sprint(ibh)
			vars := map[string]interface{}{
				"coin":   ja.Coin,
				"height": fmt.Sprint(ibh),
			}
			apiDataa := ja.getAPIdata(vars)
			blocks = append(blocks, apiDataa)
			ibh--
		}
		out := map[string]interface{}{
			"currentPage": page,
			"pageCount":   blockCount / per,
			"blocks":      blocks,
		}
		return c.JSON(out)
	}
}

func (ja *JormAPI) LastBlock() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		vars := map[string]interface{}{
			"coin": c.Params("coin"),
		}
		ja.Coin = c.Params("coin")
		ja.JORMcommands.Command = "lastblock"
		return c.JSON(ja.getAPIdata(vars))
	}
}

// func (b *bitNodesRPC) directViewBlockHeight() {
// 	v := mux.Vars(r)
// 	bh := v["blockheight"]
// 	// node := Node{}
// 	rpc := RPCSRC(bitNodes(b.path, v["coin"]), b.username, b.password)
// 	bhi, _ := strconv.Atoi(bh)
// 	block := rpc.GetBlockByHeight(bhi)

// 	bl := map[string]interface{}{
// 		"d": block,
// 	}
// 	fmt.Println("IP RPC source:", block)
// 	out, err := json.Marshal(bl)
// 	if err != nil {
// 		fmt.Println("Error encoding JSON")
// 		return
// 	}
// 	w.Write([]byte(out))
// }

// func  viewBlock() {
// 	v := mux.Vars(r)

// 	lb := map[string]interface{}{
// 		"block": e.getAPIdata(v["coin"], "block", v["block"]),
// 	}
// 	out, err := json.Marshal(lb)
// 	if err != nil {
// 		fmt.Println("Error encoding JSON")
// 		return
// 	}
// 	w.Write([]byte(out))
// }

func (ja *JormAPI) ViewTypes() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ja.Coin = c.Params("coin")
		ja.JORMcommands.Variable = c.Params("vars")
		ja.JORMcommands.Command = c.Params("type")
		per, _ := strconv.Atoi(c.Params("per"))
		page, _ := strconv.Atoi(c.Params("page"))
		vars := map[string]interface{}{
			"coin": ja.Coin,
			"per":  per,
			"page": page,
			"type": c.Params("type"),
		}
		return c.JSON(ja.getAPIdata(vars))
	}
}
func (ja *JormAPI) ViewType() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ja.Coin = c.Params("coin")
		ja.JORMcommands.Command = c.Params("type")
		vars := map[string]interface{}{
			"coin": ja.Coin,
			"id":   c.Params("id"),
			"type": c.Params("type"),
		}
		return c.JSON(map[string]interface{}{
			c.Params("type"): ja.getAPIdata(vars),
		})
	}
}

// func (b *bitNodesRPC) directViewRawMemPool() {
// 	v := mux.Vars(r)
// 	rpc := RPCSRC(bitNodes(b.path, v["coin"]), b.username, b.password)
// 	rawMemPool := rpc.APIGetRawMemPool()
// 	rmp := map[string]interface{}{
// 		"d": rawMemPool,
// 	}
// 	out, err := json.Marshal(rmp)
// 	if err != nil {
// 		fmt.Println("Error encoding JSON")
// 		return
// 	}
// 	w.Write([]byte(out))
// }
// func (b *bitNodesRPC) directViewMiningInfo() {
// 	v := mux.Vars(r)
// 	rpc := RPCSRC(bitNodes(b.path, v["coin"]), b.username, b.password)
// 	miningInfo := rpc.APIGetMiningInfo()

// 	mi := map[string]interface{}{
// 		"d": miningInfo,
// 	}
// 	out, err := json.Marshal(mi)
// 	if err != nil {
// 		fmt.Println("Error encoding JSON")
// 		return
// 	}
// 	w.Write([]byte(out))
// }
// func (b *bitNodesRPC) directViewInfo() {
// 	v := mux.Vars(r)
// 	rpc := RPCSRC(bitNodes(b.path, v["coin"]), b.username, b.password)
// 	info := rpc.APIGetInfo()

// 	in := map[string]interface{}{
// 		"d": info,
// 	}
// 	out, err := json.Marshal(in)
// 	if err != nil {
// 		fmt.Println("Error encoding JSON")
// 		return
// 	}
// 	w.Write([]byte(out))
// }
// func (b *bitNodesRPC) directViewPeers() {
// 	v := mux.Vars(r)
// 	rpc := RPCSRC(bitNodes(b.path, v["coin"]), b.username, b.password)
// 	info := rpc.APIGetPeerInfo()
// 	pi := map[string]interface{}{
// 		"d": info,
// 	}
// 	out, err := json.Marshal(pi)
// 	if err != nil {
// 		fmt.Println("Error encoding JSON")
// 		return
// 	}
// 	w.Write([]byte(out))
// }
