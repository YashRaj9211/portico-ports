package cmd

import (
	"fmt"
	"github.com/YashRaj9211/portico/portico-client/utils"
	"github.com/spf13/cobra"
)


var localPort string

var exposeCmd = &cobra.Command{
	Use: "expose [localPort]",
	Short: "Expose localport to the internet and access it\n publically using assigned url",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		localPort := args[0]
		fmt.Printf("Exposing local port: %s\n", localPort)
		url, err := utils.ClientSideTunneler(localPort)
		if err != nil {
			fmt.Printf("Tunnel failed: %v\n", err)
			return
		}
		fmt.Printf("Access your port at: %s\n", url)
		select {}
	},
}


func init() {
	rootCmd.AddCommand(exposeCmd)
}