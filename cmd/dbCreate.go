/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"log"
)

// dbCreateCmd represents the dbCreate command
var dbCreateCmd = &cobra.Command{
	Use:     "dbCreate",
	Aliases: []string{"db:create"},
	Short:   "Create database",
	Long:    `Create database`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("data base create started")
		dbUrl, err := cmd.Flags().GetString("db-url")
		if err != nil {
			log.Println(err)
			return
		}
		dbName, err := cmd.Flags().GetString("db-name")
		if err != nil {
			log.Println(err)
			return
		}

		err = createDataBase(dbUrl, dbName)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("data base created")
	},
}

func init() {
	rootCmd.AddCommand(dbCreateCmd)
	dbCreateCmd.Flags().StringP("db-url", "u", "", "Database connection URL")
	dbCreateCmd.Flags().StringP("db-name", "n", "", "Database name")
}

func createDataBase(dbUrl string, dbName string) error {

	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		return err
	}

	return nil
}
