package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	var (
		// filename string
		origin string
	)
	// flag.StringVar(&filename, "file", "ex1.html", "Filename/path for HTML file to read.")
	flag.StringVar(&origin, "origin", "https://gigizapponecounseling.com", "The origin of the site you want to create the sitemap for (https://...)")
	flag.Parse()
	// remove trailing slash if present so I can add to reletive urls
	origin = strings.TrimRight(origin, "/")

	c := http.Client{CheckRedirect: CheckRedirect}

	logFile, err := os.Create("logs/" + time.Now().String() + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(logFile, "RedirectTo: ", log.Lmsgprefix)
	resp, err := c.Get(origin)
	if err != nil {
		// redirect probably
		logger.Println(err)
	}
	if resp.StatusCode != 200 {
		url := "url will go here"
		logger.Printf("Status: %d URL: %s", resp.StatusCode, url)
	}
	loc := resp.StatusCode
	fmt.Println(loc)

	// first request
	// store origin
	// parse body and get links -- doing BFS
	// if links are relative, add origin
	// if links are of different domain, ignore
	// create a map (map1) of link urls so i can do fast lookups and prevent dups
	// for each in map1  goto url and get links
	// if not in map1 add to new map (map2) to check later
	// once map1 is done go through map2 if map2 len > 0
	// create map3 with new links and add map2 links to map1 after completion
	// how does this know to end? if there are no new urls

}

func CheckRedirect(req *http.Request, via []*http.Request) error {
	loc := req.Header.Get("Location")
	return errors.New(loc)
}
