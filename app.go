package main

import (
	"context"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	// 앱 시작 시 호출됨
}

func (a *App) Greet(name string) string {
	return "Hello, " + name
}
