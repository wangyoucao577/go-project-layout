package diagnosis

import (
	"log"
	"net"
	"os"

	"github.com/matishsiao/goInfo"
)

// New fetchs information for diagnosis.
func New() *Info {
	var info Info

	//hostname
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("get hostname failed, err %v\n", err)
	} else {
		log.Printf("Hostname: %s\n", hostname)
		info.Hostname = hostname
	}

	//ip addresses
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Printf("lookup network interface addrs failed, err %v\n", err)
	} else {
		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())
			if err != nil {
				log.Printf("ParseCIDR addrs failed, err %v\n", err)
				continue
			}

			if ip.IsLoopback() {
				log.Printf("ignore Loopback ip address-->%s\n", addr.String())
				continue
			}

			log.Printf("%s ip address-->%s\n", addr.Network(), addr.String())
			info.IPAddresses = append(info.IPAddresses, addr.String())
		}
	}

	//from goInfo
	gi := goInfo.GetInfo()
	log.Println(gi)
	info.CPUs = gi.CPUs

	return &info
}
