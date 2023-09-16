package main

import (
	"flag"
	"log"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"yubisuma-2.0/app"
)

var (
	// port番号
	addr string
)

func init() {

	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	// dbの初期化
	app.InitDB()
	// echoの処理
	e := echo.New()
	// panicが発生した時の処理
	e.Use(echomiddleware.Recover())
	// CORS周りの設定
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		Skipper:      echomiddleware.DefaultCORSConfig.Skipper,
		AllowOrigins: echomiddleware.DefaultCORSConfig.AllowOrigins,
		AllowMethods: echomiddleware.DefaultCORSConfig.AllowMethods,
		AllowHeaders: echomiddleware.DefaultCORSConfig.AllowHeaders,
	}))

	e.GET("/animals", app.HandleAnimalsGet())
	e.POST("/animal", app.HandleAnimalPost())
	e.POST("/action", app.HandleAction())

	log.Println("Server running...")
	if err := e.Start(addr); err != nil {
		log.Fatalf("failed to start server. %+v", err)
	}
}
