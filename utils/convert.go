package utils

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/mokupuro/geekhaku-cli/static"
)

func PrintAAFromTxt(fp string) {
	text := readFile(filepath.Join("aa", fp))
	fmt.Print(text)
}

func AAFromText(fp string) string {
	text := readFile(filepath.Join("aa", fp))
	return text
}

func readFile(fp string) string {
	file, err := static.Aa.Open(fp)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// ファイルを読み込んで出力
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	return buf.String()
}
