// Run with: go run ex_8_1_server.go <port> <iana_location_name> <sleep_seconds>
//	$ go run ex_8_1_server.go 8001 America/New_York 1 &
//	$ go run ex_8_1_server.go 8002 Europe/London 2 &
//	$ go run ex_8_1_server.go 8003 Asia/Tokyo 3 &
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var ianaLocationName string
var sleepSeconds int

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		location, err := time.LoadLocation(ianaLocationName)
		if err != nil {
			fmt.Println("Illegal IANA location")
			return
		}
		timeInZone := time.Now().In(location)
		_, err = io.WriteString(c, timeInZone.Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Expected IANA location name as command line input")
		os.Exit(1)
	}
	ianaLocationName = os.Args[2]
	sleepSeconds, _ = strconv.Atoi(os.Args[3])

	listener, err := net.Listen("tcp", "localhost:"+os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
