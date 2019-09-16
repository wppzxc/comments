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
