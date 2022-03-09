package main

import (
	"fmt"
	"go-simp-scan/port"
)

func main() {
	result := port.ScanPort("www.stackoverflow.com", 80)
	fmt.Println(result)
}