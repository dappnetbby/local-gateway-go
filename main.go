package main

import (
	"log"
	lib "github.com/dappnetbby/local-gateway/lib"
	"os"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
const VERSION = "0.0.1"
const GATEWAY_PORT = 10424

func main() {
    logger.Printf("Dappnet Local Gateway v%s\n", VERSION)
	logger.Println("Copyright Liam Zebedee and contributors")

	gateway := lib.NewGatewayServer(GATEWAY_PORT)
	
	go func(){
		logger.Println("")
		logger.Printf("Listening on http://0.0.0.0:%d\n", GATEWAY_PORT)
		if err := gateway.Start(); err != nil {
            logger.Printf("Error starting server: %v\n", err)
        }
	}()

	ch := make(chan bool)
	<-ch
}