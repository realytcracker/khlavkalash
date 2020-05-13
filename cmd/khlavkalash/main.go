/*

khlavkalash by ytcracker
use http to serve one thing
no matter what anyone asks for
no pizza, only khlav kalash

*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var verbose, help bool
var port uint
var mimetype, server, filename, location string

func main() {
	fmt.Println(" __   ___  __   ___  ")
	fmt.Println("|/\"| /  \")|/\"| /  \") ")
	fmt.Println("(: |/   / (: |/   /  ")
	fmt.Println("|    __/  |    __/   [khlavkalash] by [ytcracker]")
	fmt.Println("(// _  \\  (// _  \\   ")
	fmt.Println("|: | \\  \\ |: | \\  \\  ")
	fmt.Println("(__|  \\__)(__|  \\__) ")

	// program usage and defaults
	flag.BoolVar(&help, "h", false, "prints this helpful garbage")
	flag.BoolVar(&verbose, "v", false, "print incoming requests to stdout")
	flag.UintVar(&port, "p", 80, "port to serve khlav kalash on (80 may require root privs)")
	flag.StringVar(&mimetype, "m", "image/jpeg", "mime type of your khlav kalash")
	flag.StringVar(&server, "s", "nginx/1.17.10", "http server version header")
	flag.StringVar(&filename, "f", "", "path to your khlav kalash")
	flag.StringVar(&location, "l", "", "if set, 301 redirect to <parameter>, skip -f")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if fileExists(filename) == false && location == "" {
		flag.PrintDefaults()
		fmt.Println("please specify a file to serve with -f, or use the -l option.")
		os.Exit(1)
	}

	ln, err := net.Listen("tcp", ":"+strconv.FormatUint(uint64(port), 10))
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("server version header: " + server)
	if location != "" {
		fmt.Println("location: " + location)
	} else {
		fmt.Println("filename: " + filename)
		fmt.Println("mime type: " + mimetype)
	}
	fmt.Println("listening on port " + strconv.FormatUint(uint64(port), 10) + "...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go serveKhlavKalash(conn)
	}
}

// serveKhlavKalash serves the khlav kalash
func serveKhlavKalash(conn net.Conn) {
	// dunno why i keep date header
	// prob because rapper gotta know what time it is
	t := time.Now()

	headers := ""

	if location != "" {
		headers += "HTTP/1.1 301 Moved Permanently\n"
		headers += "Server: " + server + "\n"
		headers += "Date: " + t.Format(time.RFC1123) + "\n"
		headers += "Content-Type: " + mimetype + "\n"
		headers += "Location: " + location + "\n\n"
	} else {
		headers += "HTTP/1.1 200 OK\n"
		headers += "Server: " + server + "\n"
		headers += "Date: " + t.Format(time.RFC1123) + "\n"
		headers += "Content-Type: " + mimetype + "\n\n"
	}

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("error:", err.Error())
	} else if verbose == true {
		fmt.Println("[" + conn.RemoteAddr().String() + "]: " + string(buf))
	}

	f, err := ioutil.ReadFile(filename)

	conn.Write([]byte(headers))
	conn.Write(f)
	conn.Close()
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
