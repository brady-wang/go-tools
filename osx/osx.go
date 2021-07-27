package osx

import (
	"errors"
	"fmt"
	"net"
)

func GetServerIp() (string,error){
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		return "",errors.New("获取服务器ip失败")
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String(),nil
					}
				}
			}
		}
	}

	return "",errors.New("未获取到服务器ip")
}