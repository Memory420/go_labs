package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"
)

func main() {
	file, err := os.Open("sok.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	start := time.Now()

	rowResults := make(chan *image.RGBA)

	blurredImg := filterWithChannels(img, rowResults)

	fmt.Println("Время обработки:", time.Since(start))

	outFile, err := os.Create("sok_blurred.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, blurredImg)
	if err != nil {
		panic(err)
	}

	fmt.Println("Сохранёно в 'sok_blurred.png'.")
}
