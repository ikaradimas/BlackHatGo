package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/briandowns/spinner"
	"github.com/miekg/dns"
)

var (
	domain      string
	wordlist    string
	workers     int
	hostAndPort string
	hlp         bool
)

type result struct {
	IPAddress string
	Hostname  string
}

func init() {
	flag.StringVar(&domain, "d", "", "The domain to guess subdomains for")
	flag.StringVar(&wordlist, "w", "", "The wordlist to use for guessing")
	flag.IntVar(&workers, "c", 100, "The amount of workers to use")
	flag.StringVar(&hostAndPort, "h", "8.8.8.8:53", "The DNS server and port to use")
	flag.BoolVar(&hlp, "help", false, "Get help for this utility")
	flag.Parse()
}

func help() {
	fmt.Println("guessdns - A utility to guess DNS subdomains")
	flag.PrintDefaults()
}

func main() {
	fmt.Println("Starting guessdns")
	if domain == "" || wordlist == "" {
		help()
		fmt.Println("Missing required parameters")
		os.Exit(1)
	}

	fmt.Printf("Starting configuration:\n * Domain: %s\n * Wordlist: %s\n * Workers: %d\n * Host and Port: %s\n",
		domain, wordlist, workers, hostAndPort)

	var results []result
	fqdns := make(chan string, workers)
	gather := make(chan []result)
	tracker := make(chan empty)

	fh, err := os.Open(wordlist)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()

	for i := 0; i < workers; i++ {
		go worker(tracker, fqdns, gather, hostAndPort)
	}

	for scanner.Scan() {
		fqdns <- fmt.Sprintf("%s.%s", scanner.Text(), domain)
	}

	go func() {
		for resultSet := range gather {
			results = append(results, resultSet...)
		}
		var e empty
		tracker <- e
	}()

	close(fqdns)

	for i := 0; i < workers; i++ {
		<-tracker
	}
	close(gather)
	<-tracker

	s.Stop()

	if len(results) < 1 {
		fmt.Println("No subdomains found.")
		return
	}

	fmt.Println("Detected subdomains:")
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, ' ', 0)
	for _, r := range results {
		fmt.Fprintf(w, "%s\t%s\n", r.Hostname, r.IPAddress)
	}
	w.Flush()
}

func lookupA(fqdn, serverAddr string) ([]string, error) {
	var m dns.Msg
	var ips []string
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeA)
	in, err := dns.Exchange(&m, serverAddr)
	if err != nil {
		return ips, err
	}
	if len(in.Answer) < 1 {
		return ips, errors.New("no answer")
	}
	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			ips = append(ips, a.A.String())
		}
	}
	return ips, nil
}

func lookupCNAME(fqdn, serverAddr string) ([]string, error) {
	var m dns.Msg
	var ips []string
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeCNAME)
	in, err := dns.Exchange(&m, serverAddr)
	if err != nil {
		return ips, err
	}
	if len(in.Answer) < 1 {
		return ips, errors.New("no answer")
	}
	for _, answer := range in.Answer {
		if c, ok := answer.(*dns.CNAME); ok {
			ips = append(ips, c.Target)
		}
	}
	return ips, nil
}

func lookup(fqdn, serverAddr string) []result {
	var results []result
	var cfqdn = fqdn

	for {
		cnames, err := lookupCNAME(cfqdn, serverAddr)
		if err == nil && len(cnames) > 0 {
			cfqdn = cnames[0]
			continue
		}
		ips, err := lookupA(cfqdn, serverAddr)
		if err != nil {
			break
		}
		for _, ip := range ips {
			results = append(results, result{IPAddress: ip, Hostname: fqdn})
		}
		break
	}
	return results
}

type empty struct{}

func worker(tracker chan empty, fqdns chan string, gather chan []result, serverAddr string) {
	for fqdn := range fqdns {
		results := lookup(fqdn, serverAddr)
		if len(results) > 0 {
			gather <- results
		}
	}
	var e empty
	tracker <- e
}
