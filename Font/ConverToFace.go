// 將字型來源檔轉換成程式內可以使用的字型
package font

import (
	"io"
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func ConvertToFace() font.Face {
	file, err := os.Open("./Font/NotoSansTC-Bold.otf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fontBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	f, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    20,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return face
}
