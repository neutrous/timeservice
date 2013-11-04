// -*- mode: go -*-

package main

import (
	"fmt"
	"os"
	
	"github.com/neutrous/TimeService/server"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: <proto>://<address>:<port>\n")
		return
	}
	
	var addresses []string
	addresses = append(addresses, os.Args[1])
	
	server := tu.NewServer(addresses)
	if server == nil {
		fmt.Printf("Create server instance failure.\n")
		return
	}

	// Start pub the time datas.
	server.Publish()
}



















