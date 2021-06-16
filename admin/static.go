package admin

import "embed"

var (
	//go:embed dist/*
	SPA embed.FS
)
