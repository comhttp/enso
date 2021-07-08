package main

import (
	"fmt"
	"github.com/comhttp/jorm-server/app"
	"github.com/comhttp/jorm-server/app/cfg"
	"log"
)

func main() {
	j := app.NewJORM()

	fmt.Println("Listening on port: ", cfg.C.Port)
	log.Fatal(j.WWW.ListenAndServe())
}
