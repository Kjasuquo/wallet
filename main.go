package main

import (
	"flag"
	"github.com/kjasuquo/preparation/config"
	"os"

	"log"
)

func main() {
	var cfg = config.NewConfig()
	flag.IntVar(&cfg.Port, "port", 8081, "Server port to listen to")

	cfg.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	cfg.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	err := cfg.Server()
	if err != nil {
		cfg.ErrorLog.Print(err)
		log.Fatal(err)
	}
}
