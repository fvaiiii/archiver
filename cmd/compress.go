/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/fvaiiii/archiver/internal/archive"
	"github.com/fvaiiii/archiver/internal/lz77"
	"github.com/spf13/cobra"
)

var defaultWindowSize int

// compressCmd represents the compress command
var compressCmd = &cobra.Command{
	Use:   "compress [input] [output]",
	Short: "Compress file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		output := args[1]

		if _, err := os.Stat(input); os.IsNotExist(err) {
			log.Fatalf("input file not found: %s", input)
		}

		data, err := os.ReadFile(input)
		if err != nil {
			log.Fatal(err)
		}
		tokens := lz77.Compress(data, defaultWindowSize)
		if err := archive.WriteArchive(output, tokens); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Compressed: %s -> %s", input, output)
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)
	compressCmd.Flags().IntVar(&defaultWindowSize, "window", 32768, "default window size: 32768")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
