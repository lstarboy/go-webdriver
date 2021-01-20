package pipline

const (
	SuccessUtil = 1
)

type Pipline struct {
	sessionId string // 会话ID

	Key int `json:"key"` // pip 的顺序

	Data PipData `json:"data"` // 管道数据

	IsSuccessUntil int `json:"is_success_until"`
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

func (p *Pipline) SetSessionId(id string) *Pipline {
	p.sessionId = id
	return p
}

func (p Pipline) Start() (err error) {
	if p.IsSuccessUntil == SuccessUtil {
		for {
			if err = p.dispatch(); err == nil {
				break
			}
		}
	} else {
		err = p.dispatch()
	}
	return
}

func (p Pipline) dispatch() error {
	for _, ac := range p.Data.Actions {
		err := ac.WithSessionId(p.sessionId).Run()
		if err != nil {
			return err
		}
	}
	return nil
}
