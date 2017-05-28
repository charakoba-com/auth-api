package main

import (
	"log"
	"os"

	authapi "github.com/charakoba-com/auth-api"
	flags "github.com/jessevdk/go-flags"
)

type options struct {
	Listen string `short:"l" long:"listen" default:":8080" description:"Listen address"`
}

func main() {
	os.Exit(_main())
}

func _main() int {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		log.Printf("%s", err)
		return 1
	}
	if err := authapi.Run(opts.Listen); err != nil {
		log.Printf("%s", err)
		return 1
	}
	return 0
}
