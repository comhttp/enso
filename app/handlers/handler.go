package handlers

import (
	"encoding/hex"
	"log"
	"strconv"

	"github.com/comhttp/enso/pkg/utl"
	"github.com/comhttp/jorm/mod/coin"
	"github.com/comhttp/jorm/pkg/utl/img"
	"github.com/gofiber/fiber/v2"
)

// CoinsHandler handles a request for coin data
func CoinHandler(cq *coin.CoinsQueries) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		coin, err := cq.GetCoin(c.Params("coin"))
		utl.ErrorLog(err)
		return c.JSON(coin)
	}
}

// CoinsHandler handles a request for coin data
func CoinsHandler(cq *coin.CoinsQueries) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(cq.GetCoins())
	}
}

// CoinsHandler handles a request for coin data
func RestCoinsHandler(cq *coin.CoinsQueries) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(cq.GetRestCoins())
	}
}

// CoinsHandler handles a request for coin data
func CoinsWordsHandler(cq *coin.CoinsQueries) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(cq.GetCoinsWords())
	}
}

// CoinsHandler handles a request for coin data
func UsableCoinsHandler(cq *coin.CoinsQueries) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(cq.GetUsableCoins())
	}
}

// CoinsHandler handles a request for coin data
func AllCoinsHandler(cq *coin.CoinsQueries) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(cq.GetAllCoins())
	}
}

// CoinsHandler handles a request for coin data
func NodeCoinsHandler(cq *coin.CoinsQueries) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(cq.GetNodeCoins())
	}
}

// CoinsHandler handles a request for coin data
func CoinsBinHandler(cq *coin.CoinsQueries) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(cq.GetCoinsBin())
	}
}

// CoinNodesHandler handles a request for (?)
func CoinNodesHandler(cq *coin.CoinsQueries) {

}

// CoinsHandler handles a request for coin data
func AlgoCoinsHandler(cq *coin.CoinsQueries) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(cq.GetAlgoCoins())
	}
}

// NodeHandler handles a request for (?)
func NodeHandler(cq *coin.CoinsQueries) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// return c.JSON(nodes.GetNode(cq., c.Params("coin"), c.Params("nodeip")))
		return nil
	}
}

// LogoHandler handles a request for logo data
func LogoHandler(cq *coin.CoinsQueries) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		size, err := strconv.ParseFloat(c.Params("size"), 32)
		log.Print("Error encoding JSON: ", err)
		c.Type("png")
		_, err = c.Write([]byte(getLogo(cq, c.Params("coin"), size)))
		return err
	}
}

// LogoHandler handles a request for logo data
func getLogo(cq *coin.CoinsQueries, coin string, size float64) []byte {
	logoRawString, err := cq.GetLogo(coin)
	if err != nil {
		log.Print("Error encoding JSON")
	}
	logoRaw, err := hex.DecodeString(logoRawString)
	logo, _ := img.ImageResize(logoRaw, img.Options{Width: size, Height: size})
	return logo
}

// jsonHandler handles a request for json data
func JsonAlgoCoinsHandler(cq *coin.CoinsQueries) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(cq.GetAlgoCoinsLogo())
	}
}
