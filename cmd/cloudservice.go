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
	"go-microservice/service"
	"path"
)

var cloudPort int

// httpserviceCmd represents the httpservice command
var cloudserviceCmd = &cobra.Command{
	Use:   "cloudservice",
	Short: "Create cloud service with HTTP",
	Long:  `Create cloud service with HTTP`,
	Run: func(cmd *cobra.Command, args []string) {
		httpPort := viper.GetInt("cloudport")
		consulPort := viper.GetInt("consulport")
		var err error
		if dbpath, err = cmd.Flags().GetString("dbpath"); err != nil || dbpath == ""{
			dbpath = viper.GetString("dbpath")
		}
		
		dbpath = path.Join(dbpath, "cloud.db")
		service.InitSQLit(dbpath)
		service.DBMigrationAll()
		server.LauchCloudServer(httpPort)
		if consulPort != 0 {
			consulConfig := api.DefaultConfig()
			consulConfig.Address = fmt.Sprintf(":%d", consulPort)
			consulClient, _ := api.NewClient(consulConfig)
			client := consul.NewClient(consulClient)
			client.Register(&api.AgentServiceRegistration{
				ID:      fmt.Sprintf("cloudservice:%d", consulPort),
				Name:    "cloudservice",
				Tags:    append([]string{}, "cloudservice"),
				Address: "127.0.0.1",
				Port:    consulPort})
		}
	},
}

func init() {
	serviceCmd.AddCommand(cloudserviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpserviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpserviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cloudserviceCmd.Flags().IntVar(&cloudPort, "cloudport", 8060, "Cloud Service HTTP port")
	cloudserviceCmd.Flags().IntVar(&consulPort, "consulport", 8500, "Consul port")
	cloudserviceCmd.Flags().StringVar(&dbpath, "dbpath", "", "path of db data")
	viper.BindPFlag("dbpath", migrationCmd.Flags().Lookup("dbpath"))
	viper.BindPFlag("cloudport", cloudserviceCmd.Flags().Lookup("httpport"))
	viper.BindPFlag("consulport", cloudserviceCmd.Flags().Lookup("consulport"))
}
