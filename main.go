package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/wpp/comments/pkg/app"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	m := &app.MainData{
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
	if _, err := (MainWindow{
		Title:  "拼多多评论生成工具",
		Size:   Size{400, 600},
		Layout: VBox{},
		Children: []Widget{
			Composite{
				MaxSize: Size{0, 50},
				Layout:  HBox{},
				Children: []Widget{
					HSpacer{},
					TextLabel{
						Text: "AccessKey：",
					},
					LineEdit{
						AssignTo: &m.AccessKey,
						MaxSize:  Size{450, 0},
					},
				},
			},
			Composite{
				MaxSize: Size{0, 50},
				Layout:  HBox{},
				Children: []Widget{
					HSpacer{},
					TextLabel{
						Text: "商品ID：",
					},
					LineEdit{
						AssignTo: &m.ItemId,
						MaxSize:  Size{200, 0},
					},
					TextLabel{
						Text: "最低字数：",
					},
					LineEdit{
						AssignTo: &m.MinLength,
						MaxSize:  Size{100, 0},
					},
				},
			},
			Composite{
				Layout: VBox{},
				Children: []Widget{
					TextLabel{
						Text: "评语头编辑区：",
					},
					TextEdit{
						AssignTo: &m.CommentPrefix,
					},
					TextLabel{
						Text: "尾部标记词编辑区：",
					},
					TextEdit{
						AssignTo: &m.CommentSuffix,
					},
					TextLabel{
						Text: "添加过滤关键词，每个关键词之间用“,”隔离：",
					},
					TextEdit{
						AssignTo: &m.CommentFilter,
					},
					TextLabel{
						Text: "生成结果：",
					},
					TextEdit{
						AssignTo: &m.CommentResult,
						VScroll: true,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						Text:      "生成",
						OnClicked: m.Start,
						AssignTo:  &m.GenerateBtn,
					},
					PushButton{
						Text:      "清空",
						OnClicked: m.ClearCommentResult,
						AssignTo:  &m.ClearBtn,
					},
				},
			},
		},
	}).Run(); err != nil {
		fmt.Println(err)
	}
}
