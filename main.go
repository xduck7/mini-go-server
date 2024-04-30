package main

import (
	"flag"
	"fmt"
	"github.com/xduck7/mini-go-server/docs"
	"github.com/xduck7/mini-go-server/internal/app"
	"os"
)

// @title Mini-Go-Server
// @version 0.7.3
// @description API Server for my github project

func main() {
	var portFlag string
	flag.StringVar(&portFlag, "port", "8080", "use -port= for setting port")
	flag.Parse()

	port := os.Getenv("PORT")
	if port == "" {
		port = portFlag
	}

	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", port)
	app.Run(port)
}
