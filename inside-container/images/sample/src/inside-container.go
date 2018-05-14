package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

const (
	htmlStartToTitle = "<!doctype html><html><head><title>"
	htmlTitleToBody  = "</title></head><body>"
	htmlBodyToEnd    = "</body></html>"
	fmtOneInfoLine   = "<div><label>%s: </label><span>%s</span></div>"
)

func getHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = err.Error()
	}
	return hostname
}

func getIPAddresses() string {
	result := ""

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		result += err.Error()
	} else {
		for i, addr := range addrs {
			if i > 0 {
				result += ", "
			}
			result += addr.String()
		}
	}

	return result
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received.")

	fmt.Fprint(w, htmlStartToTitle)
	fmt.Fprint(w, "Inside Container")
	fmt.Fprint(w, htmlTitleToBody)
	fmt.Fprintf(w, fmtOneInfoLine, "Hostname", getHostName())
	fmt.Fprintf(w, fmtOneInfoLine, "IP Addresses", getIPAddresses())
	fmt.Fprintf(w, htmlBodyToEnd)

	fmt.Println("Request concluded.")
}

func main() {
	portNumber := os.Getenv("PORT_NUMBER")
	if portNumber == "" {
		portNumber = ":8080"
	} else {
		portNumber = ":" + portNumber
	}

	http.HandleFunc("/", rootHandler)

	fmt.Println("Hostname:" + getHostName())
	fmt.Println("IP Addresses:" + getIPAddresses())
	fmt.Printf("\n Starting server on %s\n", portNumber)

	http.ListenAndServe(portNumber, nil)
}
