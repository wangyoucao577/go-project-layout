package diagnosis

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/matishsiao/goInfo"
)

// Info includes information for diagnosis.
type Info struct {
	//static info
	Hostname    string   `json:"Hostname"`
	IPAddresses []string `json:"IP Addresses"` //readable format, like x.x.x.x or ::1
	CPUs        int      `json:"CPUs"`

	//dynamic info for per request
	RemoteAddr string `json:"Remote Endpoint"`
}

// New fetchs information for diagnosis.
func New() *Info {
	var info Info

	//hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("get hostname failed, err %v\n", err)
	} else {
		fmt.Printf("Hostname: %s\n", hostname)
		info.Hostname = hostname
	}

	//ip addresses
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("lookup network interface addrs failed, err %v\n", err)
	} else {
		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())
			if err != nil {
				fmt.Printf("ParseCIDR addrs failed, err %v\n", err)
				continue
			}

			if ip.IsLoopback() {
				fmt.Printf("ignore Loopback ip address-->%s\n", addr.String())
				continue
			}

			fmt.Printf("%s ip address-->%s\n", addr.Network(), addr.String())
			info.IPAddresses = append(info.IPAddresses, addr.String())
		}
	}

	//from goInfo
	gi := goInfo.GetInfo()
	gi.VarDump()
	info.CPUs = gi.CPUs

	return &info
}

func (info Info) String() string {
	jsonstr, err := json.Marshal(info)
	if err != nil {
		fmt.Printf("to json failed, err %v\n", err)
		return err.Error()
	}
	return string(jsonstr)
}
