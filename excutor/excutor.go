package excutor

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zhouzhe1157/go-webdriver/util"
	"strings"
	"time"
)

type Command struct {
	method      string
	path        string
	description string
	response    Response
}

var (
	NewSession              = Command{method: "POST", path: "/session", description: "NewSession"}
	DeleteSession           = Command{method: "DELETE", path: "/session/{sessionid}", description: "DeleteSession"}
	Status                  = Command{method: "GET", path: "/status", description: "Status"}
	GetTimeouts             = Command{method: "GET", path: "/session/{sessionid}/timeouts", description: "GetTimeouts"}
	SetTimeouts             = Command{method: "POST", path: "/session/{sessionid}/timeouts", description: "SetTimeouts"}
	NavigateTo              = Command{method: "POST", path: "/session/{sessionid}/url", description: "NavigateTo"}
	GetCurrentURL           = Command{method: "GET", path: "/session/{sessionid}/url", description: "GetCurrentURL"}
	Back                    = Command{method: "POST", path: "/session/{sessionid}/back", description: "Back"}
	Forward                 = Command{method: "POST", path: "/session/{sessionid}/forward", description: "Forward"}
	Refresh                 = Command{method: "POST", path: "/session/{sessionid}/refresh", description: "Refresh"}
	GetTitle                = Command{method: "GET", path: "/session/{sessionid}/title", description: "GetTitle"}
	GetWindowHandle         = Command{method: "GET", path: "/session/{sessionid}/window", description: "GetWindowHandle"}
	CloseWindow             = Command{method: "DELETE", path: "/session/{sessionid}/window", description: "CloseWindow"}
	SwitchToWindow          = Command{method: "POST", path: "/session/{sessionid}/window", description: "SwitchToWindow"}
	GetWindowHandles        = Command{method: "GET", path: "/session/{sessionid}/window/handles", description: "GetWindowHandles"}
	NewWindow               = Command{method: "POST", path: "/session/{sessionid}/window/new", description: "NewWindow"}
	SwitchToFrame           = Command{method: "POST", path: "/session/{sessionid}/frame", description: "SwitchToFrame"}
	SwitchToParentFrame     = Command{method: "POST", path: "/session/{sessionid}/frame/parent", description: "SwitchToParentFrame"}
	GetWindowRect           = Command{method: "GET", path: "/session/{sessionid}/window/rect", description: "GetWindowRect"}
	SetWindowRect           = Command{method: "POST", path: "/session/{sessionid}/window/rect", description: "SetWindowRect"}
	MaximizeWindow          = Command{method: "POST", path: "/session/{sessionid}/window/maximize", description: "MaximizeWindow"}
	MinimizeWindow          = Command{method: "POST", path: "/session/{sessionid}/window/minimize", description: "MinimizeWindow"}
	FullscreenWindow        = Command{method: "POST", path: "/session/{sessionid}/window/fullscreen", description: "FullscreenWindow"}
	GetActiveElement        = Command{method: "GET", path: "/session/{sessionid}/element/active", description: "GetActiveElement"}
	FindElement             = Command{method: "POST", path: "/session/{sessionid}/element", description: "FindElement"}
	FindElements            = Command{method: "POST", path: "/session/{sessionid}/elements", description: "FindElements"}
	FindElementFromElement  = Command{method: "POST", path: "/session/{sessionid}/element/{elementid}/element", description: "FindElementFromElement"}
	FindElementsFromElement = Command{method: "POST", path: "/session/{sessionid}/element/{elementid}/elements", description: "FindElementsFromElement"}
	IsElementSelected       = Command{method: "GET", path: "/session/{sessionid}/element/{elementid}/selected", description: "IsElementSelected"}
	GetElementAttribute     = Command{method: "GET", path: "/session/{sessionid}/element/{elementid}/attribute/{name}", description: "GetElementAttribute"}
	GetElementProperty      = Command{method: "GET", path: "/session/{sessionid}/element/{elementid}/property/{name}", description: "GetElementProperty"}
	GetElementCSSValue      = Command{method: "GET", path: "/session/{sessionid}/element/{elementid}/css/{propertyname}", description: "GetElementCSSValue"}
	GetElementText          = Command{method: "GET", path: "/session/{sessionid}/element/{elementid}/text", description: "GetElementText"}
	GetElementTagName       = Command{method: "GET", path: "/session/{sessionid}/element/{elementid}/name", description: "GetElementTagName"}
	GetElementRect          = Command{method: "GET", path: "/session/{sessionid}/element/{elementid}/rect", description: "GetElementRect"}
	IsElementEnabled        = Command{method: "GET", path: "/session/{sessionid}/element/{elementid}/enabled", description: "IsElementEnabled"}
	GetComputedRole         = Command{method: "GET", path: "/session/{sessionid}/element/{elementid}/computedrole", description: "GetComputedRole"}
	GetComputedLabel        = Command{method: "GET", path: "/session/{sessionid}/element/{elementid}/computedlabel", description: "GetComputedLabel"}
	ElementClick            = Command{method: "POST", path: "/session/{sessionid}/element/{elementid}/click", description: "ElementClick"}
	ElementClear            = Command{method: "POST", path: "/session/{sessionid}/element/{elementid}/clear", description: "ElementClear"}
	ElementSendKeys         = Command{method: "POST", path: "/session/{sessionid}/element/{elementid}/value", description: "ElementSendKeys"}
	GetPageSource           = Command{method: "GET", path: "/session/{sessionid}/source", description: "GetPageSource"}
	ExecuteScript           = Command{method: "POST", path: "/session/{sessionid}/execute/sync", description: "ExecuteScript"}
	ExecuteAsyncScript      = Command{method: "POST", path: "/session/{sessionid}/execute/async", description: "ExecuteAsyncScript"}
	GetAllCookies           = Command{method: "GET", path: "/session/{sessionid}/cookie", description: "GetAllCookies"}
	GetNamedCookie          = Command{method: "GET", path: "/session/{sessionid}/cookie/{name}", description: "GetNamedCookie"}
	AddCookie               = Command{method: "POST", path: "/session/{sessionid}/cookie", description: "AddCookie"}
	DeleteCookie            = Command{method: "DELETE", path: "/session/{sessionid}/cookie/{name}", description: "DeleteCookie"}
	DeleteAllCookies        = Command{method: "DELETE", path: "/session/{sessionid}/cookie", description: "DeleteAllCookies"}
	PerformActions          = Command{method: "POST", path: "/session/{sessionid}/actions", description: "PerformActions"}
	ReleaseActions          = Command{method: "DELETE", path: "/session/{sessionid}/actions", description: "ReleaseActions"}
	DismissAlert            = Command{method: "POST", path: "/session/{sessionid}/alert/dismiss", description: "DismissAlert"}
	AcceptAlert             = Command{method: "POST", path: "/session/{sessionid}/alert/accept", description: "AcceptAlert"}
	GetAlertText            = Command{method: "GET", path: "/session/{sessionid}/alert/text", description: "GetAlertText"}
	SendAlertText           = Command{method: "POST", path: "/session/{sessionid}/alert/text", description: "SendAlertText"}
	TakeScreenshot          = Command{method: "GET", path: "/session/{sessionid}/screenshot", description: "TakeScreenshot"}
	TakeElementScreenshot   = Command{method: "GET", path: "/session/{sessionid}/element/{elementid}/screenshot", description: "TakeElementScreenshot"}
	PrintPage               = Command{method: "POST", path: "/session/{sessionid}/print", description: "PrintPage"}
)

