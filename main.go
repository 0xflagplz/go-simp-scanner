package main

import (
	"strings"
	"os"
	"fmt"
	"github.com/AchocolatechipPancake/go-simp-scanner/port"
	"github.com/zs5460/art"										//imported ASCII Art lib
	"strconv"
)


func main() {
	if len(os.Args) < 2{
		fmt.Println("Usage:", os.Args[0], "HOST", "StartingPort#", "EndingPort#")
		return
	} else {
		host := os.Args[1]
		start := os.Args[2]
		finish := os.Args[3]

		startV, err := strconv.Atoi(start)
		if err != nil {
			fmt.Println("ERROR# Start Port had an error")
			return
		}
		finishV, err := strconv.Atoi(finish)
		if err != nil {
			fmt.Println("ERROR# Finish Port had an error")
			return
		}
		print(art.String("Slow Ass Scanner") + "\n")


		port.GetOpenPorts(host, port.PortRange{Start:startV, End:finishV})
		fmt.Println("\n\n\n" + strings.Repeat("*", 26))
		fmt.Println(strings.Repeat("*", 5) + "Scan Complete!" + strings.Repeat("*", 5))
		fmt.Println(strings.Repeat("*#", 13))
	}
	//Example Scan
	//port.GetOpenPorts("scanme.nmap.org", port.PortRange{Start:20, End:80})
}
