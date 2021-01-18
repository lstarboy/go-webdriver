package test

import (
	"zhouzhe1157/go-webdriver/proccessor"
	"zhouzhe1157/go-webdriver/proccessor/pipline"
	"zhouzhe1157/go-webdriver/util"
)

func start() {
	data := []pipline.Pipline{}
	userDataDir := "E:\\logs\\" + util.RandString(16)
	_ = proccessor.StartCase(data, userDataDir)
}
