package proccessor

import (
	"github.com/zhouzhe1157/go-webdriver/driver"
	"github.com/zhouzhe1157/go-webdriver/proccessor/pipline"
)

// 开启case
func StartCase(pips []pipline.Pipline) error {

	// 获取session
	resp, err := driver.GetSession()
	if err != nil {
		return err
	}

	sessionId := resp.SessionId
	for _, pip := range pips {
		// 写入sessionid
		pip.SetSessionId(sessionId)
		// 启用管道
		go pip.Start()
	}
	// 等待管道切换
	return nil
}
