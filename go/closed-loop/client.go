package main

import (
	"fmt"
	"net"
	"time"
)

type port string
type data string

var (
	udpPorts map[port]data
)

func main() {
    done := make(chan bool)
	// for each port
	for p, d := range udpPorts {
		fmt.Printf("starting client sending %s to port %s\n", d, p)
		// where each process sends data to a different target udp port
		serverAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("127.0.0.1:%s", p))
		checkErr(err)
		// from the same source
		localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
		checkErr(err)
		conn, err := net.DialUDP("udp", localAddr, serverAddr)
		checkErr(err)
		//defer conn.Close()
		// create a separate process
		go func(d data, c *net.UDPConn) {
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
		}(d, conn)
	}
    <-done
}

func init() {
	udpPorts = make(map[port]data)
	udpPorts["2561"] = "CMD"
	udpPorts["2563"] = "GTEL"
	udpPorts["2564"] = "GACK"
	udpPorts["2565"] = "ERR"
	udpPorts["2569"] = "SER"
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}