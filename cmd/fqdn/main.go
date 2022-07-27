package main

import (
	"flag"
	"log"

	"github.com/miekg/dns"
)

var (
	search      string
	hostAndPort string
)

func init() {
	flag.StringVar(&search, "s", "example.com", "Search string")
	flag.StringVar(&hostAndPort, "h", "8.8.8.8:53", "DNS host")
	flag.Parse()
}

func main() {
	log.Println("fqdn - FQDN lookup utility")
	log.Printf("Using dns: %s", hostAndPort)
	log.Printf("Looking up type A records for: %s", search)

	var msg dns.Msg
	fqdn := dns.Fqdn(search)
	msg.SetQuestion(fqdn, dns.TypeA)

	in, err := dns.Exchange(&msg, hostAndPort)
	if err != nil {
		log.Fatalln(err)
	}

	if len(in.Answer) < 1 {
		log.Println("No records found.")
		return
	}

	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			log.Println(a.A)
		}
	}
}
