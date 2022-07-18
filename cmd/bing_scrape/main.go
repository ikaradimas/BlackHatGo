package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/ikaradimas/blackhatgo/metadata"
)

func handler(i int, s *goquery.Selection) {
	log.Println(s.Html())
	url, ok := s.Find("a").Attr("href")
	if !ok {
		return
	}

	fmt.Printf("%d: %s\n", i, url)
	res, err := http.Get(url)
	if err != nil {
		return
	}
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return
	}

	cp, ap, err := metadata.NewProperties(r)
	if err != nil {
		return
	}

	log.Printf("%25s %25s - %s %s\n", cp.Creator, cp.LastModifiedBy,
		ap.Application, ap.GetMajorVersion())
}

func main() {
	if (len(os.Args)) != 3 {
		log.Fatalln("Missing required argument. Usage: bing_scrape domain ext")
	}
	domain := os.Args[1]
	filetype := os.Args[2]

	q := fmt.Sprintf("site:%s && filetype:%s && instreamset:(url title):%s",
		domain, filetype, filetype)
	search := fmt.Sprintf("https://www.bing.com/search?q=%s", url.QueryEscape(q))

	client := http.Client{}
	req, _ := http.NewRequest("GET", search, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	resp, err := client.Do(req)

	if err != nil {
		log.Panicln(err)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Panicln(err)
	}

	s := "html body div#b_content ol#b_results li.b_algo h2"
	log.Println(doc.Find(s).Html())
	doc.Find(s).Each(handler)

}
