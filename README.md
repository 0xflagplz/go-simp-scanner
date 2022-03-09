# go-simp-scanner
simple golang port scanner

Usage: > go-simp-scanner.exe <HOST#> <StartingPort#> <EndingPort#>

HOST#:   
          Replace with Hostname or IP  
             EXAMPLES: 
               scanme.nmap.org
               127.0.0.1 

StartingPort#:
          Replace with the starting port number. This will only include ports on the set list.
  
EndingPort#:
          Replace with the ending port number. This will only include ports on the set list.
          
          
   Example Outputs:
          go-simp-scanner.exe scanme.nmap.org 20 8080   \\ This will scan every preset port between 20 and 8080
