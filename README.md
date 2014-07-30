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
If you have set up your `$GOPATH` correctly and you have included `$GOPATH/bin` directory in your command's PATH
variable, you should be able to use this as:  

    $ simple-server [path] [options]

[path] defaults to the current working directory (./)

You can find a list of supported flags / options by executing the command as:  

    $ simple-server -h

The allowed flags / options are:

-p    Specify a port for listening to connections (default is 8080)  
-a    Specify an address to bind the server to (default is 127.0.0.1)  

Example:  

    $ simple-server -a 127.0.0.1 -p 8082

will set up a server that can be accessed at [http://localhost:8082](http://localhost:8082)
