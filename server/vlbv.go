package main

import (
	"flag"
	"strconv"

	"github.com/himu62/vlbv/server/log"
	"github.com/himu62/vlbv/server/web"
)

func main() {
	port := flag.Int("port", 8080, "web server listen port")
	isDebug := flag.Bool("debug", false, "debug flag")
	flag.Parse()

	var logger *log.Logger
	if *isDebug {
		logger = log.NewDebugLogger()
	} else {
		logger = log.NewLogger()
	}

	server := web.NewWebServer(logger)
	server.Listen(":" + strconv.Itoa(*port))
}
