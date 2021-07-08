package csrc

import (
	"encoding/json"
	"fmt"
	"github.com/comhttp/jorm-server/app/jdb"
	"github.com/comhttp/jorm-server/app/jorm/coin"
	"github.com/comhttp/jorm-server/pkg/utl"
	"io/ioutil"
	"net/http"
	"time"
)

func getCoinGecko(j *jdb.JDB) {
	fmt.Println("GetCoinGeckoStart")
	var coinsRaw []map[string]interface{}
	respcs, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
	utl.ErrorLog(err)
	defer respcs.Body.Close()
	mapBody, err := ioutil.ReadAll(respcs.Body)
	if mapBody != nil {
		json.Unmarshal(mapBody, &coinsRaw)
		for _, coinSrc := range coinsRaw {
			if coinSrc["id"] != nil {
				if coinSrc["id"] != nil && coinSrc["id"].(string) != "" {
					slug := utl.MakeSlug(coinSrc["id"].(string))
					coin.GetCoin(j, "cg", slug, getCoinGeckoCoin(slug, coinSrc))
				}
			}
		}
	}
	fmt.Println("GetCoinGeckoDone")
}

func getCoinGeckoCoin(slug string, coinSrc map[string]interface{}) func(c *coin.Coin) {
	return func(c *coin.Coin) {
		c.SetName(coinSrc["name"])
		c.SetTicker(coinSrc["symbol"])
		fmt.Println("Checked1:", c.Checked)
		c.Slug = slug
		coinDetails := make(map[string]interface{})
		respc, err := http.Get("https://api.coingecko.com/api/v3/coins/" + coinSrc["id"].(string) + "?tickers=false&market_data=false&community_data=true&developer_data=false&sparkline=false")
		utl.ErrorLog(err)
		defer respc.Body.Close()
		mapBody, err := ioutil.ReadAll(respc.Body)
		if mapBody != nil {
			json.Unmarshal(mapBody, &coinDetails)
			if coinDetails["description"] != nil {
				c.SetDescription(coinDetails["description"].(map[string]interface{})["en"])
			}
			c.SetAlgo(coinDetails["hashing_algorithm"])
			c.SetStart(coinDetails["genesis_date"])

			if coinDetails["image"] != nil {
				c.SetLogo(coinDetails["image"].(map[string]interface{})["large"])
			}
			if coinDetails["links"] != nil {
				c.SetWebSite(coinDetails["links"].(map[string]interface{})["homepage"])
				c.SetExplorer(coinDetails["links"].(map[string]interface{})["blockchain_site"].([]interface{}))
				c.SetChat(coinDetails["links"].(map[string]interface{})["chat_url"].([]interface{}))

				if coinDetails["links"].(map[string]interface{})["bitcointalk_thread_identifier"] != nil {
					c.SetBitcoinTalk(fmt.Sprintf("%f", int(coinDetails["links"].(map[string]interface{})["bitcointalk_thread_identifier"].(float64))))
				}
				c.SetTwitter(coinDetails["links"].(map[string]interface{})["twitter_screen_name"])
				c.SetTelegram(coinDetails["links"].(map[string]interface{})["telegram_channel_identifier"])
				c.SetReddit(coinDetails["links"].(map[string]interface{})["subreddit_url"])
			}
			//if coinDetails["block_time_in_minutes"] != nil {
			//	c.BlockTime = insertFloat(coinDetails["block_time_in_minutes"].(float64), c.BlockTime)
			//}
			// insertFloat(coinDetails["block_time_in_minutes"].(float64), c.BlockTime)})

			//fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
			fmt.Println("Ubo:", c.Name)
			fmt.Println("Checked1:", c.Checked)
			fmt.Println("Checked2:", c.Checked)
		}
		time.Sleep(99 * time.Millisecond)
	}
}
