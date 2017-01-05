package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/png"
	"os"

	"github.com/ftloc/exception"
)

var (
	inputfile = flag.String("input", "", "input file (png)")
	x         = flag.Int("x", 0, "left start position")
	y         = flag.Int("y", 0, "top start position")
	w         = flag.Int("w", 0, "width")
	h         = flag.Int("h", 0, "height")
)

func main() {
	flag.Parse()

	if *inputfile == "" {
		flag.Usage()
		return
	}

	f, err := os.Open(*inputfile)
	exception.ThrowOnError(err, err)

	i, err := png.Decode(f)
	exception.ThrowOnError(err, err)

	out := make([]uint64, *w)

	for yi := 0; yi < *h; yi++ {
		for xi := 0; xi < *w; xi++ {
			c := i.At(xi+*x, yi+*y)
			if g, ok := c.(color.Gray); ok {
				out[xi] += uint64(g.Y >> 8)
			} else {
				r, _, _, _ := i.At(xi+*x, yi+*y).RGBA()
				out[xi] += uint64(r)
			}
		}
	}

	for e := range out {
		fmt.Printf("%d ", (out[e] >> 8))
	}

	fmt.Println()
}
