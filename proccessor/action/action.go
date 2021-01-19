package action

import (
	"errors"
	"sync"
	"time"
	"zhouzhe1157/go-webdriver/driver"
	"zhouzhe1157/go-webdriver/excutor"
	errors2 "zhouzhe1157/go-webdriver/proccessor/errors"
	"zhouzhe1157/go-webdriver/util"
)

const (
	ACTION_NAVIGATETO = 10101 // 跳到指定页面

	ACTION_VIEW_VALUE = 10201 // 查看 （查看对象值，对象是否存在）

	ACTION_VIEW_TITLE = 10202 // 获取title

	ACTINO_SEND_KEYS = 10301 // 填写值

	ACTION_CLICK = 10401 // 页面元素点击

	ACTION_SCREENSHOT = 10501 // 页面截图

	ACTION_SWITCH_WINDOW = 10601 // 切换窗口

	ACTION_NEW_WINDOW = 10701 // 新开窗口

	ACTION_WAIT = 10801 // 等待

	ACTINO_EXCUTE_SCRIPT = 10901 // 执行脚本

	EXPECT_TYPE_EXIST = 1 // 存在与否类型

	EXPECT_TYPE_NOT_EXIST = 2 // 不存在

	EXPECT_VALUE_EQUAL = 3 // 值类型对比相等

	EXPECT_VALUE_NOT_EQUAL = 4 // 值类型对比不相等

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

	// 是否阻塞循环执行
	IsBlockUntil int `json:"is_block_until"`

	// 前置行为
	PreAction *Action `json:"pre_action"`

	// 后置行为
	SufAction *Action `json:"suf_action"`

	// session_id
	session_id string
}

// add session
func (a *Action) WithSessionId(id string) *Action {
	a.session_id = id
	return a
}

// start
func (a *Action) Run() (err error) {
	if a.IsBlockUntil == 1 {
		err = a.waitFor()
	} else {
		err = a.dispatch()
	}
	return err
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

func (a *Action) dispatch() error {
	// 操作类型
	switch a.ActionType {

	case ACTION_NAVIGATETO:
		_, err := driver.NavigateToUrl(a.session_id, a.ActionTarget)
		if err != nil {
			return err
		}
	case ACTION_VIEW_TITLE:
		rex, err := driver.GetTitle(a.session_id)
		if err != nil {
			return err
		}
		if rex.Value != a.ActionValue {
			return errors2.NewExpectError("预期值不符合")
		}
	case ACTION_CLICK:
		resp, err := driver.FindElement(a.session_id, a.getSelector())
		if err != nil {
			return err
		}
		_, err = driver.ElementClick(a.session_id, resp.Value.ElementId)
		if err != nil {
			return err
		}
	case ACTINO_SEND_KEYS:
		resp, err := driver.FindElement(a.session_id, a.getSelector())
		if err != nil {
			return err
		}
		_, err = driver.ElementSendKeys(a.session_id, resp.Value.ElementId, a.ActionValue)
		if err != nil {
			return err
		}
	case ACTION_VIEW_VALUE:
		resp, err := driver.FindElement(a.session_id, a.getSelector())
		if err != nil {
			return err
		}
		rex, err := driver.GetElementText(a.session_id, resp.Value.ElementId)
		if err != nil {
			return err
		}
		if rex.Value != a.ActionValue {
			return errors2.NewExpectError("预期值不符合")
		}
	case ACTION_SCREENSHOT:
		_, err := driver.TakeScreenshot(a.session_id, "asd.png")
		if err != nil {
			return err
		}
	case ACTION_SWITCH_WINDOW:
		resp, _ := driver.GetWindowHandles(a.session_id)
		handId := util.ToInt(a.ActionValue)
		_, err := driver.SwitchToWindow(a.session_id, resp.Value[handId])
		if err != nil {
			return err
		}
	case ACTION_NEW_WINDOW:
		_, err := driver.NewWindow(a.session_id)
		if err != nil {
			return err
		}
	case ACTINO_EXCUTE_SCRIPT:
		_, err := driver.ExcuteScript(a.session_id, a.ActionValue)
		if err != nil {
			return err
		}
	}
	if a.ActionDelay > 0 {
		time.Sleep(time.Duration(a.ActionDelay) * time.Second)
	}
	return nil
}

func (a *Action) waitFor() error {
	tick := time.Tick(1000 * time.Millisecond)
	timeout := time.After(300 * time.Second)
	end := make(chan int)
	var err error
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
	Loop:
		for {
			select {
			case <-timeout:
				err = errors.New("超时")
				break Loop
			case <-tick:
				_ = check(a, end)
			case <-end:
				break Loop
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
		wg.Done()
	}(wg)
	wg.Wait()
	return err
}

func check(a *Action, ch chan int) error {
	err := a.dispatch()
	if a.SufAction != nil {
		err = check(a.SufAction.WithSessionId(a.session_id), ch)
	}
	if err == nil {
		ch <- 1
	}
	return err
}
