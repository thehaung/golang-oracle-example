package main

import (
	"github.com/thehaung/golang-oracle-example/bootstrap"
	"github.com/thehaung/golang-oracle-example/config"
	"log"
)

func main() {
	conf, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	bootstrap.Run(conf)
}
