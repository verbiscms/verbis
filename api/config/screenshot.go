// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"os"
)

const (
	// The default screenshot name within the theme's
	// directory.
	ScreenshotName = "screenshot"
)

var (
	ScrenshotExtensions = []string{
		".png",
		".svg",
		".jpg",
	}
)

func FindScreenshot(path, theme string) (string, error) {
	for _, v := range ScrenshotExtensions {
		name := path + string(os.PathSeparator) + theme + string(os.PathSeparator) + ScreenshotName + v
		fmt.Println(name)
		info, err := os.Stat(name)
		if err != nil {
			continue
		}
		fmt.Println(info.Name())
		return "test", nil
	}
	return "", fmt.Errorf("no screenshot found")
}
