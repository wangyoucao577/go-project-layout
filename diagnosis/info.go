package diagnosis

import (
	"encoding/json"

	"github.com/golang/glog"
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

func (info Info) String() string {
	jsonstr, err := json.Marshal(info)
	if err != nil {
		glog.Errorf("to json failed, err %v\n", err)
		return err.Error()
	}
	return string(jsonstr)
}
