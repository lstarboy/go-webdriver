package excutor

type Request interface {
}

type Response interface {
}

type BaseRequest struct {
}

type BaseResponse struct {
}

type ErrorResponse struct {
	Value Error `json:"value"`
}

type Error struct {
	Err        string `json:"error"`
	Message    string `json:"message"`
	Stacktrace string `json:"stacktrace"`
}

func (e Error) Error() string {
	return e.Err
}

type ChromeOptions struct {
	UserDataDir string
	IsHeadless  bool
}

type Options struct {
	Args []string `json:"args"`
}

type Capablities struct {
	BrowserName string  `json:"browserName"` //": "chrome"}
	Options     Options `json:"goog:chromeOptions"`
}

type Element struct {
	ElementId string `json:"ELEMENT"`
}
