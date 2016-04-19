package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

)

var IP string

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><span style=\"font-size: 40px\">Hello world! IP is :%s</span></body></html>", IP)
	log.Print(r.URL.Path)
}

func Error404(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", http.StatusNotFound)
	log.Print(r.URL.Path)
}

func getip() string {
	conn, err := net.Dial("udp", "www.google.com.hk:80")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer conn.Close()

	fmt.Println(conn.LocalAddr().String())
	IP = strings.Split(conn.LocalAddr().String(), ":")[0]
	return IP
}

func main() {

	fmt.Println(getip())
	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/favicon.ico", Error404)
	log.Fatal(http.ListenAndServe(":80", nil))
}

