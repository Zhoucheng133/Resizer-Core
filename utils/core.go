package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
)

func ResizeHandler(path string, width int, height int, output string, stretch bool) string {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Sprint("ERR: ", err.Error())
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Sprint("ERR: ", err.Error())
	}

	var finalImg image.Image
	ext := strings.ToLower(filepath.Ext(output))

	if stretch {
		finalImg = resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
	} else {
		origBounds := img.Bounds()
		ratioW := float64(width) / float64(origBounds.Dx())
		ratioH := float64(height) / float64(origBounds.Dy())

		ratio := ratioW
		if ratioH < ratio {
			ratio = ratioH
		}

		newW := uint(float64(origBounds.Dx()) * ratio)
		newH := uint(float64(origBounds.Dy()) * ratio)
		resizedSubImg := resize.Resize(newW, newH, img, resize.Lanczos3)

		var bgColor color.Color = color.Transparent
		if ext == ".jpg" || ext == ".jpeg" {
			bgColor = color.White
		}

		canvas := image.NewRGBA(image.Rect(0, 0, width, height))
		draw.Draw(canvas, canvas.Bounds(), &image.Uniform{bgColor}, image.Point{}, draw.Src)

		offsetX := (width - int(newW)) / 2
		offsetY := (height - int(newH)) / 2

		draw.Draw(canvas, image.Rect(offsetX, offsetY, offsetX+int(newW), offsetY+int(newH)), resizedSubImg, image.Point{}, draw.Over)
		finalImg = canvas
	}

	out, err := os.Create(output)
	if err != nil {
		return fmt.Sprint("ERR: ", err.Error())
	}
	defer out.Close()

	switch ext {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(out, finalImg, &jpeg.Options{Quality: 90})
	case ".png":
		err = png.Encode(out, finalImg)
	case ".gif":
		err = gif.Encode(out, finalImg, nil)
	case ".webp":
		err = webp.Encode(out, finalImg, &webp.Options{Lossless: true})
	default:
		err = png.Encode(out, finalImg)
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
