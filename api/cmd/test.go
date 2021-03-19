// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/nickalie/go-webpbin"
	"github.com/spf13/cobra"
	"image"
	"image/color"
	"os"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {

			p := paths.Get()
			webpbin.Dest(p.Bin + string(os.PathSeparator) + "webp")

			width := 200
			height := 100

			img := image.NewRGBA(image.Rectangle{
				Min: image.Point{},
				Max: image.Point{X: width, Y: height},
			})

			// Set color for each pixel.
			for x := 0; x < width; x++ {
				for y := 0; y < height; y++ {
					switch {
					case x < width && y < height:
						img.Set(x, y, color.RGBA{R: 100, G: 200, B: 200, A: 0xff})
					}
				}
			}

			b := &bytes.Buffer{}
			err := webpbin.Encode(b, img)

			fmt.Println(err)
		},
	}
)
