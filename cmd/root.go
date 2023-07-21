/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-microservice",
	Short: "Build micro-service in Golang",
	Long:  `Build micro-service leverage go-kit framework`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	if len(os.Args) > 1 && os.Args[1] == "service" {
		errs := make(chan error)
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			errs <- fmt.Errorf("%s", <-c)
		}()
		fmt.Print("exit %p", <-errs)
	}
}

func init() {
	cobra.OnInitialize(initialize)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-microservice.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func initialize() {
	viper.SetConfigFile("go-service.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		viper.SetDefault("gatewayport", 8090)
		viper.SetDefault("httpport", 8080)
		viper.SetDefault("grpcport", 8070)
		viper.SetDefault("consulport", 8500)
	}
}

