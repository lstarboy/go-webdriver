package proccessor

import (
	"zhouzhe1157/go-webdriver/driver"
	"zhouzhe1157/go-webdriver/excutor"
	"zhouzhe1157/go-webdriver/proccessor/pipline"
)

// 开启case
func StartCase(pips []pipline.Pipline, logPath string) error {
	opts := excutor.ChromeOptions{IsHeadless: false}
	// 获取session
	if logPath != "" {
		opts.UserDataDir = logPath
	}
	resp, err := driver.GetSession(opts)
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
