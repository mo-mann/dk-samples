package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
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

func getEnvironmentVariables() string {
	result := ""
	env := os.Environ()

	for _, variable := range env {
		result += variable + "\n"
	}

	return result
}

func getFiles(startpath string, depth int, currentdepth int) string {
	var result string

	entries, err := ioutil.ReadDir(startpath)
	if err != nil {
		result = err.Error()
		return result
	}

	for _, entry := range entries {
		if currentdepth > 0 {
			result += strings.Repeat("   ", currentdepth)

			if entry.IsDir() {
				result += "\\"
			} else {
				result += "|"
			}

			result += "--"
		}

		result += entry.Name() + "\n"

		if entry.IsDir() && currentdepth < depth {
			result += getFiles(path.Join(startpath, entry.Name()), depth, currentdepth+1)
		}
	}

	return result
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received.")

	fmt.Fprintln(w, "Hostname: "+getHostName())
	fmt.Fprintln(w, "IP Addresses: ", getIPAddresses())

	fmt.Println("Request concluded.")
}

func envHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, getEnvironmentVariables())
}

func filesHandler(w http.ResponseWriter, r *http.Request) {

	startpath := strings.TrimPrefix(r.URL.Path, "/files")
	if startpath == "" {
		startpath = "/"
	}

	depthqs := r.URL.Query()["depth"]
	depth := 1
	if len(depthqs) > 0 {
		depth, _ = strconv.Atoi(depthqs[0])
	}

	fmt.Fprintln(w, getFiles(startpath, depth, 0))
}

func main() {
	portNumber := os.Getenv("PORT_NUMBER")
	if portNumber == "" {
		portNumber = ":8080"
	} else {
		portNumber = ":" + portNumber
	}

	http.HandleFunc("/env", envHandler)
	http.HandleFunc("/files/", filesHandler)
	http.HandleFunc("/", rootHandler)

	fmt.Println("Hostname:" + getHostName())
	fmt.Println("IP Addresses:" + getIPAddresses())
	fmt.Printf("\n Starting server on %s\n", portNumber)

	http.ListenAndServe(portNumber, nil)
}
