package lib

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"
)

func IsCloudflareIpClean(ip net.IP, sni string) bool {
	if sni == "" {
		sni = "www.cloudflare.com"
	}
	dialer := &net.Dialer{
		Timeout: 10 * time.Second,
	}
	httpC := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				ServerName:         sni,
				InsecureSkipVerify: false,
			},
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				if addr == fmt.Sprintf("%s:443", sni) {
					addr = fmt.Sprintf("%s:443", ip.String())
				}
				return dialer.DialContext(ctx, network, addr)
			},
		},
	}
	_, err := httpC.Get(fmt.Sprintf("https://%s", sni))
	return err == nil
}
