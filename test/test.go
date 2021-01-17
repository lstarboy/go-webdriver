package test

import (
	"github.com/zhouzhe1157/go-webdriver/proccessor"
	"github.com/zhouzhe1157/go-webdriver/proccessor/pipline"
)

func start() {
	data := []pipline.Pipline{}
	_ = proccessor.StartCase(data)
}
