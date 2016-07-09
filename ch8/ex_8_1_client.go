// Run with list of ports to connect to on localhost

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Provide at least one port to listen on")
		os.Exit(1)
	}
	for i := 0; i < len(os.Args[1:]); i++ {

		conn, err := net.Dial("tcp", "localhost:"+os.Args[i+1])
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(os.Stdout, conn, i)
	}

	// Dummy read to prevent program from closing
	var end string
	fmt.Scanf("%s", &end)
}

func mustCopy(dst io.Writer, src io.Reader, idx int) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
