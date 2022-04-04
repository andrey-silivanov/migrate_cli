// Package cmd /*
package cmd

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
)

// migrateDownCmd represents the migrateDown command
var migrateDownCmd = &cobra.Command{
	Use:     "migrateDown",
	Aliases: []string{"migrate:down"},
	Short:   "migrate down",
	Long:    `migrate down`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")

		dbUrl, err := cmd.Flags().GetString("db-url")
		if err != nil {
			log.Println(err)
			return
		}

		err = migrateDown(path, dbUrl)
		if err != nil {
			log.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateDownCmd)
	migrateDownCmd.Flags().StringP("db-url", "u", "", "Database connection URL")
	migrateDownCmd.Flags().StringP("path", "p", "./migrations", "path to migrate folder")
}

func migrateDown(path string, dbUrl string) error {
	files, err := filepath.Glob(path + "/*_down.sql")
	if err != nil {
		return err
	}

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
