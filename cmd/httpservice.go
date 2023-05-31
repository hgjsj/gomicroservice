/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go-microservice/server"
)

var httpPort int
var consulPort int
var proxy string

// httpserviceCmd represents the httpservice command
var httpserviceCmd = &cobra.Command{
	Use:   "httpservice",
	Short: "Create service with HTTP",
	Long:  `Create service with HTTP`,
	Run: func(cmd *cobra.Command, args []string) {
		httpPort := viper.GetInt("httpport")
		consulPort := viper.GetInt("consulport")

		proxyFlay := cmd.Flag("proxy")
		proxy = proxyFlay.Value.String()

		server.LaunchHttpSever(httpPort, proxy, consulPort)
		if consulPort != 0 {
			consulConfig := api.DefaultConfig()
			consulConfig.Address = fmt.Sprintf(":%d", consulPort)
			consulClient, _ := api.NewClient(consulConfig)
			client := consul.NewClient(consulClient)
			client.Register(&api.AgentServiceRegistration{
				ID:      fmt.Sprintf("stringservice:%d", consulPort),
				Name:    "stringservice",
				Tags:    append([]string{}, "stringservice"),
				Address: "127.0.0.1",
				Port:    consulPort})
		}
	},
}

func init() {
	serviceCmd.AddCommand(httpserviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpserviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpserviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	httpserviceCmd.Flags().IntVar(&httpPort, "httpport", 80, "Service HTTP port")
	httpserviceCmd.Flags().StringVar(&proxy, "proxy", "", "Proxy list splite with comma as 8010,8020,8030")
	httpserviceCmd.Flags().IntVar(&consulPort, "consulport", 8500, "Consul port")

	viper.BindPFlag("httpport", httpserviceCmd.Flags().Lookup("httpport"))
	viper.BindPFlag("proxy", httpserviceCmd.Flags().Lookup("proxy"))
	viper.BindPFlag("consulport", httpserviceCmd.Flags().Lookup("consulport"))
}
