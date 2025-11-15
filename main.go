package main

import (
	"fmt"
	"log"
	"net"

	"github.com/FAMASoon/IPObfuscator-go/obfuscator"
)

var addr string

func main() {
	fmt.Printf("Enter IP Address: ")
	fmt.Scanf("%s", &addr)

	ipv4 := net.ParseIP(addr).To4()
	if ipv4 == nil {
		log.Fatal(fmt.Errorf("%s is not IP Address", addr))
	}

	fmt.Printf("%s\n", obfuscator.IPtoDecimal(ipv4))
	fmt.Printf("%s\n", obfuscator.IPtoHex(ipv4))
	fmt.Printf("%s\n", obfuscator.IPtoOct(ipv4))
	fmt.Printf("%s\n", obfuscator.IPtoHexPadding(ipv4))
	fmt.Printf("%s\n", obfuscator.IPtoOctPadding(ipv4))
	fmt.Println()

	fmt.Printf("%s\n", obfuscator.IPtoHexAndDec(ipv4))
	fmt.Printf("%s\n", obfuscator.IPtoHexAndDec2(ipv4))
	fmt.Printf("%s\n", obfuscator.IPtoHexAndDec3(ipv4))
	fmt.Println()

	fmt.Printf("%s\n", obfuscator.IPtoOctAndDec(ipv4))
	fmt.Printf("%s\n", obfuscator.IPtoOctAndDec2(ipv4))
	fmt.Printf("%s\n", obfuscator.IPtoOctAndDec3(ipv4))
	fmt.Println()

	fmt.Printf("%s\n", obfuscator.HexAndDecimal(ipv4))
	fmt.Printf("%s\n", obfuscator.OctAndDecimal(ipv4))
	fmt.Printf("%s\n", obfuscator.HexOctDec(ipv4))
	fmt.Printf("%s\n", obfuscator.HexAndDec(ipv4))
	fmt.Printf("%s\n", obfuscator.OctAndDec(ipv4))
	fmt.Printf("%s\n", obfuscator.Hex2Oct2(ipv4))
	fmt.Printf("%s\n", obfuscator.HexOct3(ipv4))
}
