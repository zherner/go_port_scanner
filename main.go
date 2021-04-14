package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

// scan attempts to connect to target on target ports
func scan(t, tp string, to time.Duration, wg *sync.WaitGroup) {
	// decrease wait group count
	defer wg.Done()

	// attempt to connect to target:port with a timeout
	conn, err := net.DialTimeout("tcp", t+":"+tp, to)
	// closed port
	if err != nil {
		fmt.Printf("%s port %s closed!\n", t, tp)
		return
	}
	conn.Close()
	fmt.Printf("%s port %s opened!\n", t, tp)

	// webserver type from header on web ports
	webserver(t, tp)
}

// webserver prints the "Server" header if there is one
func webserver(t, p string) {
	if p == "80" {
		// make http GET request to target
		res, err := http.Get("http://" + t)
		if err != nil {
			log.Fatal(err)
		}
		// print "Server header value"
		fmt.Println("└─" + string(res.Header.Get("Server")))
	}
}

func main() {
	var (
		target_ip    = []string{"151.101.194.152", "172.217.10.238"}
		target_ports = []string{"80", "443", "22", "8080"}
		target_timeo = time.Second
	)

	// wait group for main proc to wait on goroutines
	var wg sync.WaitGroup

	// loop over targets
	for _, v := range target_ip {
		// inner loop over ports
		for _, k := range target_ports {
			// add a count to wait group, wg.Done() decreases count
			wg.Add(1)
			// scan in goroutine
			go scan(v, k, target_timeo, &wg)
		}
	}
	// wait for wait group to be zero
	wg.Wait()
}
