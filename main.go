package main

import (
	"fmt"
	"github.com/comhttp/enso/app"
	"github.com/comhttp/jorm/pkg/cfg"
	"log"
)

func main() {
	j := app.NewENSO()

	fmt.Println("Listening on port: ", cfg.C.Port["enso"])
	log.Fatal(j.WWW.ListenAndServe())
}
