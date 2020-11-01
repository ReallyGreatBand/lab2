package lab2

import (
	"fmt"
	"io"
	"strings"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	// TODO: Add necessary fields.
	Reader io.Reader
	Writer io.Writer
}

func (ch *ComputeHandler) Compute() error {
	// TODO: Implement.
	buf := make([]byte, 1000)
	_, err := ch.Reader.Read(buf)
	if err != nil {
		return fmt.Errorf("Unexpected error")
	}
	res, err := CalculatePrefix(strings.Trim(string(buf), "\000"))
	if err != nil {
		return err
	}
	ch.Writer.Write([]byte(res))
	return nil
}
