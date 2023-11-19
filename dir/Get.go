package dir

import (
	"fmt"
	"os"
	"path"
)

func Get() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Failed to get home dir: %s\n", err)
		os.Exit(1)
	}

	return path.Join(home, "/.local/share/rucksack/")
}
