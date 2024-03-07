package main

import (
	//"fmt"
	"os"
	"log"
	"net"
)

func main() {
	const name string = "webscrapper"
	log.SetPrefix(name + ":  ")

	if len(os.Args) != 2 {
		log.Fatal("no url specified")
	}

	host := os.Args[1]
	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatalf("lookup is: %s: %v", host, err)
	}
	if len(ips) == 0 {
		log.Fatalf("no ips found for %s", host)
	}
}