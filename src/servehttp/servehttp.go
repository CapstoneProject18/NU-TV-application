package servehttp

import (
	"fmt"
	"net"
	"net/http"
)

func ServeHttp() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "172.19.16.239:6969")
	con, err := net.DialTCP("tcp", tcpAddr, nil)
	if err != nil {
		fmt.Errorf("hello%v", err)
	}
	_ = con
	http.ListenAndServe("127.0.0.1:8080", nil)
	http.HandleFunc("/", handleRoot)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
