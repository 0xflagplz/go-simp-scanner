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
	if len(os.Args) < 2 || len(os.Args) > 4{
		fmt.Println("\n\nUsage:", "go-simp-scanner.exe", "<HOST#>", "<StartingPort#>", "<EndingPort#>\n")
		fmt.Println("\nHOST#:\n\tReplace with Hostname or IP\n\t\tHOST EXAMPLES: scanme.nmap.org 127.0.0.1\n\nStartingPort#: \n\tReplace with the starting port number. This will only include ports on the set list.\n\nEndingPort#: \n\tReplace with the ending port number. This will only include ports on the set list.\n\nExample Outputs: \n\t$> go-simp-scanner.exe scanme.nmap.org 20 8080 \t|| \tThis will scan every preset port between 20 and 8080\n\n\n")
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
		print(art.String("Slow Ass Scanner") + "\n")	// Print Header Message


		port.GetOpenPorts(host, port.PortRange{Start:startV, End:finishV})		// IF all arguments are good, begin scan
		fmt.Println("\n\n\n" + strings.Repeat("*", 26))
		fmt.Println(strings.Repeat("*", 5) + "Scan Complete!" + strings.Repeat("*", 5))
		fmt.Println(strings.Repeat("*#", 13))
	}
	//Example Scan
	//port.GetOpenPorts("scanme.nmap.org", port.PortRange{Start:20, End:80})
}