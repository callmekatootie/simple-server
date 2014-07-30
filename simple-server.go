package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
)

var (
	port *int = flag.Int("p", 8080, "port to listen for connections")
	// address may be empty to listen on all interfaces
	address   *string = flag.String("a", "", "address to bind to")
	directory string
)

func main() {
	flag.Parse()

	var err error
	directory := flag.Arg(0)

	if len(directory) == 0 {
		// No directory specified. Use current directory to host files
		// use os.Getwd() for platform-independent resolving of current working directory
		directory, err = os.Getwd()
		if err != nil {
			fmt.Printf("error retrieving current directory: %v", err)
			os.Exit(1)
		}
	}

	http.Handle("/", http.FileServer(http.Dir(directory)))

	completeAddress := net.JoinHostPort(*address, fmt.Sprint(*port))

	fmt.Printf("Starting http server at %s\n", completeAddress)

	fmt.Println("Hit CTRL-C to stop the server")

	err = http.ListenAndServe(completeAddress, nil)

	if err != nil {
		fmt.Println("\nERROR")
		fmt.Println(err)
	}
}
