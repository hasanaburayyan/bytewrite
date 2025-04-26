package cmd

import (
	"fmt"
	"os"

	"github.com/hasanaburayyan/bytewrite/src/internal/writer"
	"github.com/spf13/cobra"
)

var (
	hexMode    bool
	binaryMode bool
)

var rootCmd = &cobra.Command{
	Use:   "bytewrite",
	Short: "Write real binary bytes from human-readable input",
	Long: `Bytewrite is a tool for generating true binary files from binary or hexadecimal strings.
Type real bits or hex, not text characters.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("no input provided")
		}

		if !binaryMode && !hexMode {
			return fmt.Errorf("you must specify either -b (binary) or -h (hex) mode")
		}

		return writer.WriteBytes(os.Stdout, args, hexMode, false)
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&hexMode, "hex", "x", false, "Interpret input as hexadecimal")
	rootCmd.PersistentFlags().BoolVarP(&binaryMode, "binary", "b", false, "Interpret input as binary (default mode)")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
