package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var s string = ""
	for i := 0; i < len(ip); i++ {
		s = s + fmt.Sprintf("%v", ip[i]) + "."
	}
	s = s[:len(s)-1]
	return s
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
