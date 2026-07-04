//sample tcp chat app with client and server
//really basic and really shit...

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func HandleIncoming(conn net.Conn) {
	_, err := conn.Write([]byte("Your connection to the server was estabilished"))
	defer conn.Close()

	for {
		//buffer to store client messages
		buf := make([]byte, 64)

		if err != nil {
			log.Println("Writing failed!")
		}

		n, err := conn.Read(buf)

		if err != nil {
			log.Printf("Connection disconnected!")
		}

		fmt.Printf("Client : %s\n", string(buf[:n]))
	}
}

func HandleOutgoing(conn net.Conn) {
	defer conn.Close()

	for {

		reader := bufio.NewReader(os.Stdin)

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error while reading line\n")
		}

		_, err = conn.Write([]byte(fmt.Sprintf("server: %s\n", line)))

		if err != nil {
			log.Printf("Connection to addr: %s closed.\n", conn.RemoteAddr())
		}
	}

}

func main() {

	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Println("Error listening to port 2000. Error", err.Error())
	}

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Printf("Error: %s while connecting to addr: %s\n",
				err.Error(), conn.RemoteAddr())
		} else {
			log.Printf("Connected successfully to remote addr: %s\n",
				conn.RemoteAddr())
		}

		go HandleIncoming(conn)

		go HandleOutgoing(conn)
	}
}
