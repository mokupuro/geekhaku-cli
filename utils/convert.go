package utils

import (
	"bytes"
	"fmt"
	"path/filepath"

	_ "github.com/mokupuro/geekhaku-cli/statik"
	"github.com/rakyll/statik/fs"
)

func PrintAAFromTxt(fp string) {
	text := readFile(filepath.Join("/", fp))
	fmt.Print(text)
}

func AAFromText(fp string) string {
	text := readFile(filepath.Join("/", fp))
	return text
}

func readFile(fp string) string {
	statikFS, err := fs.New()
	if err != nil {
		fmt.Println("************************:")
		fmt.Println(err)
		fmt.Println("************************:")
	}

	file, err := statikFS.Open(fp)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// ファイルを読み込んで出力
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	return buf.String()
}
