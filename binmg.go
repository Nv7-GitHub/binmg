package binmg

import (
	"bytes"
	"fmt"
	"image"
	"io"
)

func Convert(file io.Reader) (*image.RGBA, error) {
	fmt.Println("binmg")
	return nil, nil
}

func Deconvert(img image.Image) (io.Reader, error) {
	fmt.Println("decode")
	return bytes.NewBufferString("yeee"), nil
}
