package app

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/zserge/webview"
	"testing"
)

func Test01(t *testing.T) {
	m := &MainData{
		AccessKey: &walk.LineEdit{},
		CommentFilter: &walk.TextEdit{},
	}
	//m.AccessKey.SetText("FZXFEXFQ6XLCW2WCISHOWKJ2WOEFQMXFQN2M7EXK7DJ6IS5HVUIQ10132a5")
	itemId := "35606797536"
	result := m.GetCommentResult(itemId, "150", "-----tou-----", "-----wei-----")
	fmt.Println(result)
}

func TestStart(t *testing.T) {
	
	js := `!function () {
    window.setInterval(function () {
        alert("test")
    }, 3000)
}();`
	fmt.Println(js)
	w := webview.New(webview.Settings{
		Width: 800,
		Height: 600,
		Title: "拼多多登录",
		Resizable: true,
		URL: "https://mobile.yangkeduo.com/login.html",
		Debug: true,
		ExternalInvokeCallback: syncAccessKey,
	})
	stopCh := make(chan bool)
	go w.Run()
	defer w.Exit()
	w.Dispatch(func(){
		if err := w.Eval(js); err != nil {
			fmt.Println(err)
			stopCh <- true
		}
	})
	//m := &MainData{
	//	AccessKey:     &walk.LineEdit{},
	//	ItemId:        &walk.LineEdit{},
	//	MinLength:     &walk.LineEdit{},
	//	CommentPrefix: &walk.TextEdit{},
	//	CommentSuffix: &walk.TextEdit{},
	//	CommentFilter: &walk.TextEdit{},
	//	CommentResult: &walk.TextEdit{},
	//	GenerateBtn:   &walk.PushButton{},
	//	ClearBtn:      &walk.PushButton{},
	//}
	//m.ItemId.SetText("https://mobile.yangkeduo.com/proxy/api/reviews/4103183576/list?page=1&size=10&enable_video=1&enable_group_review=1")
	//m.Start()
}

func syncAccessKey(w webview.WebView, data string) {
	fmt.Println(data)
}
