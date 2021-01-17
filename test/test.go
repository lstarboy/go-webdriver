package test

import (
	"zhouzhe1157/go-webdriver/proccessor"
	"zhouzhe1157/go-webdriver/proccessor/pipline"
)

func start() {
	data := []pipline.Pipline{}
	_ = proccessor.StartCase(data)
}
