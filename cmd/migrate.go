/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")

		dbUrl, err := cmd.Flags().GetString("db-url")
		if err != nil {
			log.Println(err)
			return
		}

		err = migrate(path, dbUrl)
		if err != nil {
			log.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().StringP("db-url", "u", "", "Database connection URL")
	migrateCmd.Flags().StringP("path", "p", "./migrations", "path to migrate folder")
}

func migrate(path string, dbUrl string) error {
	files, err := filepath.Glob(path + "/*_up.sql")
	if err != nil {
		return err
	}
	fmt.Println(files)
	for _, file := range files {
		fContent, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		err = executeSql(dbUrl, string(fContent))
		if err != nil {
			return err
		}
	}

	return nil
}

func executeSql(dbUrl string, sqlQuery string) error {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(sqlQuery)
	if err != nil {
		return err
	}

	return nil
}
