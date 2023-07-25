/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"go-microservice/service"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dbpath string

// migrationCmd represents the migration command
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Migrate database schema",
	Long:  `Migrate database schema`,
	Run: func(cmd *cobra.Command, args []string) {
		dbpath = viper.GetString("dbpath")
		service.InitSQLit(path.Join(dbpath, "cloud.db"))
		service.DBMigrationAll()
	},
}

func init() {
	dbCmd.AddCommand(migrationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	migrationCmd.Flags().StringVar(&dbpath, "dbpath", ".", "path of db data")
	viper.BindPFlag("dbpath", migrationCmd.Flags().Lookup("dbpath"))

}
