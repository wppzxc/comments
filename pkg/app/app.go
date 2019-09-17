package app

import (
	"encoding/json"
	"fmt"
	"github.com/lxn/walk"
	"github.com/wpp/comments/pkg/types"
	"github.com/wpp/comments/pkg/utils"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	commentUrl = "https://mobile.yangkeduo.com/proxy/api/reviews/%s/list?page=%d&size=10&enable_video=1&enable_group_review=1"
)

type MainData struct {
	AccessKey     *walk.LineEdit
	ItemId        *walk.LineEdit
	MinLength     *walk.LineEdit
	CommentPrefix *walk.TextEdit
	CommentSuffix *walk.TextEdit
	CommentFilter *walk.TextEdit
	CommentResult *walk.TextEdit
	GenerateBtn   *walk.PushButton
	ClearBtn      *walk.PushButton
}

func (m *MainData) Start() {
	m.SetUIEnable(false)
	itemLink := m.ItemId.Text()
	u, err := url.Parse(itemLink)
	if err != nil {
		fmt.Println(err)
		return
	}
	param, _ := url.ParseQuery(u.RawQuery)
	itemId := param["goods_id"][0]
	if len(itemId) == 0 {
		fmt.Println("商品链接错误，无法解析商品id！")
		return
	}
	minLength := m.MinLength.Text()
	commentPrefix := m.CommentPrefix.Text()
	commentSuffix := m.CommentSuffix.Text()
	go func(){
		commentResult := m.GetCommentResult(itemId, minLength, commentPrefix, commentSuffix)
		m.CommentResult.SetText(commentResult)
	}()
}

func (m *MainData) GetCommentResult(itemId, minLenth, prefix, suffix string) string {
	defer m.SetUIEnable(true)
	comments := make([]types.Comment, 0)
	for i := 1; i <= 30; i ++ {
		cms := m.GetOnePageComments(itemId, i)
		cms = utils.RemoveEmptyComments(cms)
		cms = utils.RemoveLowScoreComments(cms)
		comments = append(comments, cms...)
	}
	comments = m.FilterKeys(comments)
	comment := m.GenerateComment(comments, minLenth)
	return prefix + comment + suffix
}

func (m *MainData) GetOnePageComments(itemId string, page int) []types.Comment {
	ak := m.AccessKey.Text()
	//ak := "U2AZ4HAJYPISSDVRHV2VME7TKV7XRD5XPRA5VC6VCLXSAKIKAJWA10132a5"
	client := &http.Client{}
	var req *http.Request
	itemUrl := fmt.Sprintf(commentUrl, itemId, page)
	req, _ = http.NewRequest(http.MethodGet, itemUrl, nil)
	req.Header.Add("accesstoken", ak)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	tmp := &types.CommentData{}
	if err := json.Unmarshal(data, tmp); err != nil {
		fmt.Println(err)
		return nil
	}
	return tmp.Data
}

func (m *MainData) FilterKeys(comments []types.Comment) []types.Comment {
	str := m.CommentFilter.Text()
	filters :=strings.Split(str, ",")
	if len(str) == 0 || len(filters) == 0{
		return comments
	}
	result := make([]types.Comment, 0)
	for _, c := range comments {
		add := true
		for _, f := range filters {
			if strings.Index(c.Comment, f) > 0 {
				add = false
			}
		}
		if add {
			result = append(result, c)
		}
	}
	return result
}

func (m *MainData) GenerateComment(comments []types.Comment, minLength string) string {
	result := []byte{}
	lastComment := ""
	length, _ := strconv.Atoi(minLength)
	size := len(comments)
	for _, c := range comments {
		fmt.Println(c)
		i := rand.Intn(size -1) + 1
		if comments[i].Comment == lastComment {
			continue
		}
		result = append(result, comments[i].Comment...)
		lastComment = comments[i].Comment
		fmt.Println(utf8.RuneCountInString(string(result)))
		if utf8.RuneCountInString(string(result)) > length {
			return string(result)
		}
		result = append(result, "，"...)
	}
	return string(result)
}

func (m *MainData) SetUIEnable(enable bool) {
	m.ItemId.SetEnabled(enable)
	m.MinLength.SetEnabled(enable)
	m.CommentPrefix.SetEnabled(enable)
	m.CommentSuffix.SetEnabled(enable)
	m.CommentFilter.SetEnabled(enable)
	m.CommentResult.SetEnabled(enable)
	m.GenerateBtn.SetEnabled(enable)
	m.ClearBtn.SetEnabled(enable)
}

func (m *MainData) ClearCommentResult() {
	m.CommentResult.SetText("")
}
