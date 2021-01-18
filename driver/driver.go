package driver

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"zhouzhe1157/go-webdriver/excutor"
	"zhouzhe1157/go-webdriver/util"
)

type CommandRequest struct {
	SessionId string
	ElementId string
	Command   excutor.Command
	Request   excutor.Request
	Response  excutor.Response
}

func send(req CommandRequest) (string, error) {
	exec := excutor.CreateExcutor()
	exec.WithCommand(req.Command)
	if req.SessionId != "" {
		exec.WithSessionId(req.SessionId)
	}
	if req.ElementId != "" {
		exec.WithElementId(req.ElementId)
	}
	body, status, err := exec.Excute("http://127.0.0.1:9515", req.Request)
	if err != nil {
		return "", err
	}
	if body == "" {
		if status != http.StatusOK {
			return "", errors.New(fmt.Sprintf("请求异常，请求状态：%d", status))
		}
	} else {
		errsp := excutor.ErrorResponse{}
		_ = json.Unmarshal([]byte(body), &errsp)
		if errsp.Value.Err != "" {
			return "", errsp.Value
		}
	}
	return body, nil
}

func GetSession(opt excutor.ChromeOptions) (excutor.InitSessionResponse, error) {
	opts := []string{}
	if opt.UserDataDir != "" {
		//userDataDir := util.RandString(16)
		opts = append(opts, fmt.Sprintf("--user-data-dir=%s", opt.UserDataDir))
	}
	if opt.IsHeadless {
		opts = append(opts, fmt.Sprintf("--headless"))
	}
	req := CommandRequest{
		Command: excutor.NewSession,
		Request: excutor.InitSessionRequest{
			Capabilities: excutor.Capablities{
				BrowserName: "chrome",
				Options:     excutor.Options{Args: opts},
			},
		},
	}
	body, err := send(req)
	resp := excutor.InitSessionResponse{}
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func DeleteSession(sess_id string) (excutor.DeleteSessionResponse, error) {
	req := CommandRequest{
		SessionId: sess_id,
		Command:   excutor.DeleteSession,
	}
	body, err := send(req)
	resp := excutor.DeleteSessionResponse{}
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func NavigateToUrl(sess_id string, url string) (excutor.NavigateToResponse, error) {
	req := CommandRequest{
		SessionId: sess_id,
		Command:   excutor.NavigateTo,
		Request:   excutor.NavigateToRequest{Url: url},
	}
	body, err := send(req)
	resp := excutor.NavigateToResponse{}
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func GetTitle(sess_id string) (excutor.GetTitleResponse, error) {
	req := CommandRequest{
		SessionId: sess_id,
		Command:   excutor.GetTitle,
	}
	body, err := send(req)
	resp := excutor.GetTitleResponse{}
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// 切换窗口
func SwitchToWindow(sess_id, hand_id string) (excutor.SwitchToWindowResponse, error) {
	req := CommandRequest{
		SessionId: sess_id,
		Command:   excutor.SwitchToWindow,
		Request:   excutor.SwitchToWindowRequest{Name: hand_id},
	}
	body, err := send(req)
	resp := excutor.SwitchToWindowResponse{}
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// 获取所有的页面
func GetWindowHandles(sess_id string) (excutor.GetWindowHandlesResponse, error) {
	req := CommandRequest{
		SessionId: sess_id,
		Command:   excutor.GetWindowHandles,
	}
	body, err := send(req)
	resp := excutor.GetWindowHandlesResponse{}
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// 新开窗口
func NewWindow(sess_id string) (excutor.NewWindowResponse, error) {
	req := CommandRequest{
		SessionId: sess_id,
		Command:   excutor.NewWindow,
	}
	body, err := send(req)
	resp := excutor.NewWindowResponse{}
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func FindElement(sess_id string, selector excutor.Selector) (excutor.FindElementResponse, error) {
	req := CommandRequest{
		SessionId: sess_id,
		Command:   excutor.FindElement,
		Request:   excutor.FindElementRequest{Selector: selector},
	}
	body, err := send(req)
	resp := excutor.FindElementResponse{}
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func FindElements(sess_id string, selector excutor.Selector) (excutor.FindElementsResponse, error) {
	req := CommandRequest{
		SessionId: sess_id,
		Command:   excutor.FindElements,
		Request:   excutor.FindElementRequest{Selector: selector},
	}
	body, err := send(req)
	fmt.Println(body)
	resp := excutor.FindElementsResponse{}

	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func FindElementFromElement(sess_id, element_id string, selector excutor.Selector) (excutor.FindElementResponse, error) {
	req := CommandRequest{
		SessionId: sess_id,
		ElementId: element_id,
		Command:   excutor.FindElementFromElement,
		Request:   excutor.FindElementRequest{Selector: selector},
	}
	body, err := send(req)
	fmt.Println(body)
	resp := excutor.FindElementResponse{}

	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func FindElementsFromElement(sess_id, element_id string, selector excutor.Selector) (excutor.FindElementsResponse, error) {
	req := CommandRequest{
		SessionId: sess_id,
		ElementId: element_id,
		Command:   excutor.FindElementsFromElement,
		Request:   excutor.FindElementRequest{Selector: selector},
	}
	body, err := send(req)
	resp := excutor.FindElementsResponse{}

	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func ElementClick(sess_id, element_id string) (excutor.ElementClickResponse, error) {
	req := CommandRequest{
		SessionId: sess_id,
		ElementId: element_id,
		Command:   excutor.ElementClick,
		Request:   excutor.ElementClickRequest{},
	}
	body, err := send(req)
	resp := excutor.ElementClickResponse{}

	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func ElementSendKeys(sess_id, element_id, value string) (excutor.ElementSendKeysResponse, error) {
	req := CommandRequest{
		SessionId: sess_id,
		ElementId: element_id,
		Command:   excutor.ElementSendKeys,
		Request:   excutor.ElementSendKeysRequest{Text: value, Value: []string{value}},
	}
	body, err := send(req)
	resp := excutor.ElementSendKeysResponse{}

	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func GetElementText(sess_id, element_id string) (excutor.GetElementTextResponse, error) {
	req := CommandRequest{
		SessionId: sess_id,
		ElementId: element_id,
		Command:   excutor.GetElementText,
		Request:   excutor.GetElementTextRequest{},
	}
	body, err := send(req)
	resp := excutor.GetElementTextResponse{}

	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func TakeScreenshot(sess_id string, fileName string) (excutor.TakeScreenshotResponse, error) {
	req := CommandRequest{
		SessionId: sess_id,
		Command:   excutor.TakeScreenshot,
		Request:   excutor.TakeScreenshotRequest{},
	}
	body, err := send(req)
	resp := excutor.TakeScreenshotResponse{}

	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}

	util.SaveImage(fileName, resp.Value)

	return resp, nil
}

func GetTimeouts(sess_id string) {
	req := CommandRequest{
		SessionId: sess_id,
		Command:   excutor.TakeScreenshot,
		Request:   excutor.TakeScreenshotRequest{},
	}
	body, err := send(req)
	resp := excutor.TakeScreenshotResponse{}

	if err != nil {
		//return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		//return resp, err
	}
	//util.SaveImage(fileName, resp.Value)
	//return resp, nil
}

func ExcuteScript(sess_id string, script string) (excutor.ExecuteScriptResponse, error) {
	req := CommandRequest{
		SessionId: sess_id,
		Command:   excutor.ExecuteScript,
		Request: excutor.ExecuteScriptRequest{
			Args:   []string{},
			Script: script,
		},
	}
	body, err := send(req)
	resp := excutor.ExecuteScriptResponse{}
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
