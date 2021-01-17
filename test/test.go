package test

import (
	"go-webdriver/proccessor"
	"go-webdriver/proccessor/pipline"
)

func start() {
	data := []pipline.Pipline{}
	_ = proccessor.StartCase(data)
}
