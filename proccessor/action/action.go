package action

import (
	"fmt"
	"go-webdriver/driver"
	"go-webdriver/excutor"
	"go-webdriver/util"
	"time"
)

const (
	ACTION_NAVIGATETO = 10101 // 跳到指定页面

	ACTION_VIEW_VALUE = 10201 // 查看 （查看对象值，对象是否存在）

	ACTINO_SEND_KEYS = 10301 // 填写值

	ACTION_CLICK = 10401 // 页面元素点击

	ACTION_SCREENSHOT = 10501 // 页面截图

	ACTION_SWITCH_WINDOW = 10601 // 切换窗口

	ACTION_NEW_WINDOW = 10701 // 新开窗口

	EXPECT_TYPE_EXIST = 1 // 存在与否类型

	EXPECT_TYPE_VALUE = 2 // 值类型对比

	BLOCK_ACTION = 1 // 阻塞行为

	UN_BLOCK_ACTION = -1 // 非阻塞行为

	SELECTOR_CSS = 10

	SELECTOR_XPATH = 20
)

// 操作
type Action struct {

	// 操作名称
	ActionName string `json:"action_name"`

	// 操作类型
	ActionType int `json:"action_type"`

	// 操作目标
	ActionTarget string `json:"action_target"`

	// 操作选择器
	ActionSelector int `json:"action_selector"`

	// 操作值
	ActionValue string `json:"action_value"`

	// 操作延迟时间
	ActionDelay int `json:"action_delay"`

	// 预测结果类型
	ExpectType int `json:"expect_type"`

	// 预测结果
	ActionExpectValue string `json:"expect_value"`

	// 是否阻塞
	IsBlock int `json:"is_block"`

	// session_id
	session_id string
}

// add session
func (a *Action) WithSessionId(id string) *Action {
	a.session_id = id
	return a
}

func (a *Action) getSelector() excutor.Selector {
	selector := excutor.CSS_SELECTOR
	switch a.ActionSelector {
	case SELECTOR_XPATH:
		selector = excutor.XPATH_SELECTOR
		break
	}
	return excutor.CreateSelector(selector, a.ActionTarget)
}

func (a *Action) Run() {

	// 操作类型
	switch a.ActionType {

	case ACTION_NAVIGATETO:
		resp, err := driver.NavigateToUrl(a.session_id, a.ActionTarget)
		fmt.Print(resp, err)
		break
	case ACTION_CLICK:
		resp, _ := driver.FindElement(a.session_id, a.getSelector())
		respx, _ := driver.ElementClick(a.session_id, resp.Value.ElementId)
		fmt.Println(respx)
		break
	case ACTINO_SEND_KEYS:
		resp, _ := driver.FindElement(a.session_id, a.getSelector())
		respx, _ := driver.ElementSendKeys(a.session_id, resp.Value.ElementId, a.ActionValue)
		fmt.Println(respx)
		break
	case ACTION_VIEW_VALUE:
		resp, _ := driver.FindElement(a.session_id, a.getSelector())
		respx, _ := driver.GetElementText(a.session_id, resp.Value.ElementId)
		fmt.Println("value is:", respx.Value)
		break
	case ACTION_SCREENSHOT:
		resp, _ := driver.TakeScreenshot(a.session_id, "asd.png")
		fmt.Println("value is: ", resp)
		break
	case ACTION_SWITCH_WINDOW:
		resp, _ := driver.GetWindowHandles(a.session_id)
		handId := util.ToInt(a.ActionValue)
		respx, _ := driver.SwitchToWindow(a.session_id, resp.Value[handId])
		fmt.Sprintf("switch result", respx)
		break
	case ACTION_NEW_WINDOW:
		resp, _ := driver.NewWindow(a.session_id)
		fmt.Sprintf("open window:", resp)
		break
	}
	if a.ActionDelay > 0 {
		time.Sleep(time.Duration(a.ActionDelay) * time.Second)
	}

}
