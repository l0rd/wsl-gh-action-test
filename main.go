package main

import (
	"fmt"

	"github.com/vishvananda/netlink"
)

const defaultWSLRoute = "0.0.0.0/0"

// wslHostIP returns the Windows host's IP address. It only make
// sense to execute it when running in a WSL distribution.
// Instructions to retrieve the IP address are section "Identify IP address"
// (scenario 2) of the WSL networking documentation:
// https://learn.microsoft.com/en-us/windows/wsl/networking#identify-ip-address
func wslHostIP() string {
	routes, err := netlink.RouteList(nil, netlink.FAMILY_V4)
	if err != nil {
		return ""
	}
	for _, r := range routes {
		if r.Dst.String() == defaultWSLRoute && r.Gw != nil {
			return r.Gw.String()
		}
	}
	return ""
}

func main() {
	hip := wslHostIP()
	fmt.Println(hip)
}
