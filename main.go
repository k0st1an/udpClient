package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"
)

var (
	address string
	loop    bool
	msg     string
	msgFile string
)

func main() {
	bytes, err := ioutil.ReadFile(msgFile)
	checkErr(err)
	msg = string(bytes)

	rAddr, err := net.ResolveUDPAddr("udp", address)
	checkErr(err)

	lAddr, err := net.ResolveUDPAddr("udp", "localhost:0")
	checkErr(err)

	client, err := net.DialUDP("udp", lAddr, rAddr)
	checkErr(err)
	defer client.Close()

	if loop {
		for {
			client.Write([]byte(msg))
			time.Sleep(time.Second)
		}
	} else {
		client.Write([]byte(msg))
	}
}

func init() {
	flag.StringVar(&address, "addr", "", "address of UDP server")
	flag.BoolVar(&loop, "loop", false, "send messages endlessly")
	flag.StringVar(&msgFile, "msg-file", "./msg.txt", "file with a message")
	flag.Parse()

	if len(strings.TrimSpace(address)) == 0 {
		log.Fatalln("use -addr key")
	}
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
