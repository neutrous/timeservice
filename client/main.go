// -*- mode: go -*-

package main

import (
	"fmt"
	"os"
	
	"github.com/neutrous/TimeService/client"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: <proto>://<address>:<port>")
		return
	}

	address := os.Args[1]
	
	client := tu.NewClient(address)
	if client == nil {
		fmt.Println("Create client instance failure.")
		return
	}

	// Start receive the time datas.
	client.Subscribe()
}

