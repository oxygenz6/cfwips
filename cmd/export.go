package cmd

import (
	"bytes"
	"fmt"

	"github.com/oxygenz6/cfwips/lib/storage"
	"github.com/spf13/cobra"
	"go.etcd.io/bbolt"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export list of (previously) found whitelisted IPs.",
	Long:  `You can redirect the output to any file you want and have the list exported in txt.`,
	Run: func(cmd *cobra.Command, args []string) {
		defer storage.Instance.Close()
		whitelistedIps := ""

		storage.Instance.View(func(tx *bbolt.Tx) error {
			b := tx.Bucket([]byte("ips"))
			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				if !bytes.Equal(v, []byte("true")) {
					continue
				}
				whitelistedIps += fmt.Sprintf("%s\n", k)
			}

			return nil
		})

		fmt.Print(whitelistedIps)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
