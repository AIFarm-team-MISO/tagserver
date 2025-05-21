package main

import (
	"embed"
	"tagserver/backend" // 너의 모듈 경로에 따라 조정

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist/*
var assets embed.FS

func main() {
	app := backend.NewCrawler()

	wails.Run(&options.App{
		Title:  "TagServer",
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []interface{}{app}, // ✅ 이 줄 반드시 있어야 함!
	})
}
