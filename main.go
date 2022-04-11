package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/zapponejosh/html_link/internal/Link"
)

func main() {
	var (
		filename string
	)
	flag.StringVar(&filename, "file", "ex1.html", "Filename/path for HTML file to read.")
	flag.Parse()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file %s", filename)
	}

	links, err := Link.Parse(file)
	fmt.Println(links)

}
