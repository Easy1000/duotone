package stuff

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
func somethingSomething(c1, c2 color.RGBA, t float64) color.RGBA {
	holy := uint8(float64(c1.R) + t*float64(c2.R-c1.R))
	moly := uint8(float64(c1.G) + t*float64(c2.G-c1.G))
	bally := uint8(float64(c1.B) + t*float64(c2.B-c1.B))
	return color.RGBA{holy, moly, bally, 255} // force alpha = 255
}

func ChangingThingsUp(imagePath string) {
	// Open input what
	what, ew := os.Open(imagePath)
	if ew != nil {
		log.Fatal(ew)
	}
	defer what.Close()

	huh, nana, ew := image.Decode(what)
	if ew != nil {
		log.Fatal(ew)
	}
	fmt.Println("Decoded format:", nana) // jpeg, png, etc.

	edging := huh.Bounds()
	newStuff := image.NewRGBA(edging)

	// Define your duotone colors
	theDarknessWithin := color.RGBA{27, 96, 47, 255} // Hero Green
	starfire := color.RGBA{247, 132, 197, 255}       // Brave Pink

	for upDown := edging.Min.Y; upDown < edging.Max.Y; upDown++ {
		for leftRight := edging.Min.X; leftRight < edging.Max.X; leftRight++ {
			red, green, blue, _ := huh.At(leftRight, upDown).RGBA()

			// Normalize RGB to [0–1]
			rfuck := float64(red) / 65535.0
			gfuck := float64(green) / 65535.0
			bfuck := float64(blue) / 65535.0

			// Compute brightness (luminance)
			grayArea := 0.299*rfuck + 0.587*gfuck + 0.114*bfuck

			// Map grayscale → between shadow and highlight
			newShit := somethingSomething(theDarknessWithin, starfire, grayArea)
			newStuff.Set(leftRight, upDown, newShit)
		}
	}

	// Save result
	myType := filepath.Ext(imagePath)
	toWhere := "duotone_" + imagePath
	out, ew := os.Create(toWhere)
	if ew != nil {
		log.Fatal(ew)
	}
	defer out.Close()

	switch myType {
	case ".png":
		ew = png.Encode(out, newStuff)
	default:
		ew = jpeg.Encode(out, newStuff, &jpeg.Options{Quality: 95})
	}

	fmt.Println("Converted image:", toWhere)
	if ew != nil {
		log.Fatal(ew)
	}
}
