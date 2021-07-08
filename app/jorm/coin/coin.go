package coin

type Coins struct {
	N int      `json:"n"`
	C []string `json:"c"`
}

type NodeCoins struct {
	N int        `json:"n"`
	C []NodeCoin `json:"c"`
}
type AlgoCoins struct {
	N int        `json:"n"`
	A []string   `json:"a"`
	C []AlgoCoin `json:"c"`
}
type CoinsWords struct {
	N int    `json:"n"`
	C string `json:"c"`
}
type BaseCoins struct {
	N int        `json:"n"`
	C []BaseCoin `json:"c"`
}

// Coin stores identifying information about coins in the database
type BaseCoin struct {
	Rank   int    `json:"r"`
	Name   string `json:"n"`
	Ticker string `json:"t"`
	Slug   string `json:"s"`
}

// Coin stores identifying information about coins in the database
type NodeCoin struct {
	Rank   int    `json:"r"`
	Name   string `json:"n"`
	Ticker string `json:"t"`
	Slug   string `json:"s"`
	Algo   string `json:"a"`
}

// Coin stores identifying information about coins in the database
type AlgoCoin struct {
	Rank   int    `json:"r"`
	Name   string `json:"n"`
	Ticker string `json:"t"`
	Slug   string `json:"s"`
	Algo   string `json:"a"`
}

type Coin struct {
	Name   string `json:"name" form:"name"`
	Ticker string `json:"ticker" form:"ticker"`
	Slug   string `json:"slug" form:"slug"`

	Rank        int    `json:"rank" form:"rank"`
	Token       string `json:"token" form:"token"`
	Platform    string `json:"platform" form:"platform"`
	Algo        string `json:"algo" form:"algo"`
	Proof       string `json:"proof" form:"proof"`
	Description string `json:"description" form:"description"`

	Ico                  bool    `json:"ico" form:"ico"`
	TotalCoinSupply      float64 `json:"total"`
	BuiltOn              string  `json:"builton"`
	BlockTime            float64 `json:"blocktime"`
	DifficultyAdjustment string  `json:"diff"`
	BlockRewardReduction string  `json:"rew"`
	Start                string  `json:"start"`

	WebSite     []string `json:"web"`
	Explorer    []string `json:"explorer"`
	Chat        []string `json:"chat"`
	Twitter     string   `json:"tw"`
	Facebook    string   `json:"facebook"`
	Telegram    string   `json:"telegram"`
	Reddit      string   `json:"reddit"`
	Github      []string `json:"github"`
	BitcoinTalk string   `json:"bitcointalk"`
	WhitePaper  string   `json:"whitepaper"`
	Logo        bool     `json:"logo" form:"logo"`
	//Logo        utl.Images `json:"logo" form:"logo"`

	Published bool `json:"published" form:"published"`
	BitNode   bool `json:"bitnode" form:"bitnode"`
	Selected  bool `json:"selected" form:"selected"`
	Favorite  bool `json:"fav" form:"favorite"`

	Checked map[string]bool `json:"checked"`

	Markets map[string]string `json:"markets"`
}

func (coin *Coin) SelectCoin() *Coin {
	//coin.LogoBig = LoadLogo(coin.Slug, "img128")
	//coin.Data = LoadInfo(coin.Slug)
	return coin
}

//func LoadLogo(slug, size string) image.Image {
//	// Load logo image from database
//	logos := make(map[string]interface{})
//	fmt.Println("slug", slug)
//	err := jdb.JDB.Read(filepath.FromSlash(cfg.C.Out+"/data/"+slug), "logo", logos)
//	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(logos[size].(string)))
//	logo, _, err := image.Decode(reader)
//	utl.ErrorLog(err)
//	return logo
//}
//
//func LoadInfo(slug string) Coin {
//	// Load coin data from database
//	info := Coin{}
//	err := jdb.JDB.Read(filepath.FromSlash("data/"+slug), "info", info)
//	utl.ErrorLog(err)
//	//jsonString, _ := json.Marshal(info)
//
//	// convert json to struct
//	//s := CoinData{}
//	//json.Unmarshal(jsonString, &s)
//	return info
//}

