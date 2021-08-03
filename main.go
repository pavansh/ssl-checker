package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

func main() {

	host := flag.String("host", "example.com", "hostname (ex: -host example.com,sample.com")

	flag.Parse()

	hostname := strings.Split(*host, ",")

	for _, value := range hostname {

		conn, err := tls.Dial("tcp", value+":443", nil)
		if err != nil {
			log.Fatal("SSL Not Found: " + err.Error())
		}

		err = conn.VerifyHostname(value)
		if err != nil {
			log.Fatal("Certificate Error: " + err.Error())
		}

		expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
		fmt.Printf("Host: %s\nIssuer: %s\nExpiry: %v\n", value, conn.ConnectionState().PeerCertificates[0].Issuer, expiry.Format(time.RFC850))
		fmt.Println("Duration: ", humanize.Time(conn.ConnectionState().PeerCertificates[0].NotAfter)+"\n")
	}
}
