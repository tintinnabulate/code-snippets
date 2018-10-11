package main

import (
	"fmt"
	"net"
)

var (
	chans map[port](chan string)
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

	a, b, c, d, e := 0, 0, 0, 0, 0
	for {
		if a == 1 && b == 1 && c == 1 && d == 1 && e == 1 {
			fmt.Println("---")
			a, b, c, d, e = 0, 0, 0, 0, 0
		}
		select {
		case res := <-chans["2561"]:
			fmt.Println(res)
			a = 1
		case res := <-chans["2563"]:
			fmt.Println(res)
			b = 1
		case res := <-chans["2564"]:
			fmt.Println(res)
			c = 1
		case res := <-chans["2565"]:
			fmt.Println(res)
			d = 1
		case res := <-chans["2569"]:
			fmt.Println(res)
			e = 1
		}
	}

}

func init() {
	chans = make(map[port](chan string))
	chans["2561"] = make(chan string)
	chans["2563"] = make(chan string)
	chans["2564"] = make(chan string)
	chans["2565"] = make(chan string)
	chans["2569"] = make(chan string)
}
