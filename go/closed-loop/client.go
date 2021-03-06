package main

import (
	"fmt"
	"net"
	"time"
)

var (
	done chan bool
)

func init() {
	done := make(chan bool)
}

func client(d data, c *net.UDPConn) {
	i := 0
	for {
		msg := fmt.Sprintf("%s: %d", d, i)
		i++
		buf := []byte(msg)
		_, err := c.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
	done <- true
}

func main() {
	x := 1
	localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	checkErr(err)
	// for each port
	for p, d := range udpPorts {
		fmt.Printf("starting client sending %s to port %s\n", d, p)
		// where each process sends data to a different target udp port
		serverAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("127.0.0.1:%s", p))
		checkErr(err)
		conn, err := net.DialUDP("udp", localAddr, serverAddr)
		checkErr(err)
		//defer conn.Close()
		// create a separate process after x++ seconds
		time.Sleep(time.Duration(int(time.Second) * x))
		x++
		go client(d, conn)
	}
	<-done
}
