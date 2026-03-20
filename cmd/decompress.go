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

var force bool

// decompressCmd represents the decompress command
var decompressCmd = &cobra.Command{
	Use:   "decompress [input] [output]",
	Short: "Decompress file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		output := args[1]

		if !force {
			if _, err := os.Stat(output); err == nil {
				log.Fatalf("file already exists: %s (use --force to overwrite)", output)
			}
		}
		tokens, err := archive.ReadArchive(input)
		if err != nil {
			log.Fatal(err)
		}

		data := lz77.Decompress(tokens)

		if err := os.WriteFile(output, data, 0644); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Decompressed: %s -> %s (%d)", input, output, len(data))

	},
}

func init() {
	rootCmd.AddCommand(decompressCmd)
	decompressCmd.Flags().BoolVarP(&force, "force", "f", false, "overwrite output file, if exists")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decompressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decompressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
