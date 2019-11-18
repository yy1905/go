package publicutils

import (
	"net"
	"fmt"
	"os"
)

// 获取ip地址
func GetLocalIp() (ipslice []string) {
	var ipSlice []string

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//fmt.Println("ip:", ipnet.IP.String())
				ipSlice = append(ipSlice, ipnet.IP.String())
			}

		}
	}

	return ipSlice
}