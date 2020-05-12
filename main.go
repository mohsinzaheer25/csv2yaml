package main

import (
	_"bufio"
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	_"io"
	"os"
	"strings"
)

var (
	source string
	destination string
)

func init(){
	rootCmd.PersistentFlags().StringVarP(&source,"source", "s", "", "Source CSV Filename or File Path")
	rootCmd.PersistentFlags().StringVarP(&destination,"destination", "d", "", "Yaml File Name or Path To Save File.")
}

var rootCmd = &cobra.Command{
	Use:   "csv2yaml",
	Short: "csv2yaml is used to convert csv file to yaml",
	Long: `This tool is used to convert csv file to yaml

Usage: csv2yaml -s [csvfile] -d [yaml file]`,

	Run: func(cmd *cobra.Command, args []string) {
		if source == "" || destination == "" {
			cmd.Help()
			os.Exit(1)
		}
	},
}


func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if strings.HasSuffix(destination, ".yml") || strings.HasSuffix(destination, ".yaml") {
		if strings.HasSuffix(source, ".csv") {
			ReadCsv(source)
		} else {
			fmt.Println("Source file needs to be 'CSV'. We don't accept other extension.")
		}

	} else {
		fmt.Println("Please add '.yml' or '.yaml' to destination file")
	}
}


func ReadCsv(filename string) {

	// Open CSV file
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	r := csv.NewReader(f)
	result, _ := r.ReadAll()
	headersLength := result[0]

	// Writing Output To Yaml File.

	file, err := os.Create(destination)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(file, "---")

	for i :=0; i < (len(result) - 1); i++ {
		fmt.Fprintln( file, "- " + result[i+1][0] + ":")
		for length :=0; length < (len(headersLength)); length++ {
			if length <= (len(headersLength) - 2) {
				fmt.Fprintln( file, "    " + result[0][length+1] + ": " + result[i+1][length+1])
			}
		}
	}
	fmt.Println(destination + " File Created Successfully.")
}
