package writer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteBytes(t *testing.T) {
	t.Run("BinaryMode_Success", func(t *testing.T) {
		output := &bytes.Buffer{}
		args := []string{"00000001", "00000101", "00001010"}

		err := WriteBytes(output, args, false, true)
		require.NoError(t, err)

		expected := []byte{0x01, 0x05, 0x0A}
		assert.Equal(t, expected, output.Bytes())
	})

	t.Run("HexMode_Success", func(t *testing.T) {
		output := &bytes.Buffer{}
		args := []string{"01", "05", "0A"}

		err := WriteBytes(output, args, true, false)
		require.NoError(t, err)

		expected := []byte{0x01, 0x05, 0x0A}
		assert.Equal(t, expected, output.Bytes())
	})

	t.Run("InvalidBinaryLength", func(t *testing.T) {
		output := &bytes.Buffer{}
		args := []string{"00001"}

		err := WriteBytes(output, args, false, true)
		require.Error(t, err)
	})

	t.Run("InvalidBinaryContent", func(t *testing.T) {
		output := &bytes.Buffer{}
		args := []string{"0000AB01"}

		err := WriteBytes(output, args, false, true)
		require.Error(t, err)
	})

	t.Run("InvalidHexContent", func(t *testing.T) {
		output := &bytes.Buffer{}
		args := []string{"0G"}

		err := WriteBytes(output, args, true, false)
		require.Error(t, err)
	})

	t.Run("NoModeSelected", func(t *testing.T) {
		output := &bytes.Buffer{}
		args := []string{"01"}

		err := WriteBytes(output, args, false, false)
		require.Error(t, err)
	})

	t.Run("InvalidHexLength", func(t *testing.T) {
		output := &bytes.Buffer{}
		args := []string{"001"} // three characters, invalid hex

		err := WriteBytes(output, args, true, false)
		require.Error(t, err)
	})
}
