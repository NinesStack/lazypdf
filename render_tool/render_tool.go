package main

import (
	"context"
	"image/png"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/NinesStack/lazypdf"
)

type rasterType int

const (
	RasterPNG rasterType = iota
	RasterSVG
)

var (
	pdf   = kingpin.Arg("pdf", "PDF file").Required().String()
	page  = kingpin.Flag("page", "page").Default("1").Short('p').Int()
	size  = kingpin.Flag("size", "size").Default("0").Short('s').Int()
	scale = kingpin.Flag("scale", "scale").Default("1.5").Short('S').Float()
	out   = kingpin.Flag("out", "out").Default("").Short('o').String()
	ext   = kingpin.Flag("ext", "ext").Default("png").Short('e').String()
)

func getRasterType(ext string) rasterType {
	switch ext {
	case "png":
		return RasterPNG
	case "svg":
		return RasterSVG
	}

	return RasterPNG
}

func outputWriter(writer func(f *os.File)) {
	f, err := os.Create(*out)
	if err != nil {
		log.Fatalf("IO error: %s", err)
	}
	defer f.Close()

	writer(f)
}

func main() {
	kingpin.Parse()

	if *out == "" {
		*out = *pdf + "." + *ext
	}

	// We will rasterize one page only
	raster := lazypdf.NewRasterizer(*pdf, 1)
	err := raster.Run()
	if err != nil {
		log.Fatalf("Failed to initialize the rasteriser: %s", err)
	}

	switch getRasterType(*ext) {
	case RasterPNG:
		img, err := raster.GeneratePageImage(context.Background(), *page, *size, *scale)
		if err != nil {
			log.Fatalf("Raster error: %s", err)
		}

		outputWriter(func(f *os.File) {
			err = png.Encode(f, img)
			if err != nil {
				log.Fatalf("Failed to write image to file: %s", err)
			}
		})

	case RasterSVG:
		svgData, err := raster.GeneratePageSVG(context.Background(), *page, *size, *scale)
		if err != nil {
			log.Fatalf("Raster error: %s", err)
		}

		outputWriter(func(f *os.File) {
			_, err = f.Write(svgData)
			if err != nil {
				log.Fatalf("Failed to write SVG to file: %s", err)
			}
		})
	default:
		log.Fatalf("Can't raster to: %s", *ext)
	}
}
