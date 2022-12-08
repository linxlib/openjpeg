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

//func Parse(f string) (image.Image, error) {
//	logs.Infof("write jp2 file %s to png", f)
//	gmagick.Initialize()
//	defer gmagick.Terminate()
//
//	mw := gmagick.NewMagickWand()
//	defer mw.Destroy()
//	if err := mw.ReadImage(f); err != nil {
//		return nil, err
//	}
//
//	err := mw.SetImageFormat("png")
//	if err != nil {
//		return nil, err
//	}
//	return png.Decode(bytes.NewReader(mw.WriteImageBlob()))
//}
//
//func Write(str string) {
//	logs.Infof("write png file to jpeg2000 file")
//	gmagick.Initialize()
//	defer gmagick.Terminate()
//
//	mw := gmagick.NewMagickWand()
//	defer mw.Destroy()
//	err := mw.ReadImage(str)
//	if err != nil {
//		logs.Error(err)
//		return
//	}
//	err = mw.SetImageFormat("png")
//	if err != nil {
//		logs.Error(err)
//		return
//	}
//	err = mw.SetImageCompression(gmagick.COMPRESSION_JPEG2000)
//	if err != nil {
//		logs.Error(err)
//		return
//	}
//	err = mw.WriteImage("jp2000.jp2")
//	if err != nil {
//		logs.Error(err)
//		return
//	}
//}
//func main() {
//	i, err := Parse("jp2.jp2")
//	if err != nil {
//		logs.Error(err)
//		fmt.Scanln()
//		return
//	}
//	j, err := os.Create("1.png")
//	if err != nil {
//		logs.Error(err)
//		fmt.Scanln()
//		return
//	}
//	err = png.Encode(j, i)
//	if err != nil {
//		logs.Error(err)
//		fmt.Scanln()
//		return
//	}
//	Write("1.png")
//	fmt.Scanln()
//}

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
