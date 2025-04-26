package cmds

import (
	"fmt"
	"io"
	"os"
	"strings"

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
		var inputs []string

		if len(args) == 0 {
			// Read from stdin if no args provided
			stdinBytes, err := io.ReadAll(os.Stdin)
			if err != nil {
				return fmt.Errorf("failed to read from stdin: %w", err)
			}

			// Split stdin by whitespace
			inputText := string(stdinBytes)
			inputs = strings.Fields(inputText)
		} else {
			inputs = args
		}

		if !binaryMode && !hexMode {
			return fmt.Errorf("you must specify either -b (binary) or -h (hex) mode")
		}

		return writer.WriteBytes(os.Stdout, inputs, hexMode, binaryMode)
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
