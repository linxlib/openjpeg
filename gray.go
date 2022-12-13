package openjpeg_go

import (
	"image"
	"image/color"
)

type DicomImage struct {
	level  int16
	width  int16
	data   []int32
	bounds image.Rectangle
	stride int
}

func (p *DicomImage) WindowLevel() int16 {
	return p.level
}

func (p *DicomImage) WindowWidth() int16 {
	return p.width
}
func (p *DicomImage) SetWindowLevel(wl int16) {
	p.level = wl
}

func (p *DicomImage) SetWindowWidth(ww int16) {
	p.width = ww
}

func (p *DicomImage) ColorModel() color.Model {
	return color.Gray16Model
}

func (p *DicomImage) Bounds() image.Rectangle {
	return p.bounds
}

func (p *DicomImage) At(x, y int) color.Color {
	if !(image.Point{X: x, Y: y}.In(p.bounds)) {
		return color.Gray16{}
	}

	windowMin := p.level - p.width/2
	windowMax := windowMin + p.width

	index := p.PixOffset(x, y)
	if index >= len(p.data) {
		return color.Black
	}
	raw := int16(p.data[index])
	if raw < windowMin {
		return color.Gray{}
	} else if raw >= windowMax {
		return color.Gray{Y: 0xff}
	}
	val := float32(raw-windowMin) / float32(p.width)

	return color.Gray{Y: uint8(float32(0xff) * val)}
}

func (p *DicomImage) PixOffset(x, y int) int {
	return (y-p.bounds.Min.Y)*p.stride + (x-p.bounds.Min.X)*1
}
