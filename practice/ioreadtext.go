package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteTo(w io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}

// Write
func ExampleWriteTo() {
	lines := []string{
		"ialdkjflaksdj",
		"tom",
		"jane",
	}
	if err := WriteTo(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}
}
