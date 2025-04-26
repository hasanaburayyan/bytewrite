package writer

import (
	"fmt"
	"io"
	"strconv"
)

func WriteBytes(output io.Writer, args []string, hexMode bool, binaryMode bool) error {
	if !hexMode && !binaryMode {
		return fmt.Errorf("either hexMode or binaryMode must be set")
	}

	if hexMode {
		if err := writeHexBytes(output, args); err != nil {
			return fmt.Errorf("failed to write hex bytes: %w", err)
		}
	} else if binaryMode {
		if err := writeBinaryBytes(output, args); err != nil {
			return fmt.Errorf("failed to write binary bytes: %w", err)
		}
	}

	return nil
}

func writeBinaryBytes(output io.Writer, args []string) error {
	for _, arg := range args {
		if len(arg) != 8 {
			return fmt.Errorf("invalid binary string %s: must be exactly 8 bits", arg)
		}

		val, err := strconv.ParseUint(arg, 2, 8)
		if err != nil {
			return err
		}

		_, err = output.Write([]byte{byte(val)})
		if err != nil {
			return err
		}
	}
	return nil
}

func writeHexBytes(output io.Writer, args []string) error {
	for _, arg := range args {
		if len(arg) != 2 {
			return fmt.Errorf("invalid binary string %s: must be exactly 2 hex digits", arg)
		}

		val, err := strconv.ParseUint(arg, 16, 8)
		if err != nil {
			return err
		}

		_, err = output.Write([]byte{byte(val)})
		if err != nil {
			return err
		}
	}
	return nil
}
