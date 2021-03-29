package internal

import (
	"fmt"
	"io"
)

func writeText(of io.Writer, content string) error {
	_, err := fmt.Fprint(of, content)
	return err
}

func writePdf(of io.Writer, content string) error {
	err := GeneratePdf(of, content)
	if err != nil {
		panic(err)
	}
	return err
}
