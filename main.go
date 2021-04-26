package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
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
}

// slicePorts parses inputted ports either single or
// port range and returns a string slice of every port in rage.
// e.g. 80-82 -> []string{"80", "81", "82"}
func slicePorts(ports *string) []string {
	var portsToScan []string

	// split port on dash. If no dash single element slice.
	ps := strings.Split(*ports, "-")

	// build complete string slice of port range.
	if len(ps) >= 1 && len(ps) < 3 {
		// First element in range. Or only element in rage.
		start, err := strconv.Atoi(ps[0])
		if err != nil {
			log.Fatal("Start port not a number.")
		}
		// If only element in range, return.
		if len(ps) == 1 {
			portsToScan = append(portsToScan, strconv.Itoa(start))
			return portsToScan
		}

		// Get end num
		end, err := strconv.Atoi(ps[1])
		if err != nil {
			log.Fatal("End port not a number.")
		}

		// build start to end slice
		for i := start; i <= end; i++ {
			portsToScan = append(portsToScan, strconv.Itoa(i))
		}
		return portsToScan
	} else {
		log.Fatal("Invalid port range.")
	}

	return nil
}

func main() {

	// command line input flags
	targets := flag.String("h", "", "Comman delimited list of IPs to scan.")
	t_ports := flag.String("p", "", "Port number to scan on, or range of ports to scan. e.g. 8080-8088")
	timeout := time.Duration(*flag.Int("t", 1, "Connection timeout in seconds.")) * time.Second

	// parse flags after all set.
	flag.Parse()

	// Print flag descriptions if not enough inputs
	if flag.NFlag() < 2 {
		flag.PrintDefaults()
		return
	}

	// split targets on comma
	targetsToScan := strings.Split(*targets, ",")
	// string slice of port range. scan() needs string port as input
	portsToScan := slicePorts(t_ports)

	// wait group for main proc to wait on goroutines
	var wg sync.WaitGroup

	// loop over targets
	for _, v := range targetsToScan {
		// inner loop over ports
		for _, k := range portsToScan {
			// add a count to wait group, wg.Done() decreases count
			wg.Add(1)
			// scan in goroutine
			go scan(v, k, timeout, &wg)
		}
	}
	// wait for wait group to be zero
	wg.Wait()
}
