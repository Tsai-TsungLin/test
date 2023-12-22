package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
)

type IPInfo struct {
	CIDR      string `json:"cidr"`
	Interface string `json:"interface"`
	IP        string `json:"ip"`
	MAC       string `json:"mac"`
	NetIndex  string `json:"net_index"`
	Netmask   string `json:"netmask"`
	Network   string `json:"network"`
	SetBy     string `json:"set_by"`
	Version   string `json:"version"`
}

func main() {
	ipInfos, err := ExportIPWindows()
	if err != nil {
		log.Println("ipInfos error: ", err)
	}
	log.Println("ipInfos: ", ipInfos)
}
func ExportIPWindows() ([]IPInfo, error) {
	var ipInfos []IPInfo

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			var cidr string
			var netmask, network string

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				cidr = strings.Split(v.String(), "/")[1]
				netmask = net.IP(v.Mask).String()
				network = v.IP.Mask(v.Mask).String()
			case *net.IPAddr:
				ip = v.IP
			}

			version := "4"
			if ip.To4() == nil {
				version = "6"
			}

			setBy := getSetByWindows(iface.Name)

			ipInfo := IPInfo{
				CIDR:      cidr,
				Interface: iface.Name,
				IP:        ip.String(),
				MAC:       iface.HardwareAddr.String(),
				NetIndex:  fmt.Sprintf("%d", iface.Index),
				Netmask:   netmask,
				Network:   network,
				SetBy:     setBy,
				Version:   version,
			}

			ipInfos = append(ipInfos, ipInfo)
		}
	}

	return ipInfos, nil
}

func getSetByWindows(interfaceName string) string {
	cmd := exec.Command("netsh", "interface", "ip", "show", "config", "name="+interfaceName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return "unknown"
	}

	outputString := string(output)
	if strings.Contains(outputString, "DHCP enabled: Yes") || strings.Contains(outputString, "DHCP 已啟用 . . . . . . . . . . . ：是") {
		return "dhcp"
	} else if strings.Contains(outputString, "DHCP enabled: No") || strings.Contains(outputString, "DHCP 已啟用 . . . . . . . . . . . ：否") {
		return "static"
	}

	return "unknown"
}
