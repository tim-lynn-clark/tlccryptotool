package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "tlccryptotool",
		Short: "tlccryptotool provides methods for encrypting and decrypting string data.",
		Long:  `tlccryptotool provides methods for encrypting and decrypting string data.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	viper.SetDefault("author", "Tim Clark")
	viper.SetDefault("license", "MIT")

	rootCmd.Version = "0.1.0"
	versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
	rootCmd.SetVersionTemplate(versionTemplate)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tlccrypto.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	if err := viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author")); err != nil {
		fmt.Printf("Error binding author flag: %s", err)
	}
	if err := viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper")); err != nil {
		fmt.Printf("Error binding viper flag: %s", err)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".tlccrypto")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
