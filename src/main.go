package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"Tyrant/src/api"
	"Tyrant/src/service/impl"
)

var (
	dbPath string
	bind   string
)

func init() {
	flag.StringVar(&dbPath, "db", "", "db path")
	flag.StringVar(&bind, "bind", "", "bind addr")
	flag.Parse()
}

func main() {
	svc, err := impl.New(dbPath)
	if err != nil {
		log.Fatalf("failed to init service: %v", err)
	}
	server, err := api.New(svc, bind)
	if err != nil {
		log.Fatalf("failed to init server: %v", err)
	}
	if err := server.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	<-ch
	server.Stop()
}
