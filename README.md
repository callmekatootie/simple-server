simple-server
=============

A simple command line http-server built using Go.

Installation
------------

    $ go get github.com/callmekatootie/simple-server

This should automatically download the repository and install the application.
The application executable should now be located in your `$GOPATH/bin` directory

Ensure that you have included the `$GOPATH/bin` directory in your command line's PATH variable to use
it directly in any directory as a command

Usage
-----
You can find a list of supported flags by executing the command as:  

    $ simple-server -h

The allowed flags are:

-p    Specify a port for listening to connections
-a    Spoecify an address to bind the server to

Example:  

    $ simple-server -a 127.0.0.1 -p 8082

will set up a server that can be accessed at [http://localhost:8082](http://localhost:8082)

By default, the address is 127.0.0.1 and the port is 8080

The server will host files in the present working directory
