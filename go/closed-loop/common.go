package main

import (
	"fmt"
)

type port string
type data string

var udpPorts = map[port]data{
	"2561": "DIO",
	"2563": "GTEL",
	"2564": "GACK",
	"2565": "ERR",
	"2569": "SER",
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
