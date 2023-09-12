/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"go-microservice/server"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var httpGWPort int

// gatewayCmd represents the gateway command
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Create API gateway for service",
	Long:  `Create API gateway for service, it is used to consul for service discovery`,
	Run: func(cmd *cobra.Command, args []string) {
		httpFlag := cmd.Flag("gatewayport")
		httpGWPort, _ = strconv.Atoi(httpFlag.Value.String())

		consulFlag := cmd.Flag("consulport")
		consulPort, _ = strconv.Atoi(consulFlag.Value.String())

		server.LaunchGW(httpGWPort, consulPort)
	},
	PostRun: waitservicedone,
}

func init() {
	serviceCmd.AddCommand(gatewayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gatewayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gatewayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	gatewayCmd.Flags().IntVar(&httpGWPort, "gatewayport", 8090, "API gateway port")
	gatewayCmd.Flags().IntVar(&consulPort, "consulport", 8500, "Consul running port")
	gatewayCmd.MarkFlagRequired("consulport")

	viper.BindPFlag("gatewayport", gatewayCmd.Flags().Lookup("gatewayport"))
	viper.BindPFlag("consulport", gatewayCmd.Flags().Lookup("consulport"))
}
