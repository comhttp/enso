package app

import (
	"time"

	"github.com/comhttp/enso/app/routes"
	"github.com/comhttp/enso/pkg/utl"
	"github.com/comhttp/jorm/mod/coin"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/utils"
)

func (e *ENSO) ENSOrouter() {
	e.Router = fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "OKNO",
		AppName:       "ENSO",
	})

	e.Router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	e.Router.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))
	e.Router.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	e.Router.Use(csrf.New(csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
	}))
	e.Router.Use(etag.New())

	e.Router.Use(favicon.New(favicon.Config{
		File: "",
	}))
	// bitnodes := e.okno.GetAll("nodes", "bitnode=true&")

	// for _, bitnode := range bitnodes {
	// 	if bitnode["coin"].(map[string]interface{})["slug"] == coin {
	// 		e.BitNodes = append(e.BitNodes, nodes.BitNode{
	// 			IP:   bitnode["ip"].(string),
	// 			Port: int64(bitnode["port"].(float64)),
	// 		})
	// 	}
	// }

	coins, err := jdb.NewJDB(e.jdbServers["jdbcoins"])
	cq := coin.Queries(coins, "")
	if err != nil {
		utl.ErrorLog(err)
	} else {
		cq.WriteInfo("nodecoins", &coin.Coins{
			N: len(e.BitNoded),
			C: e.BitNoded,
		})
	}
	routes.CoinsRoutes(cq, e.okno, e.APIs, e.getJORMcommands(), e.ExJDBs, e.Router)

	// routes.Explorer(e.okno, e.APIs, e.getJORMcommands(), e.ExJDBs, e.Router)
	//s := r.Host("enso.okno.rs").Subrouter()
	// r.StrictSlash(true)
	// return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))

	// log.Fatal().Err(e.Router.Listen(":3000"))
	// log.Fatal(e.Router.Listen(":3000"))

}
