package main

import (
	"os"
	"io/ioutil"

	"golang.org/x/image/font"
	"github.com/golang/freetype/truetype"
)

func loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size: size,
		GlyphCacheEntries: 1,
	}), nil
}
