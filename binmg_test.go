package binmg

import (
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

func TestBinmg(t *testing.T) {
	// Load file
	file, err := os.Open("binmg_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Process
	out, err := Convert(file)
	if err != nil {
		t.Fatal(err)
	}

	// Save
	outFile, err := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	defer outFile.Close()
	err = png.Encode(outFile, out)
	if err != nil {
		t.Fatal(err)
	}

	// Read output
	output, err := os.Open("out.png")
	if err != nil {
		t.Fatal(err)
	}
	img, err := png.Decode(output)
	if err != nil {
		t.Fatal(err)
	}

	// Deconvert
	in, err := Deconvert(img)
	if err != nil {
		t.Fatal(err)
	}

	// Compare
	input, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	processed, err := ioutil.ReadAll(in)
	if err != nil {
		t.Fatal(err)
	}

	success := true
	if len(input) != len(processed) {
		success = false
	}
	for i, val := range input {
		if processed[i] != val {
			success = false
		}
	}

	if !success {
		t.Fatal("output isn't same as input")
	}
}
