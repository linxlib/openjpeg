package main

import (
	"fmt"
	"github.com/chai2010/webp"
	"github.com/linxlib/logs"
	openjpeg "github.com/linxlib/openjpeg"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"image"

	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
)

type encFun func(w io.Writer, i image.Image) error

const rfmts = "[jpg|jp2|png|webp|tiff|bmp|gif]"
const wfmts = "[jpg|jp2|png|webp|tiff|bmp|gif]"

func main() {

	encTab := map[string]encFun{
		".png":  png.Encode,
		".bmp":  bmp.Encode,
		".jp2":  func(w io.Writer, i image.Image) error { return openjpeg.Encode(w.(io.WriteSeeker), i, nil) },
		".gif":  func(w io.Writer, i image.Image) error { return gif.Encode(w, i, nil) },
		".jpg":  func(w io.Writer, i image.Image) error { return jpeg.Encode(w, i, nil) },
		".tiff": func(w io.Writer, i image.Image) error { return tiff.Encode(w, i, nil) },
		".webp": func(w io.Writer, i image.Image) error { return webp.Encode(w, i, nil) },
	}
	in, err := os.Open("jp2.jp2")
	if err != nil {
		logs.Error(err)
		fmt.Scanln()
		return
	}
	img, inFmt, err := image.Decode(in)
	logs.Printf("Decoded %s: %dx%d %s\n", in.Name(), img.Bounds().Dx(), img.Bounds().Dy(), inFmt)
	_ = in.Close()
	out, err := os.Create("jp2000.jpg")
	if err != nil {
		logs.Error(err)
		fmt.Scanln()
		return
	}
	ofmt := filepath.Ext(out.Name())
	log.Printf("Encoding into %s\n", out.Name())
	if encFun, ok := encTab[ofmt]; ok {
		err := encFun(out, img)
		if err != nil {
			logs.Error(err)
			fmt.Scanln()
			return
		}
		log.Printf("Done")
	} else {
		log.Printf("Unknown output format %s\n", ofmt)
	}
	fmt.Scanln()
	return
}
