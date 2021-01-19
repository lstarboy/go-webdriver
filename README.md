### go webdriver 

#### 功能说明
`go-webdriver` 是使用golang 开发的一套基于 `w3c-webdriver` 协议开发的一套chrome 浏览器驱动， 该包依赖 `chromedriver` 浏览器驱动， `chrome
` 浏览器， 用程序控制模拟用户操作浏览器，功能类似 `selenium`。 

#### 用途
* 模拟用户操作，配合用户操作，全自动化流程测试，或者半自动化测试
* 页面数据抓取，而不需要进行各种复杂的破解

#### 使用方法
* 依赖
  
  * chrome 浏览器， 此处依赖 `chrome 87.0.4280.141` `chrome 87.0.4280.88` 均正常(chrome 版本查看 `chrome 设置` -> `帮助` -> `关于google chrome`);
  * chromedriver chrome 浏览器驱动，下载和浏览器版本匹配的chromedriver。 下载地址：[chromedriver](https://npm.taobao.org/mirrors/chromedriver/)。
    使用 `chromedriver 87.0.4280.88` 等版本测试过正常；
    
```ssh
    ./chromedriver --port=9515 --allow-ips=
```

* 引用
    * 在`go.mod`中新增  
```text
   zhouzhe1157/go-webdriver => github.com/zhouzhe1157/go-webdriver v1.0.29
```

    * 加载依赖
```bigquery
    go mod tidy
    go mod vendor
```
    
* 实例， 流程流转，只需要不断的增加action, 就可以实现流程的自动化流转

```golang

    //logDir := "E:\\logs\\" + util.RandString(16)
    opts := excutor.ChromeOptions{IsHeadless: false, UserDataDir:""}
    resp, err := driver.GetSession(opts)
    if err != nil {
        return err
    }

    // 构建单通道
    pip := pipline.Pipline{Data: pipline.PipData{Actions: []action.Action{}}}

    // 邮箱名称
    len := util.RandInt(9, 14)
    username := util.RandString(len)

    // 操作步骤
    action1 := action.Action {
        ActionName: "打开页面",
        ActionType: action.ACTION_NAVIGATETO,
        ActionTarget: "https://www.baidu.com",
    }
    
    action2 := action.Action {
        ActionName: "输入参数",
        ActionType: action.ACTINO_SEND_KEYS,
		ActionTarget: "#kw",
		ActionValue: "golang",
	}
	
	action3 := action.Action {
        ActionName: "搜索",
        ActionType: action.ACTION_CLICK,
        ActionTarget: "#su",
        ActionDelay: 1,
	}

	pip.Data.Actions = append(pip.Data.Actions, action1, action2, action3)
	_ = pip.SetSessionId(resp.SessionId).Start()
	
```

* 示例 需要人为介入或者阻塞流程使用方法(等待用户输入了关键词之后，才会自动执行搜索操作)

```golang

    //logDir := "E:\\logs\\" + util.RandString(16)
    opts := excutor.ChromeOptions{IsHeadless: false, UserDataDir:""}
    resp, err := driver.GetSession(opts)
    if err != nil {
        return err
    }

    // 构建单通道
    pip := pipline.Pipline{Data: pipline.PipData{Actions: []action.Action{}}}

    // 邮箱名称
    len := util.RandInt(9, 14)
    username := util.RandString(len)

    // 操作步骤
    action1 := action.Action {
        ActionName: "打开页面",
        ActionType: action.ACTION_NAVIGATETO,
        ActionTarget: "https://www.baidu.com",
    }
    
    action2 := action.Action {
        ActionName: "输入参数",
        ActionType: action.ACTION_VIEW_VALUE,
        ActionTarget: "#kw",
        ExpectType: action.EXPECT_TYPE_EXIST,
	}
	
	action3 := action.Action {
        ActionName: "搜索",
        ActionType: action.ACTION_CLICK,
        ActionTarget: "#su",
        ActionDelay: 1,
        PreAction: &action2
	}

	pip.Data.Actions = append(pip.Data.Actions, action1, action3)
	_ = pip.SetSessionId(resp.SessionId).Start()
	
```

#### 待完成功能
* 更多版本浏览器，以及chromedriver的兼容
* 对`webdriver` `Command` 支持完善
* 更多场景的支持
* 对多浏览器的支持

#### 欢迎咨询，期待更多的小伙伴一起加入（QQ: 1157667735）