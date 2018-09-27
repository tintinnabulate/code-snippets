package main

import (
	"fmt"
	"net"
)

type port string
type data string

var (
	udpPorts map[port]data
	chans    map[port](chan string)
)

func main() {
	for p, _ := range udpPorts {
		fmt.Printf("starting server listening on port %s\n", p)
		serverAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%s", p))
		checkErr(err)
		serverConn, err := net.ListenUDP("udp", serverAddr)
		checkErr(err)
		defer serverConn.Close()
		go func(p port, c chan<- string, conn *net.UDPConn) {
			for {
				buf := make([]byte, 1024)
				n, addr, err := serverConn.ReadFromUDP(buf)
				c <- fmt.Sprintf("%s Received %s from %s", p, string(buf[0:n]), addr)
				if err != nil {
					fmt.Println("Error: ", err)
				}
			}
		}(p, chans[p], serverConn)
	}

	for {
		select {
		case res := <-chans["2561"]:
			fmt.Println(res)
		case res := <-chans["2563"]:
			fmt.Println(res)
		case res := <-chans["2564"]:
			fmt.Println(res)
		case res := <-chans["2565"]:
			fmt.Println(res)
		case res := <-chans["2569"]:
			fmt.Println(res)
		}
	}

}

func init() {
	udpPorts = make(map[port]data)
	udpPorts["2561"] = "CMD"
	udpPorts["2563"] = "GTEL"
	udpPorts["2564"] = "GACK"
	udpPorts["2565"] = "ERR"
	udpPorts["2569"] = "SER"
	chans = make(map[port](chan string))
	chans["2561"] = make(chan string)
	chans["2563"] = make(chan string)
	chans["2564"] = make(chan string)
	chans["2565"] = make(chan string)
	chans["2569"] = make(chan string)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
