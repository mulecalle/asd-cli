package cmd

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"notes-cli/utils"
	"strings"
)

var Commands = map[string]command{}
var ValidArgs []string

var note string
var domain string

//go:embed data/*
var StaticData embed.FS

var cmd = &cobra.Command{
	Use:   "notes",
	Long:  "use [flags] to get notes per domain",
	Short: "Get notes per domain",
	PreRun: func(cmd *cobra.Command, args []string) {
		// todo need to validate note and domains
	},
	Run: func(cmd *cobra.Command, args []string) {

		if domain == "" && note == "" {
			utils.PrintSliceString("DOMAIN", ValidArgs)
			return
		}

		runCommands()

	},
}

type command struct {
	path    string
	entries map[string]interface{}
}

func init() {
	initAvailableCommands()
	initValidArgs()

	// Disable completion from help
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(cmd)

	cmd.Flags().StringVarP(&note, "note", "n", "", "Set the name of the note to output")
	cmd.Flags().StringVarP(&domain, "domain", "d", "", "List notes of a given domain")
}

func runCommands() {
	var value interface{}
	if domain != "" && note != "" {
		value = Commands[domain].entries[note]
	} else if domain != "" {
		value = Commands[domain].entries
	}

	s, err := utils.MapValueToYAMLString(value)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(s)
}

func initValidArgs() {
	for key := range Commands {
		ValidArgs = append(ValidArgs, key)
	}
}

func initAvailableCommands() {
	Commands = make(map[string]command)

	dir := "data"

	files, err := StaticData.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		parts := strings.Split(file.Name(), ".")
		filePath := dir + "/" + file.Name()
		entries := getEntriesFromFile(filePath)

		Commands[parts[0]] = command{
			path:    filePath,
			entries: entries,
		}
	}
}

func getEntriesFromFile(path string) map[string]interface{} {
	// Read the YAML file
	yamlFile, err := StaticData.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return utils.SortEntries(yamlFile)
}
