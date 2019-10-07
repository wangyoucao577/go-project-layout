package diagnosis

import (
	"net"
	"os"

	"github.com/golang/glog"
	"github.com/matishsiao/goInfo"
)

// New fetchs information for diagnosis.
func New() *Info {
	var info Info

	//hostname
	hostname, err := os.Hostname()
	if err != nil {
		glog.Warningf("get hostname failed, err %v\n", err)
	} else {
		glog.V(2).Infof("Hostname: %s\n", hostname)
		info.Hostname = hostname
	}

	//ip addresses
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		glog.Warningf("lookup network interface addrs failed, err %v\n", err)
	} else {
		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())
			if err != nil {
				glog.Warningf("ParseCIDR addrs failed, err %v\n", err)
				continue
			}

			if ip.IsLoopback() {
				glog.V(2).Infof("ignore Loopback ip address-->%s\n", addr.String())
				continue
			}

			glog.V(2).Infof("%s ip address-->%s\n", addr.Network(), addr.String())
			info.IPAddresses = append(info.IPAddresses, addr.String())
		}
	}

	//from goInfo
	gi := goInfo.GetInfo()
	glog.V(2).Infoln(gi)
	info.CPUs = gi.CPUs

	return &info
}
