package ImageToASCII

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"bytes"
	"os"
	"errors"
	"fmt"
)

type ImageASCII struct {
	Img   image.Image
	Chars string
}

func (this *ImageASCII) Width() int {
	return this.Img.Bounds().Size().X
}

func (this *ImageASCII) Height() int {
	return this.Img.Bounds().Size().Y
}

func (this *ImageASCII) getBrightness(point image.Point) int {
	R, G, B, _ := this.Img.At(point.X, point.Y).RGBA()
	var my float64 = (float64(R) + float64(G) + float64(B))/3
	defer func() {
		recover()
	}()
	return int(float64(len(this.Chars)) / (65536 / my))
}

func (this *ImageASCII) getBrightnessWH(point image.Point, width, height int) int {
	var sum float64 = 0
	defer func() {
		recover()
	}()
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			R, G, B, _ := this.Img.At(point.X + w, point.Y + h).RGBA()
			sum += (float64(R) + float64(G) + float64(B))/3
		}
	}
	var my float64 = sum / (float64(height*width))
	return int(float64(len(this.Chars)) / (65536 / my))
}

func (this *ImageASCII) process(w, h int) (*bytes.Buffer, error) {
	if this.Width() < w || this.Height() < h {
		return nil, errors.New("new width and height must be less or equal to image size")
	}
	var buf bytes.Buffer

	if w == 0 || h == 0 {
		for i := 0; i < this.Height(); i++ {
			for j := 0; j < this.Width(); j++ {
				buf.WriteByte(this.Chars[this.getBrightness(image.Point{j, i})])
			}
			buf.WriteRune('\n')
		}
		return &buf, nil
	}

	var block_w, block_h int = this.Width()/w, this.Height()/h
	for i := 0; i < this.Height(); i += block_h+1 {
		for j := 0; j < this.Width(); j+=block_w+1 {
			buf.WriteByte(this.Chars[this.getBrightnessWH(image.Point{j, i}, block_w, block_h)])
		}
		buf.WriteRune('\n')
	}

	return &buf, nil
}

func (this *ImageASCII) ToStringWH(width, height int) string {
	var buf *bytes.Buffer
	buf, _ = this.process(width,height)
	return buf.String()
}

func (this *ImageASCII) ToString() string {
	return this.ToStringWH(0,0)
}

func (this *ImageASCII) ToWriterWH(w *io.Writer, width, height int) {
	fmt.Fprint(*w, this.ToStringWH(width, height))
}

func (this *ImageASCII) ToWriter(w *io.Writer) {
	fmt.Fprint(*w, this.ToString())
}

func LoadFromImage(img image.Image) (*ImageASCII, error) {
	var i ImageASCII
	i.Img = img
	return &i, nil
}

func LoadFromStream(r io.Reader) (*ImageASCII, error) {
	var img ImageASCII
	var err error

	img.Img, _, err = image.Decode(r)
	if err != nil {
		return nil, err
	}

	img.Chars = ".,:;ox%#@"

	return &img, nil
}

func LoadFromFile(file_name string) (*ImageASCII, error) {
	file, err := os.Open(file_name)
	if err != nil {
		return nil, err
	}
	return LoadFromStream(file)
}