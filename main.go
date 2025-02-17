package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "wails-app",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
			// Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//     // 排除 API 路由
			//     if !strings.HasPrefix(r.URL.Path, "/api") {
			//         // 重置路径为前端入口
			//         r.URL.Path = "/"
			//     }
			//     // 调用默认处理器
			//     assetserver.DefaultAssets.ServeHTTP(w, r)
			// }),
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
