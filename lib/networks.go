package lib

import (
	"encoding/binary"
	"net"
)

func ConvertCIDRToIPs(network net.IPNet) []*net.IP {
	mask := binary.BigEndian.Uint32(network.Mask)
	start := binary.BigEndian.Uint32(network.IP)

	// find the final address
	finish := (start & mask) | (mask ^ 0xffffffff)

	// loop through addresses as uint32
	ips := make([]*net.IP, 0)
	for i := start; i <= finish; i++ {
		// convert back to net.IP
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		ips = append(ips, &ip)
	}
	return ips
}
