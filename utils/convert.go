package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrintAAFromTxt(fp string) {
	path, err := filepath.Abs(fp)
	if err != nil {
		fmt.Println(err)
	}

	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
}

func AAFromText(fp string) string {
	b, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
