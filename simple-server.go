package main

import (
    "fmt"
    "flag"
    "net/http"
)

var (
    port *int = flag.Int("p", 8080, "port to listen for connections")
    address *string = flag.String("a", "127.0.0.1", "address to bind to")
    directory string
)

func main() {
    flag.Parse()

    directory := flag.Arg(0)

    if len(directory) == 0 {
        //No directory specified. Use current directory to host files
        directory = "./"
    }

    http.Handle("/", http.FileServer(http.Dir(directory)))

    fmt.Printf("Starting http server at %s on port %d \n", *address, *port)

    fmt.Println("Hit CTRL-C to stop the server")

    completeAddress := fmt.Sprintf("%s:%d", *address, *port)

    err := http.ListenAndServe(completeAddress, nil)

    if err != nil {
        fmt.Println("\nERROR")
        fmt.Println(err);
    }
}
