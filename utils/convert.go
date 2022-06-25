package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrintAAFromTxt(fp string) {

	p, _ := os.Getwd()
	fmt.Println(p)
	b, err := os.ReadFile(filepath.Join(p, "aa", fp))
	if err != nil {
		fmt.Println("read file")
		fmt.Println(err)
	}
	fmt.Print(string(b))
}

func AAFromText(fp string) string {

	p, _ := os.Getwd()
	fmt.Println(p)
	b, err := os.ReadFile(filepath.Join(p, "aa", fp))
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
