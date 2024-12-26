package main

import (
	"image"
	"image/color"
	"image/draw"
	"sync"
)

var blurMatrix = [3][3]float64{
	{0.0625, 0.125, 0.0625},
	{0.125, 0.25, 0.125},
	{0.0625, 0.125, 0.0625},
}

func applyGaussian(img image.Image, x, y int) color.NRGBA {
	var rSum, gSum, bSum, sum float64
	bounds := img.Bounds()

	for ky := -1; ky <= 1; ky++ {
		ny := y + ky
		if ny < bounds.Min.Y || ny >= bounds.Max.Y {
			continue
		}
		for kx := -1; kx <= 1; kx++ {
			nx := x + kx
			if nx < bounds.Min.X || nx >= bounds.Max.X {
				continue
			}
			pixel := color.NRGBAModel.Convert(img.At(nx, ny)).(color.NRGBA)
			kernelWeight := blurMatrix[ky+1][kx+1]

			rSum += float64(pixel.R) * kernelWeight
			gSum += float64(pixel.G) * kernelWeight
			bSum += float64(pixel.B) * kernelWeight
			sum += kernelWeight
		}
	}

	return color.NRGBA{
		R: uint8(rSum / sum),
		G: uint8(gSum / sum),
		B: uint8(bSum / sum),
		A: 255,
	}
}

func processRowWithChannel(img image.Image, y int, rowChan chan<- *image.RGBA, wg *sync.WaitGroup) {
	defer wg.Done()

	bounds := img.Bounds()
	rowImg := image.NewRGBA(image.Rect(bounds.Min.X, y, bounds.Max.X, y+1))

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		newColor := applyGaussian(img, x, y)
		rowImg.Set(x, y, newColor)
	}

	rowChan <- rowImg
}

func filterWithChannels(img image.Image, rowResults chan *image.RGBA) draw.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)
	var wg sync.WaitGroup

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		wg.Add(1)
		go processRowWithChannel(img, y, rowResults, &wg)
	}

	go func() {
		wg.Wait()
		close(rowResults)
	}()

	for row := range rowResults {
		draw.Draw(dst, row.Bounds(), row, image.Point{0, row.Rect.Min.Y}, draw.Src)
	}

	return dst
}
