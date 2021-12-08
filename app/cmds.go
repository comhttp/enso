package app

import (
	"github.com/comhttp/enso/app/handlers"
)

func (e *ENSO) getAPIs() {
	e.APIs = make(map[string][]handlers.API)

	apis := e.okno.GetAll("apis", "")
	for _, api := range apis {
		coinsRaw := api["coin"].([]interface{})
		if len(coinsRaw) > 0 {
			for _, coinRaw := range coinsRaw {
				coin := coinRaw.(map[string]interface{})["slug"]
				apiCommands := make(map[string]string)
				for _, command := range api["commands"].([]interface{}) {
					cmd := command.(map[string]interface{})
					apiCommands[cmd["type"].(string)] = cmd["command"].(string)
				}
				e.APIs[coin.(string)] = append(e.APIs[coin.(string)],
					handlers.API{
						UrlFormat: api["urlformat"].(string),
						Commands:  apiCommands,
						Endpoint:  api["url"].(string),
					})
			}
		}
	}
	return
}

func (e *ENSO) getJORMcommands() *handlers.JormCmds {
	e.jormCommands = &handlers.JormCmds{}
	e.jormCommands.CMDs = make(map[string](func(vars map[string]interface{}) map[string]interface{}))
	e.jormCommands.CMDs["status"] = func(vars map[string]interface{}) map[string]interface{} {
		status := e.ExJDBs[vars["coin"].(string)].GetExplorer(vars["coin"].(string))
		d := make(map[string]interface{})
		d["blocks"] = status.Blocks
		d["addresses"] = status.Addresses
		d["txs"] = status.Txs
		return d
	}
	e.jormCommands.CMDs["lastblock"] = func(vars map[string]interface{}) map[string]interface{} {
		lastblock := e.ExJDBs[vars["coin"].(string)].GetLastBlock(vars["coin"].(string))
		d := make(map[string]interface{})
		d["lastblock"] = lastblock
		return d
	}
	e.jormCommands.CMDs["blocks"] = func(vars map[string]interface{}) map[string]interface{} {
		blocks := e.ExJDBs[vars["coin"].(string)].GetBlocks(vars["coin"].(string), vars["per"].(int), vars["page"].(int))
		d := make(map[string]interface{})
		d["blocks"] = blocks
		return d
	}

	e.jormCommands.CMDs["block"] = func(vars map[string]interface{}) map[string]interface{} {
		block := e.ExJDBs[vars["coin"].(string)].GetBlock(vars["coin"].(string), vars["id"].(string))
		return block
	}
	e.jormCommands.CMDs["tx"] = func(vars map[string]interface{}) map[string]interface{} {
		tx := e.ExJDBs[vars["coin"].(string)].GetTx(vars["coin"].(string), vars["id"].(string))
		return tx
	}
	e.jormCommands.CMDs["addr"] = func(vars map[string]interface{}) map[string]interface{} {
		addr := e.ExJDBs[vars["coin"].(string)].GetAddr(vars["coin"].(string), vars["id"].(string))
		return addr
	}
	// e.jormCommands["blocks"] = func(vars map[string]interface{}) map[string]interface{} {
	// 	lastblock := e.ExJDBs[vars["coin"]].ViewBlocks(vars["coin"])
	// 	d := make(map[string]interface{})
	// 	d["lastblock"] = lastblock
	// 	return d
	// }
	return e.jormCommands
}
