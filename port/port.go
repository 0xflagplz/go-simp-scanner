
import (
	"fmt"
	"strconv"
)


common := map[int]string {

	"20": "alt-ftp",
	"21": "FTP",
	"22": "SSH",
	"23": "TELNET",
	"25": "SMTP",
	"53": "DNS",
	"68": "DCHP",
	"67": "DCHP",
	"69": "TFTP",
	"80": "HTTP",
	"109": "POP2",
	"110": "POP3",
	"122": "alt-ssh",
	"123": "NTP",
	"135": "EPMAP / RPC Locator service",
	"137": "NETBIOS-NS",
	"138": "NETBIOS-DGM",
	"139": "NETBIOS-SSN--SMB",
	"143": "IMAP",
	"156": "SQL-SERVER",
	"161": "SNMP",
	"162": "SNMPTRAP",
	"220": "IMAPv3",
	"280": "http-mgmt",
	"369": "Rpc2portmap",
	"389": "LDAP",
	"443": "HTTPS",
	"445": "SMB",
	"465": "alt-smtp",
	"513": "rlogin",
	"530": "rpc",
	"540": "uucp",
	"543": "Kerberos login",
	"544": "Kerberos Remote shell",
	"546": "DHCP-CLIENT",
	"554": "rtsp",
	"563": "NNTPS",
	"587": "smtp-alt",
	"547": "DHCP-SERVER",
	"591": "FileMaker 6.0",
	"593": "HTTP RPC",
	"636": "LDAPS",
	"873": "rsync",
	"902" : "VMWARE",
	"989": "FTPS",
	"990": "FTPS",
	"995": "POP3-SSL",
	"993": "IMAP-SSL",
	"1080": "SOCKS proxy",
	"1119": "Battle.net Chat/Game",
	"1241": "Unofficial-Nessus Scanner",
	"1194": "OpenVPN",
	"1352": "IBM Lotus Notes/Domino (RPC) protocol",
	"1360": "Mimer SQL",
	"1414": "IBM WebSphere MQ",
	"1433": "MSSQL Server",
	"1434": "MSSQL Monitor",
	"1527": "Oracle Net Services",
	"2086": "WHM/CPANEL",
	"2087": "WHM/CPANEL",
	"2082": "CPANEL",
	"2083": "CPANEL",
	"3306": "MYSQL-Database",
	"3389": "RDP",
	"5000": "UNPN",
	"8443": "alt-https",
	"8080": "alt-http",
	"9090": "alt-http/updog",
	"10000": "VIRTUALMIN/WEBMIN",
	
}

func ScanPort(hostname string, port int) bool{
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", address, 60*time.Second)  // modify the "tcp" in order to create udp option
	if err != nil {
		return false
	}
	defer conn.Close()  // defer will wait till the function is done
	return true
}
}