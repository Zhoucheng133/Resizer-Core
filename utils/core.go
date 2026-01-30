package utils

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
)

func ResizeHandler(path string, width int, height int, output string) string {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Sprint("ERR: ", err.Error())
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Sprint("ERR: ", err.Error())
	}
	resized := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
	out, err := os.Create(output)
	if err != nil {
		return fmt.Sprint("ERR: ", err.Error())
	}
	defer out.Close()

	ext := strings.ToLower(filepath.Ext(output))
	switch ext {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(out, resized, &jpeg.Options{Quality: 90})
	case ".png":
		err = png.Encode(out, resized)
	case ".gif":
		err = gif.Encode(out, resized, nil)
	case ".webp":
		err = webp.Encode(out, resized, &webp.Options{Lossless: true})
	default:
		err = png.Encode(out, resized)
	}

	if err != nil {
		return fmt.Sprint("ERR: ", err.Error())
	}
	return "OK"
}

func GetSizeHandler(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Sprint("ERR: ", err.Error())
	}
	defer file.Close()
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return fmt.Sprint("ERR: ", err.Error())
	}
	return fmt.Sprint(config.Width, "x", config.Height)
}
