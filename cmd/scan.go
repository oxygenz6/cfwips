package cmd

import (
	"fmt"
	"sync"

	"github.com/oxygenz6/cfwips/cfg"
	"github.com/oxygenz6/cfwips/lib"
	"github.com/oxygenz6/cfwips/lib/storage"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Start scanning IPs",
	Run: func(cmd *cobra.Command, args []string) {
		defer storage.Instance.Close()
		p, _ := ants.NewPool(4)
		defer p.Release()
		nets := lib.FetchCloudflareIPv4Networks()

		wg := &sync.WaitGroup{}
		for _, network := range nets {
			ips := lib.ConvertCIDRToIPs(*network)
			for _, ip := range ips {
				if storage.IpExists(ip) {
					continue
				}
				wg.Add(1)
				fn := func() {
					fmt.Printf("[%q] Scan started\n", ip.String())
					isClean := lib.IsCloudflareIpClean(*ip, cfg.Instance.SNI)
					storage.StoreIpStatus(ip, isClean)
					wg.Done()
					fmt.Printf("[%q] Scan finished, Result -> %t\n", ip.String(), isClean)
				}
				p.Submit(fn)
				if err := storage.Instance.Sync(); err != nil {
					panic(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
