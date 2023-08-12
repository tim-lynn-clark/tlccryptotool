package cmd

import (
	"fmt"
	"github.com/tim-lynn-clark/tlccrypto/aes_helpers"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var aesEncryptCmd = &cobra.Command{
	Use:   "encrypt -v value -k key",
	Short: "Encrypt a value with a key",
	Long:  `Encrypt a value with a key`,
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := cmd.Flags().GetString("key")
		if err != nil {
			return err
		}

		value, err := cmd.Flags().GetString("value")
		if err != nil {
			return err
		}

		return encrypt(key, value)
	},
}

func init() {
	aesCmd.AddCommand(aesEncryptCmd)

	aesEncryptCmd.Flags().StringP("key", "k", "", "Key to use for encryption")
	if err := aesEncryptCmd.MarkFlagRequired("key"); err != nil {
		fmt.Printf("Error marking key flag required: %v\n", err)
	}
	if err := viper.BindPFlag("key", aesEncryptCmd.Flags().Lookup("key")); err != nil {
		fmt.Printf("Error binding key flag: %v\n", err)
	}

	aesEncryptCmd.Flags().StringP("value", "v", "", "Value to encrypt")
	if err := aesEncryptCmd.MarkFlagRequired("value"); err != nil {
		fmt.Printf("Error marking value flag required: %v\n", err)
	}
	if err := viper.BindPFlag("value", aesEncryptCmd.Flags().Lookup("value")); err != nil {
		fmt.Printf("Error binding value flag: %v\n", err)
	}
}

func encrypt(key string, value string) error {
	fmt.Printf("Key: %s\n", key)
	fmt.Printf("Value: %s\n", value)

	var cryptoHelper aes_helpers.CryptoHelper

	if !cryptoHelper.IsInitialized {
		if err := cryptoHelper.InitEncryption(key); err != nil {
			return err
		}
	}

	//encrypting
	encryptedData, encryptErr := cryptoHelper.Encrypt(value)
	if encryptErr != nil {
		return encryptErr
	}

	fmt.Printf("Encrypted data: %s\n", encryptedData)
	return nil
}
