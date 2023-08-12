package cmd

import (
	"fmt"
	"github.com/tim-lynn-clark/tlccrypto/aes_helpers"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var aesDecryptCmd = &cobra.Command{
	Use:   "decrypt -v value -k key",
	Short: "Decrypt a value with a key",
	Long:  `Decrypt a value with a key`,
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := cmd.Flags().GetString("key")
		if err != nil {
			return err
		}

		value, err := cmd.Flags().GetString("value")
		if err != nil {
			return err
		}

		return decrypt(key, value)
	},
}

func init() {
	aesCmd.AddCommand(aesDecryptCmd)

	aesDecryptCmd.Flags().StringP("key", "k", "", "Key to use for encryption")
	if err := aesDecryptCmd.MarkFlagRequired("key"); err != nil {
		fmt.Printf("Error marking key flag required: %v\n", err)
	}
	if err := viper.BindPFlag("key", aesDecryptCmd.Flags().Lookup("key")); err != nil {
		fmt.Printf("Error binding key flag: %v\n", err)
	}

	aesDecryptCmd.Flags().StringP("value", "v", "", "Value to encrypt")
	if err := aesDecryptCmd.MarkFlagRequired("value"); err != nil {
		fmt.Printf("Error marking value flag required: %v\n", err)
	}
	if err := viper.BindPFlag("value", aesDecryptCmd.Flags().Lookup("value")); err != nil {
		fmt.Printf("Error binding value flag: %v\n", err)
	}
}

func decrypt(key string, value string) error {
	var cryptoHelper aes_helpers.CryptoHelper

	if !cryptoHelper.IsInitialized {
		if err := cryptoHelper.InitEncryption(key); err != nil {
			return err
		}
	}

	//decrypting
	decryptedData, decryptErr := cryptoHelper.Decrypt(value)
	if decryptErr != nil {
		return decryptErr
	}

	fmt.Printf("Decrypted data: %s\n", decryptedData)
	return nil
}
