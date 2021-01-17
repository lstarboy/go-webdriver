package pipline

import "zhouzhe1157/go-webdriver/proccessor/action"

type pip interface {

	// 获取会话ID
	GetSessionId() string

	// 启动
	Start()
}

// 管道数据
type PipData struct {
	Actions []action.Action `json:"actions"` // 操作列表
}
