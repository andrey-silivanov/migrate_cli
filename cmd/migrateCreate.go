/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"time"
)

// migrateCreateCmd represents the migrateCreate command
var migrateCreateCmd = &cobra.Command{
	Use:     "migrateCreate",
	Aliases: []string{"migrate:create"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		fileName, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Println(err)
			return
		}

		err = migrateCreate(path, fileName)

		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("migrate " + fileName + " created")
	},
}

func init() {
	rootCmd.AddCommand(migrateCreateCmd)
	migrateCreateCmd.Flags().StringP("path", "p", "./migrations", "path to migrate folder")
	migrateCreateCmd.Flags().StringP("name", "n", "", "File name")
}

func migrateCreate(path string, name string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			log.Fatal(111, err)
			return err
		}
	}
	fullName := generateFileName(path, name)
	err := createUpFile(fullName, name)
	if err != nil {
		return err
	}

	err = createDownFile(fullName, name)

	if err != nil {
		return err
	}

	return nil
}

func generateFileName(path string, name string) string {
	now := time.Now()

	timestampStr := strconv.Itoa(int(now.Unix()))

	return path + "/" + name + "_" + timestampStr
}

func createUpFile(path string, name string) error {
	return os.WriteFile(path+"_up.sql", []byte("CREATE TABLE "+name+" (\n...\n);"), 0755)
}

func createDownFile(path string, name string) error {
	return os.WriteFile(path+"_down.sql", []byte("DROP TABLE "+name+";"), 0755)
}
