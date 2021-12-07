package main

import (
	"flag"

	"github.com/rs/zerolog/log"

	"github.com/comhttp/enso/app"
	"github.com/rs/zerolog"
)

func main() {
	// Get cmd line parameters
	path := flag.String("path", "/var/db/jorm", "Path")
	loglevel := flag.String("loglevel", "debug", "Logging level (debug, info, warn, error)")
	flag.Parse()

	//j.Log.SetLevel(parseLogLevel(*loglevel))
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Default level for this example is info, unless debug flag is present

	switch *loglevel {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

	e := app.NewENSO(*path)

	// fmt.Println("Listening on port: ", cfg.C.Port["enso"])
	e.ENSOrouter()

	log.Fatal().Err(e.Router.Listen(":14433"))

}