//func LoadCoinsBase() Coins {
	//usableCoins := Coins{N: 0}
	//nodeCoins := NodeCoins{N: 0}
	//algoCoins := AlgoCoins{N: 0}
	//coinsWords := CoinsWords{N: 0}
	//restCoins := Coins{N: 0}
	//
	//coinsBin := Coins{N: 0}
	//allCoins := Coins{N: 0}
	//c := getCoins(j)
	//fmt.Println("cccccccccc::",c)

	//var bitNodes []string
	//if err := jdb.JDB.Read(filepath.FromSlash(cfg.C.Out+"/info"), "bitnoded", &bitNodes); err != nil {
	//	fmt.Println("Error", err)
	//}
	//for _, coiNn := range c {
	//	for _, bitNode := range bitNodes {
	//		if bitNode == coiNn.Slug {
	//			coiNn.BitNode = true
	//			jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/coins"), coiNn.Slug, coiNn)
	//		}
	//	}
	//}

	//for i, coin := range c {
	//if coin.BitNode {
	//	nodeCoins.N++
	//	nodeCoins.C = append(nodeCoins.C, NodeCoin{
	//		Rank:   coin.Rank,
	//		Name:   coin.Name,
	//		Ticker: coin.Ticker,
	//		Slug:   coin.Slug,
	//		Algo:   coin.Algo,
	//	})
	//} else {
	//	if coin.Algo != "N/A" && coin.Algo != "" {
	//		algoCoins.N++
	//		algoCoins.C = append(algoCoins.C, AlgoCoin{
	//			Rank:   coin.Rank,
	//			Name:   coin.Name,
	//			Ticker: coin.Ticker,
	//			Slug:   coin.Slug,
	//			Algo:   coin.Algo,
	//		})
	//		for _, a := range algoCoins.A {
	//			if a != coin.Algo {
	//				algoCoins.A = append(algoCoins.A, coin.Algo)
	//				fmt.Println("111aaa", a)
	//				fmt.Println("111coin.Algo", coin.Algo)
	//			}
	//			fmt.Println("222aaa", a)
	//			fmt.Println("2222coin.Algo", coin.Algo)
	//		}
	//	} else {
	//		if coin.Description != "" {
	//			//len(c[i].WebSite) > 0 &&
	//			// len(coin.WebSite) > 0 &&
	//			//if c[i].Platform != "token" &&
	//			restCoins.N++
	//			restCoins.C = append(restCoins.C, coin.Slug)
	//		} else {
	//			coinsBin.N++
	//			coinsBin.C = append(coinsBin.C, coin.Slug)
	//		}
	//	}
	//	usableCoins.N = i
	//	usableCoins.C = append(usableCoins.C, coin.Slug)
	//	coinsWords.C = coinsWords.C + " " + coin.Name
	//	coinsWords.N = usableCoins.N
	//}
	//allCoins.N = i
	//allCoins.C = append(allCoins.C, coin.Slug)
	//}
	//jdb.JDB.Write(cfg.C.Out+"/info", "restcoins", restCoins)
	//jdb.JDB.Write(cfg.C.Out+"/info", "algos", algoCoins)
	//jdb.JDB.Write(cfg.C.Out+"/info", "coinswords", coinsWords)
	//jdb.JDB.Write(cfg.C.Out+"/info", "usableinfo", usableCoins)
	//jdb.JDB.Write(cfg.C.Out+"/info", "allcoins", allCoins)
	//
	////jdb.JDB.Write("jorm/info", "bitnodes", LoadCoinsBase(true, true))
	//jdb.JDB.Write(cfg.C.Out+"/info", "nodecoins", nodeCoins)
	//jdb.JDB.Write(cfg.C.Out+"/info", "coinsbin", coinsBin)
	//return c
//}

//func getCoins(j *jdb.JDB) Coins {
//	coins, err := j.ReadAll("coins", "coins_")
//	utl.ErrorLog(err)
//	allCoins := Coins{N: 0}
//	for i, c := range coins {
//		allCoins.C = append(allCoins.C, c)
//		allCoins.N = i
//	}
//	//data, err := jdb.JDB.ReadAll(cfg.C.Out + "/coins")
//
//	//coins := make([][]byte, len(data))
//	//for i := range data {
//	//	coins[i] = []byte(data[i])
//	//}
//	//cs := make([]Coin, len(coins))
//	//for i := range coins {
//	//	if err := json.Unmarshal(coins[i], &cs[i]); err != nil {
//	//		fmt.Println("Error", err)
//	//	}
//	//}
//	//return cs
//	return allCoins
//}
