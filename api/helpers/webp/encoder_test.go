package webp

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"golang.org/x/image/webp"
	"image/jpeg"
	"os"
	"testing"
)

func TestEncode(t *testing.T) {
	f, err := os.Open("source.jpg")
	assert.Nil(t, err)
	imgSource, err := jpeg.Decode(f)
	assert.Nil(t, err)
	var b bytes.Buffer
	err = Encode(&b, imgSource)
	assert.Nil(t, err)
	imgTarget, err := webp.Decode(bytes.NewReader(b.Bytes()))
	assert.Nil(t, err)
	assert.Equal(t, imgSource.Bounds(), imgTarget.Bounds())
}
