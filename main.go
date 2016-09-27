package main

import (
	"fmt"
	flag "github.com/ogier/pflag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const USAGE = `
❯ Minimal directory hosting

flags:
  --verbose  not terse     (default: false)
  --port     binding port  (defalt: 8080)
  --path     url path      (default: /)
`

func main() {
	verbose := flag.Bool("verbose", false, "not terse (default false)")
	port := flag.String("port", "8080", "binding port")
	path := flag.String("path", "/", "URL path")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, USAGE)
	}
	flag.Parse()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir(dir))
	http.Handle(*path, fs)
	if *verbose {
		log.Printf("Serving %s on port :%s at %s", dir, *port, *path)
	}
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
