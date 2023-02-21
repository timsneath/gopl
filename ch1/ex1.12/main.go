package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{
	color.Black, // background
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF},
	color.RGBA{0x00, 0xFF, 0xFF, 0xFF},
	color.RGBA{0x00, 0x00, 0xFF, 0xFF},
	color.RGBA{0xFF, 0x00, 0xFF, 0xFF},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var cycles int

		formValue := r.FormValue("cycles")
		if formValue == "" {
			cycles = 5
		}

		cycles, err := strconv.Atoi(formValue)
		if err != nil {
			cycles = 5
		}
		lissajous(w, cycles)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	maxT := float64(cycles) * 2.0 * math.Pi
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < maxT; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				(uint8(math.Round(maxT/t))%5)+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
