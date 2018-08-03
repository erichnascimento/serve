package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/erichnascimento/rocket/middleware"
	"github.com/erichnascimento/rocket/server"
)

const Version = "0.0.1"

var (
	port int
	dir  string
	ver  bool
)

func main() {
	flag.IntVar(&port, "port", 8080, "specify the port. Default 8080")
	flag.StringVar(&dir, "dir", ".", `specify the root directory. Default "."`)
	flag.BoolVar(&ver, "version", false, `print the version`)
	flag.Parse()

	if ver {
		fmt.Println(Version)
		os.Exit(0)
	}

	address := fmt.Sprintf(`:%d`, port)
	fmt.Printf("Listening and serving \"%s\" dir on addresss: %s\n", dir, address)
	fmt.Printf("Access http://localhost:%s\n\n", address)

	s := server.NewServer()
	s.Use(middleware.NewLogger())
	s.Use(newFileHandler(dir))

	log.Fatal(s.ListenAndServe(address))
}

func newFileHandler(dir string) middleware.Middleware {
	h := func(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		http.FileServer(http.Dir(dir)).ServeHTTP(rw, req)
		next(rw, req)
	}

	return middleware.NewMiddleFunc(h)
}
