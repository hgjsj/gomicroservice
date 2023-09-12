/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Launch micro service",
	Long:  `Launch micro service`,
	// Run: func(cmd *cobra.Command, args []string) {
	//  	cmd.Help()
	// },

}

func init() {
	rootCmd.AddCommand(serviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func waitservicedone(cmd *cobra.Command, args []string) {
	errs := make(chan error)
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("exit %p", <-errs)
}
