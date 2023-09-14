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

var gRPCPort int

// grpcserviceCmd represents the grpcservice command
var grpcserviceCmd = &cobra.Command{
	Use:   "grpcservice",
	Short: "Create service with gRPC",
	Long:  `Create service with gRPC`,
	Run: func(cmd *cobra.Command, args []string) {
		grpcFlag := cmd.Flag("grpcport")
		gRPCPort, _ = strconv.Atoi(grpcFlag.Value.String())
		server.LaunchgRPCSever(gRPCPort)
	},
	PostRun: waitServerDone,
}

func init() {
	serviceCmd.AddCommand(grpcserviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcserviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcserviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	grpcserviceCmd.Flags().IntVar(&gRPCPort, "grpcport", 8070, "Service gRPC port")
	viper.BindPFlag("grpcport", grpcserviceCmd.Flags().Lookup("grpcport"))
}
