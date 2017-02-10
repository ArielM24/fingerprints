package main

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
	"github.com/nfnt/resize"
	_ "github.com/nfnt/resize"
)

func loadImage(name string) *image.Gray {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	m = resize.Resize(450, 450, m, resize.Bilinear)

	bounds := m.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	gray := image.NewGray(bounds)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := m.At(x, y)
			grayColor := color.GrayModel.Convert(oldColor)
			gray.Set(x, y, grayColor)
		}
	}

	return gray
}

func appMain(driver gxui.Driver) {
	imgSrc := processImage()
	widthMax := imgSrc.Bounds().Max

	theme := flags.CreateTheme(driver)
	window := theme.CreateWindow(widthMax.X, widthMax.Y, "Image viewer")
	window.SetScale(flags.DefaultScaleFactor)
	window.SetBackgroundBrush(gxui.WhiteBrush)

	img := theme.CreateImage()
	window.AddChild(img)

	gray := image.NewRGBA(imgSrc.Bounds())
	draw.Draw(gray, imgSrc.Bounds(), imgSrc, image.ZP, draw.Src)
	texture := driver.CreateTexture(gray, 1)
	img.SetTexture(texture)

	window.OnClose(driver.Terminate)
}

func processImage() *image.Gray {
	img := loadImage("corpus/nist2.jpg")
	normalizeGray(img)
	img = ComputeDirectional(img)
	return img
}

func main() {
	gl.StartDriver(appMain)
}
