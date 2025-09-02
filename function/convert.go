package function

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

// linear interpolation between two colors
func lerpColor(c1, c2 color.RGBA, t float64) color.RGBA {
	r := uint8(float64(c1.R) + t*float64(c2.R-c1.R))
	g := uint8(float64(c1.G) + t*float64(c2.G-c1.G))
	b := uint8(float64(c1.B) + t*float64(c2.B-c1.B))
	return color.RGBA{r, g, b, 255} // force alpha = 255
}

func Convert(imagePath string) {
	// Open input file
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decoded format:", format) // jpeg, png, etc.

	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	// Define your duotone colors
	shadow := color.RGBA{27, 96, 47, 255}       // Hero Green
	highlight := color.RGBA{247, 132, 197, 255} // Brave Pink

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			// Normalize RGB to [0–1]
			rf := float64(r) / 65535.0
			gf := float64(g) / 65535.0
			bf := float64(b) / 65535.0

			// Compute brightness (luminance)
			gray := 0.299*rf + 0.587*gf + 0.114*bf

			// Map grayscale → between shadow and highlight
			newColor := lerpColor(shadow, highlight, gray)
			newImg.Set(x, y, newColor)
		}
	}

	// Save result
	ext := filepath.Ext(imagePath)
	outputPath := "duotone.jpg"
	out, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	switch ext {
	case ".png":
		err = png.Encode(out, newImg)
	default:
		err = jpeg.Encode(out, newImg, &jpeg.Options{Quality: 95})
	}

	if err != nil {
		log.Fatal(err)
	}
}
