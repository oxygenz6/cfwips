package lib

import (
	"errors"
	"io"
	"net"
	"net/http"
	"strings"
)

func FetchCloudflareIPv4Networks() []*net.IPNet {
	resp, err := http.DefaultClient.Get("https://www.cloudflare.com/ips-v4")
	if err != nil {
		panic(errors.Join(
			errors.New("failed to send http request"),
			err,
		))
	}
	netListRaw, err := io.ReadAll(resp.Body)
	if err != nil {

		panic(errors.Join(
			errors.New("failed to read http response"),
			err,
		))
	}
	strNets := strings.Split(string(netListRaw), "\n")
	ipNets := make([]*net.IPNet, 0)
	for _, v := range strNets {
		_, ipNet, err := net.ParseCIDR(v)
		if err != nil {
			panic(errors.Join(
				errors.New("failed to parse network"),
				err,
			))
		}
		ipNets = append(ipNets, ipNet)
	}
	return ipNets
}
