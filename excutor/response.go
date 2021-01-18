package excutor

type InitSessionResponse struct {
	BaseResponse
	SessionId string `json:"sessionId"`
	Value     struct {
		Capabilities Capablities `json:"capabilities"`
		SessionId    string      `json:"sessionId"`
	} `json:"value"`
}

type DeleteSessionResponse struct {
	BaseResponse
}

type NavigateToResponse struct {
	BaseResponse
}

type GetTitleResponse struct {
	BaseResponse
	SessionId string `json:"sessionId"`
	Status    int    `json:"status"`
	Value     string `json:"value"`
}

type GetWindowHandlesResponse struct {
	BaseResponse
	SessionId string   `json:"sessionId"`
	Status    int      `json:"status"`
	Value     []string `json:"value"`
}

type SwitchToWindowResponse struct {
	BaseResponse
	SessionId string `json:"sessionId"`
	Status    int    `json:"status"`
}

type NewWindowResponse struct {
	BaseResponse
}

type FindElementResponse struct {
	BaseResponse
	Value Element `json:"value"`
}

type FindElementsResponse struct {
	BaseResponse
	Value []Element `json:"value"`
}

type ElementClickResponse struct {
	BaseResponse
	Value interface{} `json:"value"`
}

type ElementSendKeysResponse struct {
	BaseResponse
	Value interface{} `json:"value"`
}

type GetElementTextResponse struct {
	BaseResponse
	Value interface{} `json:"value"`
}

type TakeScreenshotResponse struct {
	BaseResponse
	Value string `json:"value"`
}

type ExecuteScriptResponse struct {
	BaseResponse
	Value  interface{} `json:"value"`
	Status int         `json:"status"`
}
