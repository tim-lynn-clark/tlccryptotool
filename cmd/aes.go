package cmd

import "github.com/spf13/cobra"

var aesCmd = &cobra.Command{
	Use:   "aes",
	Short: "AES encryption helper methods",
	Long:  `AES encryption helper methods`,
}

func init() {
	rootCmd.AddCommand(aesCmd)
}
