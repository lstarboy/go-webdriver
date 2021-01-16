package pipline

type Pipline struct {
	sessionId string // 会话ID

	Key int `json:"key"` // pip 的顺序

	Data PipData `json:"data"` // 管道数据
}

func CreatePipline(key int, data PipData) *Pipline {
	pip := &Pipline{}
	pip.Data = data
	pip.Key = key
	return pip
}

func (p Pipline) GetSessionId() string {
	return p.sessionId
}

func (p *Pipline) SetSessionId(id string) {
	p.sessionId = id
}

func (p Pipline) Start() {
	for _, ac := range p.Data.Actions {
		ac.WithSessionId(p.sessionId).Run()
	}
}
