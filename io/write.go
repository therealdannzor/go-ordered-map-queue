package io

import (
	"fmt"
	"os"
	"time"
)

func WriteFile(data []byte) error {
	// filename is current timestamp
	name := time.Now().String()
	f, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("failed to create filename %s: %w", name, err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}

	return nil
}
