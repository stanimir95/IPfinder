package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Print("Enter Domain Name: ")
	var input string
	fmt.Scanln(&input)

	iprecords, _ := net.LookupIP(input)
	for _, ip := range iprecords {
		fmt.Println(ip)
	}
}
