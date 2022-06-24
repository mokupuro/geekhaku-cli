package utils

import (
	"fmt"
	"os"
)

func PrintAAFromTxt(fp string) {
	b, err := os.ReadFile(fp)
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