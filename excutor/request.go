package excutor

type InitSessionRequest struct {
	BaseRequest
	Capabilities Capablities `json:"desiredCapabilities"`
}

type NavigateToRequest struct {
	BaseRequest
	Url string `json:"url"`
}

type SwitchToWindowRequest struct {
	BaseRequest
	Name string `json:"name"`
}

type FindElementRequest struct {
	BaseRequest
	Selector
}

type ElementClickRequest struct {
	BaseRequest
}

type ElementSendKeysRequest struct {
	BaseRequest
	Text  string   `json:"text"`
	Value []string `json:"value"`
}

type GetElementTextRequest struct {
	BaseRequest
}

type TakeScreenshotRequest struct {
	BaseRequest
}

type ExecuteScriptRequest struct {
	BaseRequest
	Script string   `json:"script"`
	Args   []string `json:"args"`
}
