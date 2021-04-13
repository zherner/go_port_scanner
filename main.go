package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

// scan attempts to connect to target on target ports
func scan(t string, tp []string, to time.Duration) {
	// range on slice of ports
	for _, v := range tp {
		// attempt to connect to target:port with a timeout
		conn, err := net.DialTimeout("tcp", t+":"+v, to)
		// closed port
		if err != nil {
			fmt.Printf("Port %s closed\n", v)
			continue
		}
		conn.Close()
		fmt.Printf("Port %s open\n", v)

		// for webserver ports
		webserver(t, v)
	}
}

// webserver prints the "Server" header if there is one
func webserver(t, p string) {
	if p == "80" || p == "443" {
		res, err := http.Get("http://" + t)
		if err != nil {
			log.Fatal(err)
		}
		// resp, _ := ioutil.ReadAll(res.Body)
		//res.Body.Close()
		// print "Server header value"
		fmt.Println("└─" + string(res.Header.Get("Server")))
	}
}

func main() {
	var (
		target_ip    = "151.101.194.152"
		target_ports = []string{"80", "443", "22", "8080"}
		target_timeo = time.Second
	)

	scan(target_ip, target_ports, target_timeo)

}
