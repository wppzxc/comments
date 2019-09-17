package app

import (
	"fmt"
	"github.com/lxn/walk"
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
	m := &MainData{
		AccessKey:     &walk.LineEdit{},
		ItemId:        &walk.LineEdit{},
		MinLength:     &walk.LineEdit{},
		CommentPrefix: &walk.TextEdit{},
		CommentSuffix: &walk.TextEdit{},
		CommentFilter: &walk.TextEdit{},
		CommentResult: &walk.TextEdit{},
		GenerateBtn:   &walk.PushButton{},
		ClearBtn:      &walk.PushButton{},
	}
	m.ItemId.SetText("https://mobile.yangkeduo.com/proxy/api/reviews/4103183576/list?page=1&size=10&enable_video=1&enable_group_review=1")
	m.Start()
}
