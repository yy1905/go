package publicutils

import "net"


func DnsToIP(uL string) ([]string, error) {
	ns, err := net.LookupHost(uL)
	if err != nil {
		return nil, err
	}
	return ns, err
}

