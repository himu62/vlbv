package main

import (
	"flag"
	"strconv"

	"github.com/labstack/echo"
)

var e *echo.Echo

func main() {
	port := flag.Int("port", 8080, "web server listen port")
	isDebug := flag.Bool("debug", false, "debug flag")
	flag.Parse()

	e := echo.New()
	e.Debug = *isDebug

	e.Static("/", "public_html")

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(*port)))
}
