package main

import (
	"flag"
	"fmt"
)

var PersistAddr string
var AIAddr string 
func init() {
	flag.StringVar(&PersistAddr, "persist", "", "Address Persist service can be reached at")
	flag.StringVar(&AIAddr, "ai", "", "Address AI service can be reached at")
}

func main() {
	flag.Parse()
	fmt.Println("Persist", PersistAddr)
	fmt.Println("AI", AIAddr)
}