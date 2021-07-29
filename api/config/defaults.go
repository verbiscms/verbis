// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import "github.com/verbiscms/verbis/api/domain"

var (
	// DefaultTheme is Global configuration for the theme,
	// sets defaults to ensure that there are no
	// empty values within the themes Config
	// to prevent any errors.
	DefaultTheme = domain.ThemeConfig{
		Theme:         domain.Theme{},
		Resources:     nil,
		AssetsPath:    "assets",
		FileExtension: ".cms",
		TemplateDir:   "templates",
		LayoutDir:     "layouts",
		Media: domain.MediaConfig{
			UploadPath: "/uploads",
			AllowedFileTypes: []string{
				"image/png",
				"image/jpeg",
				"image/gif",
				"image/webp",
				"image/bmp",
				"image/svg+xml",
				"video/mpeg",
				"video/mp4",
				"video/webm",
				"application/pdf",
				"application/msword",
			},
		},
		Editor: domain.Editor{
			Modules: []string{
				"blockquote",
				"code_block",
				"code_block_highlight",
				"hardbreak",
				"h1",
				"h2",
				"h3",
				"h4",
				"h5",
				"h6",
				"paragraph",
				"hr",
				"ul",
				"ol",
				"bold",
				"code",
				"italic",
				"link",
				"strike",
				"underline",
				"history",
				"search",
				"trailing_node",
				"color",
			},
			Options: map[string]interface{}{
				// TODO: Add Verbis colours for brand.
				"palette": []string{
					"#4D4D4D", "#999999", "#FFFFFF", "#F44E3B", "#FE9200", "#FCDC00",
					"#DBDF00", "#A4DD00", "#68CCCA", "#73D8FF", "#AEA1FF", "#FDA1FF",
					"#333333", "#808080", "#CCCCCC", "#D33115", "#E27300", "#FCC400",
					"#B0BC00", "#68BC00", "#16A5A5", "#009CE0", "#7B64FF", "#FA28FF",
					"#000000", "#666666", "#B3B3B3", "#9F0500", "#C45100", "#FB9E00",
					"#808900", "#194D33", "#0C797D", "#0062B1", "#653294", "#AB149E",
				},
			},
		},
	}
)
