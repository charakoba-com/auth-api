package main

import (
	"os"

	authapi "github.com/charakoba-com/auth-api"
	flags "github.com/jessevdk/go-flags"
	logger "github.com/nasa9084/go-logger"
)

type options struct {
	Listen string `short:"l" long:"listen" default:":8080" description:"Listen address"`
}

func main() {
	os.Exit(_main())
}

func _main() int {
	log := logger.New(os.Stdout, "", logger.InfoLevel)
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		log.Errorf.Printf("%s", err)
		return 1
	}
	if err := authapi.Run(opts.Listen); err != nil {
		log.Errorf("%s", err)
		return 1
	}
	return 0
}
