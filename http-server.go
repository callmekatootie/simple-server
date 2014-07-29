package main

import (
    "fmt"
    "flag"
    "net/http"
)

type serverConfig struct {
    port *int
    address *string
}

func (o *serverConfig) parseFlags() {
    o.port = flag.Int("p", 8080, "port to listen for connections")
    o.address = flag.String("a", "127.0.0.1", "address to bind to")

    flag.Parse()
}

func main() {
    options := serverConfig{}

    options.parseFlags()

    http.Handle("/", http.FileServer(http.Dir("./")))

    fmt.Printf("Starting http server at %s on port %d \n", *options.address, *options.port)

    fmt.Println("Hit CTRL-C to stop the server")

    address := fmt.Sprintf("%s:%d", *options.address, *options.port)

    err := http.ListenAndServe(address, nil)

    if err != nil {
        fmt.Println("\nERROR")
        fmt.Println(err);
    }
}
