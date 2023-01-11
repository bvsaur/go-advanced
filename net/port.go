package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

var site = flag.String("site", "scanme.nmap.org", "url to scan")

func checkPort(site string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", port))
	if err != nil {
		return
	}
	conn.Close()
	fmt.Println("Port", port, "is open")
}

func main() {
	flag.Parse()
	start := time.Now()
	var wg sync.WaitGroup

	fmt.Println("Site being scanned:", *site)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go checkPort(*site, i, &wg)
	}
	wg.Wait()
	fmt.Printf("Time taken: %s\n", time.Since(start))

	// Without concurrency -> 16s
	// With concurrency -> 182ms
}
