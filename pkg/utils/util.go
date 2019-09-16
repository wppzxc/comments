package utils

import "github.com/wpp/comments/pkg/types"

const (
	EmptyComment = "此用户未填写文字评论"
	FullScore = 5
)

func RemoveEmptyComments(comments []types.Comment) []types.Comment {
	result := make([]types.Comment, 0)
	for _, c := range comments {
		if c.Comment != EmptyComment {
			result = append(result, c)
		}
	}
	return result
}

func RemoveLowScoreComments(comments []types.Comment) []types.Comment {
	result := make([]types.Comment, 0)
	for _, c := range comments {
		if c.DescScore == FullScore &&
			c.LogisticsScore == FullScore && c.ServiceScore == FullScore {
			result = append(result, c)
		}
	}
	return result
}