type Excutor struct {
	cmd          Command
	session_id   string
	name         string
	element_id   string
	propertyname string
}

func CreateExcutor() *Excutor {
	return &Excutor{}
}

func (exe *Excutor) WithCommand(cmd Command) *Excutor {
	exe.cmd = cmd
	return exe
}

func (exe *Excutor) WithSessionId(sess_id string) *Excutor {
	exe.session_id = sess_id
	return exe
}

func (exe *Excutor) WithName(name string) *Excutor {
	exe.name = name
	return exe
}

func (exe *Excutor) WithElementId(element_id string) *Excutor {
	exe.element_id = element_id
	return exe
}

func (exe *Excutor) WithPropertyname(propertyname string) *Excutor {
	exe.propertyname = propertyname
	return exe
}

func (exe *Excutor) parse() (string, error) {
	var path = exe.cmd.path
	if strings.Contains(exe.cmd.path, "{sessionid}") {
		if exe.session_id == "" {
			return "", errors.New("session_id不能为空")
		}
		path = strings.Replace(exe.cmd.path, "{sessionid}", exe.session_id, -1)
	}
	if strings.Contains(path, "{name}") {
		if exe.session_id == "" {
			return "", errors.New("name不能为空")
		}
		path = strings.Replace(path, "{name}", exe.name, -1)
	}
	if strings.Contains(path, "{propertyname}") {
		if exe.session_id == "" {
			return "", errors.New("propertyname不能为空")
		}
		path = strings.Replace(path, "{propertyname}", exe.propertyname, -1)
	}
	if strings.Contains(path, "{elementid}") {
		if exe.session_id == "" {
			return "", errors.New("elementid不能为空")
		}
		path = strings.Replace(path, "{elementid}", exe.element_id, -1)
	}
	return path, nil
}

func (exe *Excutor) Excute(domain string, req Request) (string, int, error) {
	ctx := context.Background()
	defer ctx.Done()
	var (
		body   string
		status int
		err    error
	)
	path, err := exe.parse()
	path = domain + path
	bytes := []byte{}
	if req != nil {
		bytes, _ = json.Marshal(req)
	}
	body, status, err = util.DoRequest(&ctx, path, string(bytes), exe.cmd.method, 30*time.Second, util.ExtParams{})
	return body, status, err
}
